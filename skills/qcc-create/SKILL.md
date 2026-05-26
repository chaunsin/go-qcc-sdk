---
name: qcc-create
description: Create or update Go SDK interfaces for Qichacha QCC OpenAPI docs. Use when the user invokes /qcc:create, qcc:create, /qcc-create, or asks to implement, audit, repair, or refresh one or more go-qcc-sdk APIs from QCC ApiCodes such as 886 or official URLs such as https://openapi.qcc.com/dataApi/886.
---

# QCC Create

## Purpose

Implement or audit this repository's Go SDK interface code from the official Qichacha OpenAPI documentation. Treat `/qcc:create`, `qcc:create`, and `/qcc-create` as equivalent explicit invocations of this skill.

## Workflow

1. Normalize the target.
   - Accept either an official URL like `https://openapi.qcc.com/dataApi/886` or an ApiCode like `886`.
   - Extract each numeric ApiCode, derive the canonical docs URL `https://openapi.qcc.com/dataApi/{apicode}`, and derive the target file `{apicode}.go` at the repository root.
   - For each ApiCode, first determine how many interfaces the docs page contains. Most ApiCodes map to one interface, but some map to multiple interfaces; for example `https://openapi.qcc.com/dataApi/213` contains both `查询公司年报` and `查询公司年报概况`.
   - If the user provided both URL and ApiCode and they disagree, stop and ask for clarification.

2. Choose serial or parallel execution for batched work.
   - If the prompt contains multiple ApiCodes, multiple docs URLs, or multiple clearly isolated interface tasks, ask the user whether to run in parallel or serial unless they already chose.
   - Recommend parallel mode when tasks map to disjoint `{apicode}.go` files or independent docs audits; recommend serial mode when tasks share the same file, endpoint family, public naming decision, or likely conflict.
   - Only spawn subagents after the user explicitly chooses parallel. Give each subagent a disjoint ApiCode/file ownership boundary, the official docs source for that boundary, and an instruction not to edit unrelated files or overwrite other agents' work.
   - Keep one coordinating agent responsible for conflict checks, final validation, and the combined create/update/no-op summary.

3. Read the official interface facts before touching code.
   - Fetch or inspect the QCC docs page and extract: interface address, HTTP method, request format, request headers, request parameters, response fields, response example, pagination, enum meanings, required flags, and important notes.
   - Build a complete request-parameter inventory before implementation. For every documented request parameter, record exact wire name and case, location (`query`, `header`, `body`, or path), required/optional status, type, default, enum/range, and description.
   - When one ApiCode page contains multiple interfaces, build a separate fact record for each interface, including interface name, request path, method, parameters, response shape, and example.
   - Unless the user explicitly selected only one interface from a multi-interface ApiCode, audit/create/update every interface documented under that ApiCode.
   - If the page requires login, dynamic browser rendering, or the user explicitly asks to use `@chrome`, use the available Chrome/browser workflow instead of guessing from memory; get explicit user approval before relying on logged-in sessions, cookies, or account pages, and do not save credentials, cookies, or private docs dumps.
   - If official docs cannot be accessed, ask the user to paste the relevant docs content. Do not invent endpoint paths, parameters, field names, enum values, or response shapes.

4. Inspect the current repository implementation.
   - Read `{apicode}.go` if it exists; also search for `dataApi/{apicode}`, the documented endpoint path, and related method/type names.
   - A single ApiCode can contain multiple endpoint methods; audit the full file and verify every documented interface is present instead of stopping after the first matching method.
   - Load `references/go-qcc-sdk-patterns.md` for repository-specific implementation conventions.

5. Decide whether this is create, update, or no-op.
   - Create when no corresponding implementation exists.
   - Update when an implementation exists but differs from official docs.
   - No-op when the existing implementation matches the official docs and comments are adequate.

6. Implement with minimal, style-preserving changes.
   - Keep the existing root-level `{apicode}.go` layout, license header, `package api`, exported request/response types, `context` and `fmt` usage, `auth()` call, `Token` and `Timespan` headers, `key` query parameter, status-code check, and `resp.Status != "200"` check.
   - Implement every documented request parameter from the inventory. Each parameter must have a corresponding request struct field, exact `SetQueryParam`/`SetHeader`/body mapping, and a field comment when docs provide a description, enum, range, or default.
   - Do not drop optional parameters. Optional strings may be omitted only when empty; optional numeric parameters may be omitted on zero only when zero is not a documented valid value. If zero is valid or unset must be distinguishable, use a pointer field or another explicit representation rather than silently skipping zero.
   - Do not add undocumented business request parameters. The only common parameters allowed outside the docs inventory are the SDK auth/base parameters already used by the repo, such as `key`, `Token`, and `Timespan`.
   - Preserve existing public names when they are already reasonable; only rename exported API when required to match docs or fix a clear error.
   - Ensure query parameter names and response JSON tags exactly match official docs, including letter case.
   - Add or update succinct field comments when official docs provide descriptions, enum values, or units.
   - Do not store real QCC credentials, tokens, cookies, copied secrets, or paid API responses in the repository.

7. Validate and summarize.
   - Run `gofmt` on touched Go files when interface code changes.
   - Add or adjust focused `httptest` coverage when paths, query parameters, or response parsing changed; tests should assert every documented request parameter is sent with the exact wire name and expected value.
   - Before finalizing, compare the request-parameter inventory with the implemented request struct fields and request builder calls, and explicitly resolve any missing, extra, or case-mismatched parameters.
   - Run targeted `go test` when possible; if local networking is blocked, report the sandbox limitation and the exact command to rerun.
   - If parallel mode was used, reconcile each subagent's file ownership, verify no overlapping edits conflict, and merge summaries by ApiCode before reporting.
   - For a create result, summarize added file, methods, request parameters, response structs, and tests.
   - For an update result, summarize changed endpoint paths, methods, parameters, field names or JSON tags, comments, and tests.
   - For a no-op result, summarize what official facts were compared and that no code changes were needed.
