#!/usr/bin/env python3
"""Audit go-qcc-sdk ApiCode coverage against extracted QCC docs JSON."""

from __future__ import annotations

import argparse
import html
import json
import os
import platform
import re
import sys
from dataclasses import dataclass
from pathlib import Path
from typing import Any, Iterable


METHOD_RE = re.compile(r"func\s*\(\s*\w+\s+\*Api\s*\)\s+([A-Za-z_]\w*)\s*\(")
GOOS_BY_PLATFORM = {
    "darwin": "darwin",
    "linux": "linux",
    "win32": "windows",
    "cygwin": "windows",
    "freebsd": "freebsd",
    "openbsd": "openbsd",
    "netbsd": "netbsd",
}
GOARCH_BY_MACHINE = {
    "x86_64": "amd64",
    "amd64": "amd64",
    "aarch64": "arm64",
    "arm64": "arm64",
    "i386": "386",
    "i686": "386",
}
UNIX_GOOS = {
    "aix",
    "android",
    "darwin",
    "dragonfly",
    "freebsd",
    "hurd",
    "illumos",
    "ios",
    "linux",
    "netbsd",
    "openbsd",
    "solaris",
}


@dataclass(frozen=True)
class LocalApi:
    apicode: str
    file: str
    method_count: int
    methods: list[str]


@dataclass(frozen=True)
class DocApi:
    apicode: str
    title: str
    url: str
    interface_count: int
    interfaces: list[str]


def first_value(mapping: dict[str, Any], keys: Iterable[str], default: Any = "") -> Any:
    for key in keys:
        if key in mapping and mapping[key] not in (None, ""):
            return mapping[key]
    return default


def api_sort_key(apicode: str) -> tuple[int, Any]:
    return (0, int(apicode)) if apicode.isdigit() else (1, apicode)


def default_build_tags() -> set[str]:
    goos = os.environ.get("GOOS") or GOOS_BY_PLATFORM.get(sys.platform, sys.platform)
    goarch = os.environ.get("GOARCH") or GOARCH_BY_MACHINE.get(platform.machine().lower(), platform.machine().lower())
    tags = {goos, goarch, "gc"}
    if goos in UNIX_GOOS:
        tags.add("unix")
    if os.environ.get("CGO_ENABLED", "1") != "0":
        tags.add("cgo")
    # Go release tags are monotonically true for modern Go versions; include a
    # broad range so common version constraints do not exclude active files.
    tags.update(f"go1.{minor}" for minor in range(1, 40))
    return tags


def tokenize_build_expr(expr: str) -> list[str]:
    tokens: list[str] = []
    index = 0
    while index < len(expr):
        char = expr[index]
        if char.isspace():
            index += 1
            continue
        if expr.startswith("&&", index) or expr.startswith("||", index):
            tokens.append(expr[index : index + 2])
            index += 2
            continue
        if char in "!()":
            tokens.append(char)
            index += 1
            continue
        match = re.match(r"[A-Za-z0-9_./-]+", expr[index:])
        if not match:
            raise ValueError(f"invalid go:build expression near {expr[index:]!r}")
        tokens.append(match.group(0))
        index += len(match.group(0))
    return tokens


class BuildExprParser:
    def __init__(self, tokens: list[str], tags: set[str]):
        self.tokens = tokens
        self.tags = tags
        self.index = 0

    def parse(self) -> bool:
        value = self.parse_or()
        if self.index != len(self.tokens):
            raise ValueError("unexpected token in go:build expression")
        return value

    def accept(self, token: str) -> bool:
        if self.index < len(self.tokens) and self.tokens[self.index] == token:
            self.index += 1
            return True
        return False

    def parse_or(self) -> bool:
        value = self.parse_and()
        while self.accept("||"):
            rhs = self.parse_and()
            value = value or rhs
        return value

    def parse_and(self) -> bool:
        value = self.parse_unary()
        while self.accept("&&"):
            rhs = self.parse_unary()
            value = value and rhs
        return value

    def parse_unary(self) -> bool:
        if self.accept("!"):
            return not self.parse_unary()
        if self.accept("("):
            value = self.parse_or()
            if not self.accept(")"):
                raise ValueError("missing closing parenthesis in go:build expression")
            return value
        if self.index >= len(self.tokens):
            raise ValueError("unexpected end of go:build expression")
        tag = self.tokens[self.index]
        self.index += 1
        return tag in self.tags


def eval_go_build_expr(expr: str, tags: set[str]) -> bool:
    return BuildExprParser(tokenize_build_expr(expr), tags).parse()


def eval_plus_build_line(line: str, tags: set[str]) -> bool:
    options = line.split()
    if not options:
        return True
    for option in options:
        terms = option.split(",")
        if all((term[1:] not in tags) if term.startswith("!") else (term in tags) for term in terms):
            return True
    return False


def leading_build_constraints(text: str) -> tuple[list[str], list[str]]:
    go_build: list[str] = []
    plus_build: list[str] = []
    for line in text.splitlines():
        stripped = line.strip()
        if stripped == "":
            # A blank line ends the leading comment block only after constraints.
            if go_build or plus_build:
                break
            continue
        if stripped.startswith("//go:build "):
            go_build.append(stripped.removeprefix("//go:build ").strip())
            continue
        if stripped.startswith("// +build "):
            plus_build.append(stripped.removeprefix("// +build ").strip())
            continue
        if stripped.startswith("//") or stripped.startswith("/*") or stripped.startswith("*"):
            continue
        break
    return go_build, plus_build


def build_constraints_match(text: str, tags: set[str]) -> bool:
    go_build, plus_build = leading_build_constraints(text)
    if go_build and not all(eval_go_build_expr(expr, tags) for expr in go_build):
        return False
    if plus_build and not all(eval_plus_build_line(line, tags) for line in plus_build):
        return False
    return True


def strip_go_comments_and_literals(text: str) -> str:
    output: list[str] = []
    index = 0
    state = "code"
    while index < len(text):
        char = text[index]
        nxt = text[index + 1] if index + 1 < len(text) else ""

        if state == "code":
            if char == "/" and nxt == "/":
                output.extend("  ")
                index += 2
                state = "line_comment"
                continue
            if char == "/" and nxt == "*":
                output.extend("  ")
                index += 2
                state = "block_comment"
                continue
            if char == "`":
                output.append(" ")
                index += 1
                state = "raw_string"
                continue
            if char == '"':
                output.append(" ")
                index += 1
                state = "string"
                continue
            if char == "'":
                output.append(" ")
                index += 1
                state = "rune"
                continue
            output.append(char)
            index += 1
            continue

        if state == "line_comment":
            output.append("\n" if char == "\n" else " ")
            index += 1
            if char == "\n":
                state = "code"
            continue

        if state == "block_comment":
            if char == "*" and nxt == "/":
                output.extend("  ")
                index += 2
                state = "code"
                continue
            output.append("\n" if char == "\n" else " ")
            index += 1
            continue

        if state == "raw_string":
            output.append("\n" if char == "\n" else " ")
            index += 1
            if char == "`":
                state = "code"
            continue

        if state in {"string", "rune"}:
            if char == "\\" and index + 1 < len(text):
                output.extend("  ")
                index += 2
                continue
            output.append("\n" if char == "\n" else " ")
            index += 1
            if (state == "string" and char == '"') or (state == "rune" and char == "'"):
                state = "code"
            continue

    return "".join(output)


def scan_local(repo: Path) -> list[LocalApi]:
    if not repo.exists():
        raise FileNotFoundError(f"repo does not exist: {repo}")
    if not repo.is_dir():
        raise NotADirectoryError(f"repo is not a directory: {repo}")

    apis: list[LocalApi] = []
    tags = default_build_tags()
    for path in sorted(repo.glob("*.go")):
        if not path.stem.isdigit():
            continue
        text = path.read_text(encoding="utf-8")
        if not build_constraints_match(text, tags):
            continue
        methods = METHOD_RE.findall(strip_go_comments_and_literals(text))
        apis.append(
            LocalApi(
                apicode=path.stem,
                file=str(path.relative_to(repo)),
                method_count=len(methods),
                methods=methods,
            )
        )
    return sorted(apis, key=lambda item: api_sort_key(item.apicode))


def stringify_interface(item: Any) -> str:
    if isinstance(item, str):
        return item.strip()
    if isinstance(item, dict):
        desc = first_value(
            item,
            (
                "description",
                "desc",
                "title",
                "name",
                "interface",
                "interfaceName",
                "apiName",
                "path",
                "url",
            ),
        )
        path = first_value(item, ("path", "endpoint", "url"), "")
        if desc and path and path not in str(desc):
            return f"{desc} ({path})".strip()
        return str(desc).strip()
    return str(item).strip()


def normalize_interfaces(record: dict[str, Any]) -> tuple[int, list[str]]:
    raw = first_value(record, ("interfaces", "endpoints", "apis", "items"), None)
    interfaces: list[str] = []
    if isinstance(raw, list):
        interfaces = [text for text in (stringify_interface(item) for item in raw) if text]
    elif isinstance(raw, dict):
        interfaces = [text for text in (stringify_interface(item) for item in raw.values()) if text]
    elif isinstance(raw, str):
        interfaces = [line.strip() for line in raw.splitlines() if line.strip()]

    count_value = first_value(record, ("interface_count", "interfaceCount", "count", "apiCount"), None)
    if count_value is None:
        count = len(interfaces)
    else:
        try:
            count = int(count_value)
        except (TypeError, ValueError) as exc:
            raise ValueError(f"invalid interface count for ApiCode {record!r}") from exc

    if count_value is None and not interfaces:
        apicode = first_value(record, ("apicode", "ApiCode", "apiCode", "code", "api_code"), "<unknown>")
        raise ValueError(
            f"ApiCode {apicode} is missing interfaces/interface_count; "
            "provide detail-page extraction data instead of index-only metadata"
        )

    return count, interfaces


def load_docs(path: Path) -> list[DocApi]:
    data = json.loads(path.read_text(encoding="utf-8"))
    records = data.get("docs", data.get("data", data.get("items"))) if isinstance(data, dict) else data
    if not isinstance(records, list):
        raise ValueError("docs JSON must be a list or an object containing a docs/data/items list")

    docs: list[DocApi] = []
    seen: set[str] = set()
    for record in records:
        if not isinstance(record, dict):
            raise ValueError(f"each docs record must be an object, got {type(record).__name__}")
        apicode = str(first_value(record, ("apicode", "ApiCode", "apiCode", "code", "api_code"))).strip()
        if not apicode:
            raise ValueError(f"docs record is missing ApiCode: {record!r}")
        if apicode in seen:
            raise ValueError(f"duplicate ApiCode in docs JSON: {apicode}")
        seen.add(apicode)

        interface_count, interfaces = normalize_interfaces(record)
        title = str(first_value(record, ("title", "name", "description", "desc"), "")).strip()
        url = str(first_value(record, ("url", "doc_url", "detail_url", "detailUrl"), "")).strip()
        if not url:
            url = f"https://openapi.qcc.com/dataApi/{apicode}"

        docs.append(
            DocApi(
                apicode=apicode,
                title=title,
                url=url,
                interface_count=interface_count,
                interfaces=interfaces,
            )
        )
    return sorted(docs, key=lambda item: api_sort_key(item.apicode))


def local_json(local: list[LocalApi]) -> dict[str, Any]:
    return {
        "local": [
            {
                "apicode": item.apicode,
                "file": item.file,
                "method_count": item.method_count,
                "methods": item.methods,
            }
            for item in local
        ],
        "summary": {
            "apicode_count": len(local),
            "method_count": sum(item.method_count for item in local),
        },
    }


def compare(docs: list[DocApi], local: list[LocalApi]) -> dict[str, Any]:
    docs_by_code = {item.apicode: item for item in docs}
    local_by_code = {item.apicode: item for item in local}
    missing: list[dict[str, Any]] = []
    deprecated: list[dict[str, Any]] = []
    implemented: list[dict[str, Any]] = []

    for apicode in sorted(docs_by_code, key=api_sort_key):
        doc = docs_by_code[apicode]
        loc = local_by_code.get(apicode)
        local_count = loc.method_count if loc else 0
        if local_count < doc.interface_count:
            missing.append(
                {
                    "apicode": apicode,
                    "title": doc.title,
                    "url": doc.url,
                    "official_count": doc.interface_count,
                    "local_count": local_count,
                    "missing_count": doc.interface_count - local_count,
                    "interfaces": doc.interfaces,
                    "local_methods": loc.methods if loc else [],
                }
            )
        elif local_count > doc.interface_count:
            deprecated.append(
                {
                    "apicode": apicode,
                    "reason": "local method count exceeds official interface count",
                    "file": loc.file if loc else "",
                    "official_count": doc.interface_count,
                    "local_count": local_count,
                    "local_methods": loc.methods if loc else [],
                    "interfaces": doc.interfaces,
                    "url": doc.url,
                }
            )
        else:
            implemented.append(
                {
                    "apicode": apicode,
                    "title": doc.title,
                    "count": doc.interface_count,
                    "file": loc.file if loc else "",
                    "methods": loc.methods if loc else [],
                }
            )

    for apicode in sorted(set(local_by_code) - set(docs_by_code), key=api_sort_key):
        loc = local_by_code[apicode]
        deprecated.append(
            {
                "apicode": apicode,
                "reason": "official ApiCode not found in docs index",
                "file": loc.file,
                "official_count": 0,
                "local_count": loc.method_count,
                "local_methods": loc.methods,
                "interfaces": [],
                "url": f"https://openapi.qcc.com/dataApi/{apicode}",
            }
        )

    return {
        "summary": {
            "official_apicode_count": len(docs),
            "official_interface_count": sum(item.interface_count for item in docs),
            "local_apicode_count": len(local),
            "local_method_count": sum(item.method_count for item in local),
            "matched_apicode_count": len(implemented),
            "missing_apicode_count": len(missing),
            "deprecated_or_extra_count": len(deprecated),
        },
        "missing": missing,
        "deprecated": sorted(deprecated, key=lambda item: api_sort_key(item["apicode"])),
        "implemented": implemented,
    }


def join_items(items: list[str], limit: int = 4) -> str:
    if not items:
        return "-"
    shown = items[:limit]
    suffix = "" if len(items) <= limit else f"; ... +{len(items) - limit}"
    return "; ".join(shown) + suffix


def md_cell(value: Any) -> str:
    text = html.escape(str(value), quote=False)
    text = text.replace("\r\n", "\n").replace("\r", "\n").replace("\n", "<br>")
    return text.replace("|", r"\|")


def markdown_report(result: dict[str, Any]) -> str:
    summary = result["summary"]
    missing = result["missing"]
    deprecated = result["deprecated"]
    implemented = result["implemented"]
    official_total = summary["official_apicode_count"]
    matched = summary["matched_apicode_count"]
    coverage = (matched / official_total * 100) if official_total else 0.0

    lines = [
        "# QCC Interface Coverage Audit",
        "",
        "## Summary",
        "",
        f"- Official ApiCodes: {summary['official_apicode_count']}",
        f"- Official interfaces: {summary['official_interface_count']}",
        f"- Local ApiCode files: {summary['local_apicode_count']}",
        f"- Local SDK methods: {summary['local_method_count']}",
        f"- Count-matched ApiCodes: {matched} ({coverage:.1f}%)",
        f"- ApiCodes with missing interfaces: {summary['missing_apicode_count']}",
        f"- Deprecated or suspicious local entries: {summary['deprecated_or_extra_count']}",
        "",
        "## 当前缺失接口",
        "",
    ]

    if missing:
        lines.extend(
            [
                "| ApiCode | 文档描述 | 官方接口数 | 本地方法数 | 缺失数 | 接口描述 |",
                "| --- | --- | ---: | ---: | ---: | --- |",
            ]
        )
        for item in missing:
            title = md_cell(item["title"] or "-")
            interfaces = md_cell(join_items(item["interfaces"]))
            lines.append(
                f"| {md_cell(item['apicode'])} | {title} | {item['official_count']} | "
                f"{item['local_count']} | {item['missing_count']} | {interfaces} |"
            )
        lines.append("")
        lines.append("说明：数量审查无法在本地数量不足时精确证明哪一个接口缺失，表中接口描述是该 ApiCode 下需要核对的官方候选接口。")
    else:
        lines.append("- 未发现官方接口数大于本地方法数的 ApiCode。")

    lines.extend(["", "## 已经废弃或疑似多余的本地实现", ""])
    if deprecated:
        lines.extend(
            [
                "| ApiCode | 原因 | 本地文件 | 官方接口数 | 本地方法数 | 本地方法 |",
                "| --- | --- | --- | ---: | ---: | --- |",
            ]
        )
        for item in deprecated:
            lines.append(
                f"| {md_cell(item['apicode'])} | {md_cell(item['reason'])} | {md_cell(item['file'])} | "
                f"{item['official_count']} | {item['local_count']} | {md_cell(join_items(item['local_methods']))} |"
            )
    else:
        lines.append("- 未发现官方文档不存在但本地存在，或本地方法数超过官方接口数的 ApiCode。")

    lines.extend(["", "## 已经存在实现", ""])
    if implemented:
        total_interfaces = sum(item["count"] for item in implemented)
        samples = implemented[:8]
        sample_text = ", ".join(f"{item['apicode']}({item['count']})" for item in samples)
        if len(implemented) > len(samples):
            sample_text += f", ... +{len(implemented) - len(samples)}"
        lines.append(f"- 数量匹配的 ApiCodes: {len(implemented)}，覆盖官方接口数: {total_interfaces}。")
        lines.append(f"- 样例: {sample_text}")
        lines.append("- 为保持报告简洁，已存在实现不完整罗列；如需明细，可要求输出 JSON 或指定 ApiCode。")
    else:
        lines.append("- 未发现数量完全匹配的 ApiCode。")

    lines.extend(
        [
            "",
            "## 后续建议",
            "",
            "- 缺失接口：指定 ApiCode 后，使用当前项目的 `qcc-create` skill 创建对应 `{ApiCode}.go` 实现。",
            "- 数量不一致：使用 `qcc-create` 按官方文档刷新对应 ApiCode，再重新运行本审查。",
            "- 官方已不存在但本地存在：请先选择保留兼容、删除实现，或标记 deprecated；不要在未确认兼容策略前直接删除。",
        ]
    )
    return "\n".join(lines) + "\n"


def parse_args(argv: list[str]) -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--repo", default=".", help="Path to go-qcc-sdk repository root")
    parser.add_argument("--local-json", action="store_true", help="Print local scan JSON and exit")
    parser.add_argument("--docs-json", help="Official docs extraction JSON to compare against")
    parser.add_argument("--output", help="Optional path for the Markdown report")
    return parser.parse_args(argv)


def main(argv: list[str]) -> int:
    args = parse_args(argv)
    repo = Path(args.repo).resolve()
    local = scan_local(repo)

    if args.local_json:
        print(json.dumps(local_json(local), ensure_ascii=False, indent=2))
        return 0

    if not args.docs_json:
        print("--docs-json is required unless --local-json is used", file=sys.stderr)
        return 2

    docs = load_docs(Path(args.docs_json))
    result = compare(docs, local)
    report = markdown_report(result)
    if args.output:
        Path(args.output).write_text(report, encoding="utf-8")
    else:
        print(report, end="")
    return 0


if __name__ == "__main__":
    raise SystemExit(main(sys.argv[1:]))
