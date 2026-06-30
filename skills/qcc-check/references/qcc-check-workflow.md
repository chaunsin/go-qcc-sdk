# QCC Check Workflow Reference

Use this reference when running `$qcc-check` or validating the audit report. Keep `SKILL.md` concise and load this file only when detailed extraction or report-shape guidance is needed.

## Official Documentation Extraction

Use `https://openapi.qcc.com/dataApi` as the official index. Extract one record per ApiCode:

```json
{
  "apicode": "213",
  "title": "企业年报信息",
  "description": "optional index description",
  "url": "https://openapi.qcc.com/dataApi/213"
}
```

The public index is currently rendered by Nuxt SSR and embeds API records in the HTML payload with keys such as `ApiCode`, `Title`, and `Description`. Do not assume lowercase `apiCode`, and do not treat category-count endpoints as sufficient because they do not provide per-ApiCode interface counts.

The extractor may parse endpoint address strings such as `https://api.qichacha.com/...` from documentation HTML, but an audit must not request that host or any real paid QCC API endpoint.

Prefer the bundled extractor for a full audit:

```bash
python3 skills/qcc-check/scripts/qcc_fetch_docs.py --summary-json --verbose
```

The extractor writes to a unique JSON file in Python's platform temp directory by default. With `--summary-json`, stdout is a one-line machine-readable JSON object:

```json
{"status":"ok","apicode_count":167,"interface_count":194,"warning_count":0,"out":"/path/to/qcc_official_docs.json"}
```

Continue only when `status` is `ok`, `warning_count` is `0`, and `out` is non-empty. If the current runtime needs a stable artifact location, pass `--out <path>` with a path valid for that computer or agent tool. To inspect Python's platform temp directory without fetching network docs, run:

```bash
python3 skills/qcc-check/scripts/qcc_fetch_docs.py --print-temp-dir
```

Then run the comparator:

```bash
python3 skills/qcc-check/scripts/qcc_check.py --repo . --docs-json <official-docs-json>
```

The comparator prints Markdown to stdout by default. Only pass `--output <report-path>` when the user or tool needs a file artifact, and choose a runtime-valid path instead of hardcoding macOS-only locations such as `/private/tmp`.

`qcc_check.py` rejects official docs JSON with top-level `warnings`, record-level `error`/`warnings`, or non-zero `summary.warning_count` by default. Use `--allow-doc-warnings` only for explicit manual inspection; do not use it for a final coverage claim.

The extractor opens each detail page and extracts every documented interface under that ApiCode. A detail record should look like:

```json
{
  "apicode": "213",
  "title": "企业年报信息",
  "url": "https://openapi.qcc.com/dataApi/213",
  "interfaces": [
    {
      "description": "企业年报信息-查询公司年报信息",
      "path": "/AR/GetAnnualReport"
    },
    {
      "description": "企业年报信息-查询公司年报概况",
      "path": "/AR/GetAnnualReportSummary"
    }
  ]
}
```

If the extraction tool cannot preserve `path`, keep `description`; the coverage audit only requires counts and human-readable descriptions.

## Extraction Validation Checklist

Before reporting a final coverage result, verify the official extraction JSON summary:

- `summary.apicode_count` matches the official index count shown by the extractor.
- The JSON has one `docs` record per extracted ApiCode.
- `summary.interface_count` is the sum of all detail-page interface counts.
- `summary.warning_count` is `0` for a normal full audit.
- `warnings` is empty; otherwise stop and report the affected ApiCodes instead of declaring coverage.
- Detail records with `interface_count: 0`, rendered 404 pages, login prompts, dynamic-rendering failures, timeout errors, or tab/API-address count mismatches are blockers unless the user supplies verified replacement docs content.

In restricted approval environments, prefer the single Python extractor command. Avoid ad-hoc shell pipelines such as `curl | python` because approval reviewers evaluate each shell segment separately and failures can look like extraction failures. Avoid hardcoded temp paths in reusable skill docs; use Python's platform temp directory, a user-supplied artifact directory, or a repo-local ignored output path selected for the current runtime. For conservative network environments, rerun with `--workers 1`; the extractor also retries transient HTTP/network failures with exponential backoff by default.

## Expected JSON Input

The helper script accepts either a top-level array of detail records or an object with a `docs` array:

```json
{
  "docs": [
    {
      "ApiCode": "213",
      "title": "企业年报信息",
      "url": "https://openapi.qcc.com/dataApi/213",
      "interfaces": [
        "企业年报信息-查询公司年报信息",
        "企业年报信息-查询公司年报概况"
      ]
    }
  ]
}
```

Supported aliases include `ApiCode`, `apiCode`, `apicode`, `code`, `interfaces`, `endpoints`, `interface_count`, and `interfaceCount`.

Every docs record must include either an explicit `interfaces`/`endpoints` list or an explicit `interface_count`/`interfaceCount` value. Index-only records that only contain a title or description are rejected because the audit must not guess interface counts.

## Local Implementation Rules

- Only scan repository-root files named exactly `{ApiCode}.go`.
- Ignore `api.go`, `types.go`, tests, generated caches, nested folders, and non-numeric Go files.
- Count methods declared with an `*Api` receiver, not files. The scanner strips comments and string literals before counting and skips files whose Go build constraints do not match the current default platform.
- Do not judge implementation details such as query parameters, response structs, JSON tags, comments, or test coverage.

## Report Categories

Generate a Markdown report with these sections:

1. **当前缺失接口**
   - Include ApiCode, official title/description, official interface count, local method count, missing count, and interface descriptions.
   - When the local count is lower than official count, quantity-only comparison cannot prove which exact interface is absent. State that the listed official descriptions are the candidates to inspect.

2. **已经废弃或疑似多余的本地实现**
   - Include local `{ApiCode}.go` files whose ApiCode no longer appears in official docs.
   - Include ApiCodes where local method count is greater than official interface count.
   - Tell the user to choose between preserving compatibility, deleting, or marking deprecated.

3. **已经存在实现**
   - Summarize counts and coverage percentage.
   - Include a few examples only; do not list every implemented interface unless requested.

4. **后续建议**
   - Missing or quantity-mismatch work should be created or refreshed with the repository's `qcc-create` skill.
   - Deprecated candidates require a user decision before deletion.

Also include a short validation footer in the user-facing answer:

- Official source URL and extraction date/time.
- Official ApiCode count and interface count.
- Local ApiCode file count and SDK method count.
- Extractor JSON path and Markdown report path when files were written; mention when the report was printed to stdout instead of saved.
- Extraction warning count.
- Explicit scope reminder: this audit checks only ApiCode/interface-count coverage, not request parameters, response fields, JSON tags, tests, or runtime correctness.

## Skill Maintenance Check

When editing this skill or its bundled scripts, run:

```bash
python3 skills/qcc-check/scripts/qcc_skill_quality.py
```

This validates Python syntax, eval metadata shape, offline extraction of quoted and unquoted `ApiCode` index forms, warning fail-closed behavior, comparator warning rejection, and report output directory creation.
