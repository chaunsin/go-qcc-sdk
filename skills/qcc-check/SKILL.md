---
name: qcc-check
description: Audit this go-qcc-sdk repository against official Qichacha QCC OpenAPI documentation for interface coverage by ApiCode. Use when the user invokes $qcc-check, /qcc:check, qcc:check, /qcc-check, or asks to compare local {ApiCode}.go implementations with https://openapi.qcc.com/dataApi docs, find missing interfaces, identify deprecated local ApiCodes, or generate a QCC SDK coverage report.
---

# QCC Check

## Purpose

Audit whether this repository implements the same number of interfaces documented by the official QCC OpenAPI pages. Count coverage by ApiCode and endpoint count only; do not review request parameters, response fields, JSON tags, or implementation correctness.

## Workflow

1. Read the official documentation first.
   - Use `https://openapi.qcc.com/dataApi` as the only official index source.
   - Prefer the bundled extractor instead of ad-hoc shell pipelines:
     ```bash
     python3 skills/qcc-check/scripts/qcc_fetch_docs.py --summary-json --verbose
     ```
   - Read the one-line JSON summary from stdout. Continue only when `status` is `ok`, `warning_count` is `0`, and `out` contains the official docs JSON path. Pass `--out <path>` when the current agent/tool needs an explicit workspace or artifact directory.
   - The extractor reads the Nuxt SSR index, then reads every `https://openapi.qcc.com/dataApi/{ApiCode}` page and extracts the documented interface tabs and API addresses.
   - Treat ApiCode as the unique grouping key. One ApiCode can document multiple interfaces, so never assume one ApiCode equals one endpoint.
   - If the extractor reports warnings, missing interface tabs, rendered 404 pages, login requirements, dynamic rendering, or network failures, stop and ask the user for accessible docs content or an approved browser/Chrome workflow. Do not guess interface counts.

2. Parallelize official docs extraction safely when possible.
   - The bundled extractor uses a thread pool by default, which is preferred in restricted shell environments because it avoids brittle `curl | python` pipelines and keeps the network request in one auditable command.
   - Split ApiCodes into independent batches and assign multiple subagents only when the user or current tool policy explicitly allows subagent work.
   - Keep one coordinator agent responsible for merging, local scanning, and final reporting.
   - If subagents are unavailable or approval does not allow them, use the bundled extractor or run the same extraction sequentially and mention the fallback in the report.

3. Scan local SDK implementation.
   - Run `python3 skills/qcc-check/scripts/qcc_check.py --repo <repo> --local-json` to list local numeric Go files and method counts.
   - The scanner only considers root-level files named `{ApiCode}.go`.
   - It counts SDK methods with an `*Api` receiver inside each file; for example, `213.go` can contain two implemented methods.

4. Compare official docs with local implementation.
   - Store official extraction results as JSON, then run:
     ```bash
     python3 skills/qcc-check/scripts/qcc_check.py --repo <repo> --docs-json <official-docs-json>
     ```
   - Let the report print to stdout unless the user or current tool needs a file artifact; if so, pass `--output <report-path>` using a path valid for the current computer or agent runtime.
   - `qcc_check.py` rejects docs JSON that contains extractor warnings/errors by default. Use `--allow-doc-warnings` only for explicit manual inspection, not for a final coverage claim.
   - Load `references/qcc-check-workflow.md` when preparing or validating the JSON shape and report content.
   - Validate the extractor summary before trusting the comparison: official ApiCode count, detail record count, official interface count, and warning count must be reported; `warning_count` should be `0` for a normal full audit.
   - Classify results as missing, deprecated/suspicious, and implemented-summary. Do not list every implemented interface unless the user asks.

5. Report next actions.
   - For missing or out-of-date interfaces, recommend using this repository's `qcc-create` skill with the target ApiCode.
   - For official docs that no longer exist but local code exists, ask the user to choose whether to keep for compatibility, delete, or mark as deprecated.
   - Do not modify SDK code as part of this audit skill.

6. Validate skill tooling after editing this skill.
   - Run the bundled quality check when changing `qcc-check` instructions or scripts:
     ```bash
     python3 skills/qcc-check/scripts/qcc_skill_quality.py
     ```

## Safety

- Only access official documentation pages at `https://openapi.qcc.com/dataApi` and `https://openapi.qcc.com/dataApi/{ApiCode}`.
- Do not call real paid QCC API endpoints during an audit, including `https://api.qichacha.com/...`; endpoint addresses from that host may be parsed as strings from docs pages but must not be requested.
- Do not save API keys, secret keys, generated tokens, cookies, account data, paid responses, or private documentation dumps.
- Use single-command Python scripts for network extraction in approval-restricted environments; avoid shell pipelines that obscure which segment needs network approval.
- Do not hardcode macOS-only temp paths such as `/private/tmp` in skill instructions or reusable reports. Prefer the extractor default temp path, a user-provided artifact path, or a repo-local ignored output path chosen for the current runtime.
- If official documentation cannot be verified, stop and ask for source material instead of inventing endpoints or counts.
