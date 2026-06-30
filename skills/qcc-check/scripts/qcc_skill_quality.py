#!/usr/bin/env python3
"""Local quality checks for the qcc-check skill package."""

from __future__ import annotations

import json
import py_compile
import re
import subprocess
import sys
import tempfile
from pathlib import Path

ROOT = Path(__file__).resolve().parents[3]
SKILL = ROOT / "skills/qcc-check"
FETCH = SKILL / "scripts/qcc_fetch_docs.py"
CHECK = SKILL / "scripts/qcc_check.py"
EVALS = SKILL / "evals/evals.json"


def run(args: list[str], check: bool = True) -> subprocess.CompletedProcess[str]:
    return subprocess.run(args, cwd=ROOT, text=True, capture_output=True, check=check)


def assert_true(condition: bool, message: str) -> None:
    if not condition:
        raise AssertionError(message)


def check_python_files() -> None:
    py_compile.compile(str(FETCH), doraise=True)
    py_compile.compile(str(CHECK), doraise=True)


def check_evals_schema() -> None:
    data = json.loads(EVALS.read_text(encoding="utf-8"))
    assert_true(data.get("skill_name") == "qcc-check", "evals.skill_name must be qcc-check")
    evals = data.get("evals")
    assert_true(isinstance(evals, list) and evals, "evals must be a non-empty list")
    seen: set[int] = set()
    for item in evals:
        assert_true(isinstance(item, dict), "each eval must be an object")
        for key in ("id", "prompt", "expected_output", "files", "expectations"):
            assert_true(key in item, f"eval missing {key}: {item}")
        assert_true(item["id"] not in seen, f"duplicate eval id: {item['id']}")
        seen.add(item["id"])
        assert_true(isinstance(item["expectations"], list) and item["expectations"], f"eval {item['id']} needs expectations")


def make_success_fixture(base: Path) -> tuple[Path, Path]:
    details = base / "details"
    details.mkdir(parents=True)
    index = base / "index.html"
    index.write_text(
        '<script>window.__NUXT__={list:[{ApiCode:"213",Title:"企业年报信息"},{"ApiCode":"1124","Title":"企业财税数据"}]}</script>',
        encoding="utf-8",
    )
    (details / "213.html").write_text(
        """<!doctype html><html><body><h1 class="title">企业年报信息</h1><ul id="refTab"><li><a>查询公司年报信息</a></li><li><a>查询公司年报概况</a></li></ul><span>接口地址：</span> <span>https://api.qichacha.com/AR/GetAnnualReport</span><span>接口地址：</span> <span>https://api.qichacha.com/AR/GetAnnualReportSummary</span></body></html>""",
        encoding="utf-8",
    )
    (details / "1124.html").write_text(
        """<!doctype html><html><body><h1 class="title">企业财税数据</h1><ul id="refTab"><li><a>数据下单</a></li><li><a>验证码发送</a></li><li><a>数据获取</a></li></ul><span>接口地址：</span> <span>https://api.qichacha.com/TaxData/CreateOrder</span><span>接口地址：</span> <span>https://api.qichacha.com/TaxData/SendCode</span><span>接口地址：</span> <span>https://api.qichacha.com/TaxData/GetData</span></body></html>""",
        encoding="utf-8",
    )
    return index, details


def check_success_flow(tmp: Path) -> None:
    index, details = make_success_fixture(tmp / "success")
    fetched = run([
        sys.executable,
        str(FETCH),
        "--index-html",
        str(index),
        "--detail-dir",
        str(details),
        "--summary-json",
    ])
    summary = json.loads(fetched.stdout)
    assert_true(summary["status"] == "ok", f"unexpected fetch status: {summary}")
    assert_true(summary["apicode_count"] == 2, f"unexpected apicode count: {summary}")
    assert_true(summary["interface_count"] == 5, f"unexpected interface count: {summary}")
    assert_true(summary["warning_count"] == 0, f"unexpected warning count: {summary}")
    assert_true(summary["out"], "success summary must include out path")

    report = tmp / "artifact" / "nested" / "report.md"
    run([sys.executable, str(CHECK), "--repo", str(ROOT), "--docs-json", summary["out"], "--output", str(report)])
    assert_true(report.exists(), "qcc_check.py --output should create parent directories")


def check_warning_flow(tmp: Path) -> None:
    base = tmp / "warning"
    details = base / "details"
    details.mkdir(parents=True)
    index = base / "index.html"
    index.write_text('<script>window.__NUXT__={list:[{ApiCode:"213",Title:"企业年报信息"}]}</script>', encoding="utf-8")

    failed = run([
        sys.executable,
        str(FETCH),
        "--index-html",
        str(index),
        "--detail-dir",
        str(details),
        "--summary-json",
    ], check=False)
    assert_true(failed.returncode == 1, "warning extraction must fail by default")
    failed_summary = json.loads(failed.stdout)
    assert_true(failed_summary["status"] == "failed", f"unexpected failed summary: {failed_summary}")
    assert_true(failed_summary["out"] is None, "failed extraction must not emit a consumable out path")

    allowed = run([
        sys.executable,
        str(FETCH),
        "--index-html",
        str(index),
        "--detail-dir",
        str(details),
        "--summary-json",
        "--allow-warnings",
    ])
    allowed_summary = json.loads(allowed.stdout)
    rejected = run([sys.executable, str(CHECK), "--repo", str(ROOT), "--docs-json", allowed_summary["out"]], check=False)
    assert_true(rejected.returncode == 1, "qcc_check.py must reject warning docs by default")
    assert_true(re.search(r"warnings/errors", rejected.stderr), f"unexpected rejection stderr: {rejected.stderr}")


def main() -> int:
    check_python_files()
    check_evals_schema()
    with tempfile.TemporaryDirectory(prefix="qcc-check-quality-") as tmp_dir:
        tmp = Path(tmp_dir)
        check_success_flow(tmp)
        check_warning_flow(tmp)
    print("qcc-check skill quality checks passed")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
