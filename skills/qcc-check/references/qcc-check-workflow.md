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

Then open each detail page and extract every documented interface under that ApiCode. A detail record should look like:

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
