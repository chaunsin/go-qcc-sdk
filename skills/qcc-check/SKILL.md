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
   - For every ApiCode in the index, read `https://openapi.qcc.com/dataApi/{ApiCode}` and extract the list of documented interfaces under that ApiCode.
   - Treat ApiCode as the unique grouping key. One ApiCode can document multiple interfaces, so never assume one ApiCode equals one endpoint.
   - If the page requires login, dynamic rendering, or cannot be fetched, ask the user to provide the accessible docs content or use an approved browser/Chrome workflow. Do not guess interface counts.

2. Parallelize official docs extraction when possible.
   - Split ApiCodes into independent batches and assign multiple subagents to extract detail-page facts as structured JSON.
   - Keep one coordinator agent responsible for merging, local scanning, and final reporting.
   - If subagents are unavailable or approval does not allow them, run the same extraction sequentially and mention the fallback in the report.

3. Scan local SDK implementation.
   - Run `python3 skills/qcc-check/scripts/qcc_check.py --repo <repo> --local-json` to list local numeric Go files and method counts.
   - The scanner only considers root-level files named `{ApiCode}.go`.
   - It counts SDK methods with an `*Api` receiver inside each file; for example, `213.go` can contain two implemented methods.

4. Compare official docs with local implementation.
   - Store official extraction results as JSON, then run:
     ```bash
     python3 skills/qcc-check/scripts/qcc_check.py --repo <repo> --docs-json <official-docs.json>
     ```
   - Load `references/qcc-check-workflow.md` when preparing or validating the JSON shape and report content.
   - Classify results as missing, deprecated/suspicious, and implemented-summary. Do not list every implemented interface unless the user asks.

5. Report next actions.
   - For missing or out-of-date interfaces, recommend using this repository's `qcc-create` skill with the target ApiCode.
   - For official docs that no longer exist but local code exists, ask the user to choose whether to keep for compatibility, delete, or mark as deprecated.
   - Do not modify SDK code as part of this audit skill.

## Safety

- Only access official documentation pages under `https://openapi.qcc.com/dataApi`.
- Do not call real paid QCC API endpoints during an audit.
- Do not save API keys, secret keys, generated tokens, cookies, account data, paid responses, or private documentation dumps.
- If official documentation cannot be verified, stop and ask for source material instead of inventing endpoints or counts.
