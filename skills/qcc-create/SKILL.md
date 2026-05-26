---
name: qcc-create
description: Create or update Go SDK interfaces for Qichacha QCC OpenAPI docs. Use when the user invokes /qcc:create, qcc:create, /qcc-create, or asks to implement, audit, repair, or refresh a go-qcc-sdk API from a QCC ApiCode such as 886 or an official URL such as https://openapi.qcc.com/dataApi/886.
---

# QCC Create

## Purpose

Implement or audit this repository's Go SDK interface code from the official Qichacha OpenAPI documentation. Treat `/qcc:create`, `qcc:create`, and `/qcc-create` as equivalent explicit invocations of this skill.

## Workflow

1. Normalize the target.
   - Accept either an official URL like `https://openapi.qcc.com/dataApi/886` or an ApiCode like `886`.
   - Extract the numeric ApiCode, derive the canonical docs URL `https://openapi.qcc.com/dataApi/{apicode}`, and derive the target file `{apicode}.go` at the repository root.
   - If the user provided both URL and ApiCode and they disagree, stop and ask for clarification.

2. Read the official interface facts before touching code.
   - Fetch or inspect the QCC docs page and extract: interface address, HTTP method, request format, request headers, request parameters, response fields, response example, pagination, enum meanings, required flags, and important notes.
   - If the page requires login, dynamic browser rendering, or the user explicitly asks to use `@chrome`, use the available Chrome/browser workflow instead of guessing from memory; get explicit user approval before relying on logged-in sessions, cookies, or account pages, and do not save credentials, cookies, or private docs dumps.
   - If official docs cannot be accessed, ask the user to paste the relevant docs content. Do not invent endpoint paths, parameters, field names, enum values, or response shapes.

3. Inspect the current repository implementation.
   - Read `{apicode}.go` if it exists; also search for `dataApi/{apicode}`, the documented endpoint path, and related method/type names.
   - A single ApiCode can contain multiple endpoint methods; audit the full file instead of assuming a one-to-one file-to-method mapping.
   - Load `references/go-qcc-sdk-patterns.md` for repository-specific implementation conventions.

4. Decide whether this is create, update, or no-op.
   - Create when no corresponding implementation exists.
   - Update when an implementation exists but differs from official docs.
   - No-op when the existing implementation matches the official docs and comments are adequate.

5. Implement with minimal, style-preserving changes.
   - Keep the existing root-level `{apicode}.go` layout, license header, `package api`, exported request/response types, `context` and `fmt` usage, `auth()` call, `Token` and `Timespan` headers, `key` query parameter, status-code check, and `resp.Status != "200"` check.
   - Preserve existing public names when they are already reasonable; only rename exported API when required to match docs or fix a clear error.
   - Ensure query parameter names and response JSON tags exactly match official docs, including letter case.
   - Add or update succinct field comments when official docs provide descriptions, enum values, or units.
   - Do not store real QCC credentials, tokens, cookies, copied secrets, or paid API responses in the repository.

6. Validate and summarize.
   - Run `gofmt` on touched Go files when interface code changes.
   - Add or adjust focused `httptest` coverage when paths, query parameters, or response parsing changed.
   - Run targeted `go test` when possible; if local networking is blocked, report the sandbox limitation and the exact command to rerun.
   - For a create result, summarize added file, methods, request parameters, response structs, and tests.
   - For an update result, summarize changed endpoint paths, methods, parameters, field names or JSON tags, comments, and tests.
   - For a no-op result, summarize what official facts were compared and that no code changes were needed.
