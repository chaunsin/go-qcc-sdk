#!/usr/bin/env python3
"""Fetch public QCC OpenAPI dataApi docs into qcc_check.py JSON input."""

from __future__ import annotations

import argparse
import concurrent.futures
import datetime as dt
import html
import json
import re
import sys
import time
import tempfile
import urllib.error
import urllib.request
import uuid
from pathlib import Path
from typing import Any

INDEX_URL = "https://openapi.qcc.com/dataApi"
DETAIL_URL = "https://openapi.qcc.com/dataApi/{code}"
API_HOST = "https://api.qichacha.com"
USER_AGENT = (
    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "
    "AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126 Safari/537.36"
)


def default_output_path() -> str:
    return str(Path(tempfile.gettempdir()) / f"qcc_official_docs_{uuid.uuid4().hex}.json")


def api_sort_key(apicode: str) -> tuple[int, Any]:
    return (0, int(apicode)) if apicode.isdigit() else (1, apicode)


def fetch_html(url: str, timeout: float, retries: int = 0, retry_delay: float = 0.75) -> str:
    last_exc: BaseException | None = None
    for attempt in range(retries + 1):
        request = urllib.request.Request(
            url,
            headers={
                "User-Agent": USER_AGENT,
                "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
                "Accept-Language": "zh-CN,zh;q=0.9,en;q=0.7",
            },
        )
        try:
            with urllib.request.urlopen(request, timeout=timeout) as response:
                return response.read().decode("utf-8", "ignore")
        except urllib.error.HTTPError as exc:
            if exc.code != 429 and exc.code < 500:
                raise
            last_exc = exc
        except (OSError, TimeoutError, urllib.error.URLError) as exc:
            last_exc = exc

        if attempt < retries:
            time.sleep(retry_delay * (2**attempt))

    if last_exc:
        raise last_exc
    raise RuntimeError(f"failed to fetch {url}")


def clean_text(value: str) -> str:
    value = re.sub(r"<script\b.*?</script>", " ", value, flags=re.S | re.I)
    value = re.sub(r"<style\b.*?</style>", " ", value, flags=re.S | re.I)
    value = re.sub(r"<[^>]+>", " ", value)
    value = html.unescape(value)
    return re.sub(r"\s+", " ", value).strip()


def first_match(patterns: list[str], text: str) -> str:
    for pattern in patterns:
        match = re.search(pattern, text, flags=re.S | re.I)
        if match:
            return clean_text(match.group(1))
    return ""


def ordered_unique(values: list[str]) -> list[str]:
    seen: set[str] = set()
    result: list[str] = []
    for value in values:
        value = value.strip()
        if value and value not in seen:
            seen.add(value)
            result.append(value)
    return result


def extract_index_codes(index_html: str) -> list[str]:
    # The Nuxt SSR payload currently stores index records with an ApiCode key.
    codes: list[str] = []
    patterns = [
        r"['\"]?[Aa]piCode['\"]?\s*:\s*['\"](\d+)['\"]",
        r"\\['\"]?[Aa]piCode\\?['\"]?\s*:\s*\\?['\"](\d+)\\?['\"]",
    ]
    for pattern in patterns:
        codes.extend(re.findall(pattern, index_html))
    return sorted(ordered_unique(codes), key=lambda code: int(code))


def extract_ref_tab_names(detail_html: str) -> list[str]:
    tab = re.search(r'<ul[^>]+id=["\']refTab["\'][^>]*>(.*?)</ul>', detail_html, flags=re.S | re.I)
    if not tab:
        return []

    names: list[str] = []
    for match in re.finditer(r"<li\b[^>]*>\s*<a\b[^>]*>(.*?)</a>\s*</li>", tab.group(1), flags=re.S | re.I):
        name = clean_text(match.group(1))
        if name:
            names.append(name)
    return ordered_unique(names)


def extract_endpoint_urls(detail_html: str) -> list[str]:
    urls = re.findall(
        r"接口地址：\s*</span>\s*<span[^>]*>\s*(https?://api\.qichacha\.com/[^<\s]+)\s*</span>",
        detail_html,
        flags=re.S | re.I,
    )
    if not urls:
        urls = re.findall(r"https?://api\.qichacha\.com/[A-Za-z0-9_./-]+", detail_html)
    return ordered_unique([html.unescape(url) for url in urls])


def title_from_html(detail_html: str) -> str:
    title = first_match(
        [
            r'<h1[^>]*class=["\'][^"\']*title[^"\']*["\'][^>]*>(.*?)</h1>',
            r'<title[^>]*>(.*?)</title>',
        ],
        detail_html,
    )
    suffixes = [
        "API数据接口 - 企查查开放平台",
        " - 企查查开放平台",
    ]
    for suffix in suffixes:
        if title.endswith(suffix):
            title = title[: -len(suffix)].strip()
    return title


def description_from_html(detail_html: str) -> str:
    return first_match(
        [
            r'<meta[^>]+name=["\']description["\'][^>]+content=["\']([^"\']*)["\']',
            r'<div[^>]+class=["\'][^"\']*describe[^"\']*["\'][^>]*>\s*<span[^>]*>描述：</span>\s*<span[^>]*>(.*?)</span>',
        ],
        detail_html,
    )


def is_rendered_404(detail_html: str) -> bool:
    return any(marker in detail_html for marker in ("error404", "statusCode:404", "This page could not be found"))


def interfaces_from_detail(detail_html: str) -> tuple[list[dict[str, str]], list[str]]:
    names = extract_ref_tab_names(detail_html)
    urls = extract_endpoint_urls(detail_html)
    warnings: list[str] = []

    if names and urls and len(names) != len(urls):
        warnings.append(f"interface tab count {len(names)} does not match API address count {len(urls)}")

    interfaces: list[dict[str, str]] = []
    for index in range(max(len(names), len(urls))):
        name = names[index] if index < len(names) else ""
        endpoint = urls[index] if index < len(urls) else ""
        path = endpoint.removeprefix(API_HOST) if endpoint.startswith(API_HOST) else endpoint
        interfaces.append(
            {
                "description": name or path or f"interface-{index + 1}",
                "path": path,
            }
        )

    if not interfaces:
        warnings.append("no interface tabs or API addresses found")
    return interfaces, warnings


def build_detail_record(code: str, detail_html: str) -> dict[str, Any]:
    interfaces, warnings = interfaces_from_detail(detail_html)
    record: dict[str, Any] = {
        "apicode": code,
        "title": title_from_html(detail_html),
        "description": description_from_html(detail_html),
        "url": DETAIL_URL.format(code=code),
        "interfaces": interfaces,
        "interface_count": len(interfaces),
    }
    if is_rendered_404(detail_html):
        warnings.insert(0, "rendered 404 page")
    if warnings:
        record["warnings"] = warnings
    return record


def fetch_detail_record(code: str, timeout: float, retries: int, retry_delay: float) -> dict[str, Any]:
    url = DETAIL_URL.format(code=code)
    try:
        return build_detail_record(code, fetch_html(url, timeout, retries, retry_delay))
    except urllib.error.HTTPError as exc:
        return {
            "apicode": code,
            "title": "",
            "description": "",
            "url": url,
            "interfaces": [],
            "interface_count": 0,
            "error": f"HTTP {exc.code}",
        }
    except Exception as exc:  # noqa: BLE001 - CLI should preserve fetch failure details.
        return {
            "apicode": code,
            "title": "",
            "description": "",
            "url": url,
            "interfaces": [],
            "interface_count": 0,
            "error": f"{type(exc).__name__}: {exc}",
        }


def read_offline_detail(detail_dir: Path, code: str) -> dict[str, Any]:
    path = detail_dir / f"{code}.html"
    if not path.exists():
        return {
            "apicode": code,
            "title": "",
            "description": "",
            "url": DETAIL_URL.format(code=code),
            "interfaces": [],
            "interface_count": 0,
            "error": f"missing offline detail HTML: {path}",
        }
    return build_detail_record(code, path.read_text(encoding="utf-8", errors="ignore"))


def collect_warnings(docs: list[dict[str, Any]]) -> list[dict[str, Any]]:
    warnings: list[dict[str, Any]] = []
    for record in docs:
        messages = []
        if record.get("error"):
            messages.append(str(record["error"]))
        messages.extend(str(item) for item in record.get("warnings", []))
        if messages:
            warnings.append(
                {
                    "apicode": record.get("apicode", ""),
                    "url": record.get("url", ""),
                    "title": record.get("title", ""),
                    "messages": messages,
                }
            )
    return warnings


def load_index_html(args: argparse.Namespace) -> str:
    if args.index_html:
        return Path(args.index_html).read_text(encoding="utf-8", errors="ignore")
    return fetch_html(INDEX_URL, args.timeout, args.retries, args.retry_delay)


def fetch_docs(codes: list[str], args: argparse.Namespace) -> list[dict[str, Any]]:
    if args.detail_dir:
        detail_dir = Path(args.detail_dir)
        return sorted((read_offline_detail(detail_dir, code) for code in codes), key=lambda item: api_sort_key(item["apicode"]))

    docs: list[dict[str, Any]] = []
    with concurrent.futures.ThreadPoolExecutor(max_workers=max(1, args.workers)) as executor:
        futures = {
            executor.submit(fetch_detail_record, code, args.timeout, args.retries, args.retry_delay): code
            for code in codes
        }
        for index, future in enumerate(concurrent.futures.as_completed(futures), start=1):
            docs.append(future.result())
            if args.verbose and (index % 25 == 0 or index == len(codes)):
                print(f"detail pages fetched: {index}/{len(codes)}", file=sys.stderr)
    return sorted(docs, key=lambda item: api_sort_key(item["apicode"]))


def build_payload(index_html: str, docs: list[dict[str, Any]]) -> dict[str, Any]:
    warnings = collect_warnings(docs)
    return {
        "source": INDEX_URL,
        "fetched_at": dt.datetime.now(dt.timezone.utc).isoformat(),
        "docs": docs,
        "warnings": warnings,
        "summary": {
            "apicode_count": len(docs),
            "interface_count": sum(int(record.get("interface_count") or 0) for record in docs),
            "warning_count": len(warnings),
            "index_html_bytes": len(index_html.encode("utf-8")),
        },
    }


def write_payload(path: Path, payload: dict[str, Any], exclusive: bool) -> None:
    path.parent.mkdir(parents=True, exist_ok=True)
    text = json.dumps(payload, ensure_ascii=False, indent=2)
    if exclusive:
        with path.open("x", encoding="utf-8") as handle:
            handle.write(text)
            handle.write("\n")
    else:
        path.write_text(text + "\n", encoding="utf-8")


def print_warning_summary(payload: dict[str, Any]) -> None:
    warnings = payload["warnings"]
    for warning in warnings[:20]:
        print(
            f"warning {warning['apicode']}: " + "; ".join(warning["messages"]),
            file=sys.stderr,
        )
    if len(warnings) > 20:
        print(f"... +{len(warnings) - 20} more warnings", file=sys.stderr)


def emit_summary(status: str, payload: dict[str, Any], out: Path | None, started: float, summary_json: bool) -> None:
    summary = payload["summary"]
    elapsed = time.time() - started
    if summary_json:
        message = {
            "status": status,
            "apicode_count": summary["apicode_count"],
            "interface_count": summary["interface_count"],
            "warning_count": summary["warning_count"],
            "out": str(out) if out else None,
            "elapsed_seconds": round(elapsed, 3),
        }
        print(json.dumps(message, ensure_ascii=False, sort_keys=True))
        return

    parts = [
        "official docs extracted:",
        f"status={status}",
        f"apicodes={summary['apicode_count']}",
        f"interfaces={summary['interface_count']}",
        f"warnings={summary['warning_count']}",
    ]
    if out:
        parts.append(f"out={out}")
    parts.append(f"elapsed={elapsed:.1f}s")
    print(" ".join(parts))


def parse_args(argv: list[str]) -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--out", help="Path to write official docs JSON; defaults to a unique file in the platform temp directory")
    parser.add_argument("--print-temp-dir", action="store_true", help="Print Python's platform temp directory and exit")
    parser.add_argument("--print-out-path-only", action="store_true", help="Print the selected output path and exit without fetching")
    parser.add_argument("--summary-json", action="store_true", help="Print a machine-readable one-line JSON summary")
    parser.add_argument("--workers", type=int, default=8, help="Concurrent detail-page fetch workers")
    parser.add_argument("--timeout", type=float, default=25.0, help="HTTP timeout in seconds")
    parser.add_argument("--retries", type=int, default=2, help="Retry attempts for transient HTTP/network failures")
    parser.add_argument("--retry-delay", type=float, default=0.75, help="Initial retry delay in seconds; later retries use exponential backoff")
    parser.add_argument("--allow-warnings", action="store_true", help="Write JSON and exit 0 even when extraction warnings exist")
    parser.add_argument("--verbose", action="store_true", help="Print progress to stderr")
    parser.add_argument("--index-html", help="Offline index HTML fixture for parser validation")
    parser.add_argument("--detail-dir", help="Offline directory containing {ApiCode}.html detail fixtures")
    return parser.parse_args(argv)


def main(argv: list[str]) -> int:
    args = parse_args(argv)
    if args.print_temp_dir:
        print(tempfile.gettempdir())
        return 0
    out = Path(args.out or default_output_path())
    out_is_default = args.out is None
    if args.print_out_path_only:
        print(out)
        return 0

    started = time.time()

    index_html = load_index_html(args)
    codes = extract_index_codes(index_html)
    if not codes:
        print("no ApiCode entries found in official dataApi index HTML", file=sys.stderr)
        return 2

    docs = fetch_docs(codes, args)
    payload = build_payload(index_html, docs)
    if payload["warnings"] and not args.allow_warnings:
        print_warning_summary(payload)
        print("refusing to pass with extraction warnings; use --allow-warnings only for manual inspection", file=sys.stderr)
        emit_summary("failed", payload, None, started, args.summary_json)
        return 1

    status = "warning" if payload["warnings"] else "ok"
    write_payload(out, payload, exclusive=out_is_default)
    emit_summary(status, payload, out, started, args.summary_json)
    return 0


if __name__ == "__main__":
    raise SystemExit(main(sys.argv[1:]))
