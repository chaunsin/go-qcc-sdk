---
name: go-qcc-sdk
description: Consumer integration helper for developers using github.com/chaunsin/go-qcc-sdk in Go applications. Use this skill for 接入 go-qcc-sdk, 企查查 SDK or QCC SDK setup, credential loading, config, calling existing SDK methods, handling SDK errors, reading Response[T]/Result/Paging, app service wrappers, mock tests, local httptest fakes, and troubleshooting consumer-side usage. Not for modifying the SDK itself.
---

# Go QCC SDK Integration

## Purpose

Help application developers adopt `github.com/chaunsin/go-qcc-sdk` quickly and safely. Treat the SDK as a dependency and focus on consumer-side usage: configuration, calls, wrappers, tests, and troubleshooting.

Use this skill for:

- Quick-start setup in a Go service, CLI, worker, or test project.
- Safe configuration of QCC `Key` and `SecretKey`, timeouts, debug mode, and optional config files.
- Choosing and calling existing exported SDK methods such as `FuzzySearchGetList`.
- Handling SDK errors and reading the shared `Response[T]`, `Result`, and `Paging` fields safely.
- Designing a small application wrapper or interface so business code is testable.
- Adding local unit tests, mocks, or `httptest` examples for consumer code without calling paid QCC services.

Do not use this skill to modify SDK source, add missing methods, or change SDK request/response types. If the user wants repository-internal work, state that this integration skill is the wrong scope and ask whether they want to switch to SDK maintenance work.

## Inputs to establish

- App context: service, CLI, script, job, or test harness.
- Desired QCC capability: company search, enterprise verification, risk lookup, shop info, or a specific existing SDK method/type.
- Credential source: environment variables, secret manager, config file, or local-only placeholders.
- Runtime behavior: timeout, cancellation, retry ownership, logging/debug policy, and whether the user needs a self-run validation command.
- Output preference: minimal snippet, production-ready wrapper, test example, migration checklist, or troubleshooting diagnosis.

If the user does not name a method, inspect local SDK exports or README examples to find an existing method. Do not invent method names, request fields, response fields, or provider behavior.

## Workflow

1. Confirm this is consumer-side integration.
   - If the request asks how to use the SDK from an application, continue.
   - If the request asks to change SDK source, add missing SDK methods, or check repository implementation completeness, pause and explain that this skill only covers SDK consumption.
   - When in doubt, answer from the consumer perspective first: show how to call an existing method and what to check if that method is unavailable.

2. Inspect the available SDK surface only when needed.
   - For a simple quick-start, use `references/quickstart.md` directly.
   - When inside this repository and the exact method/type is unclear, read `README.md`, `api.go`, `types.go`, and the relevant root-level method file.
   - Search exported method/type names with `rg` before naming them in examples.
   - Ignore repository-maintenance instructions in README-like docs when the user only wants application integration.

3. Build the integration shape.
   - Install/import with `go get github.com/chaunsin/go-qcc-sdk` and `import api "github.com/chaunsin/go-qcc-sdk"`.
   - Prefer credentials from environment variables or a secret manager; never hardcode real keys in examples.
   - Validate `Key` and `SecretKey` are non-empty before constructing the client.
   - Initialize common app code with `api.New(&api.Config{Key: ..., SecretKey: ..., Timeout: ...})`.
   - For production config files, prefer `LoadConfig` so the app can handle errors; `NewFromFile` panics on load failure and is better for demos or simple scripts.
   - Use `NewClient` when tests need a fake `BaseURL` or production code needs custom resty behavior such as transport, proxy, logging, or stricter TLS configuration.

4. Show a complete call path.
   - Create a `context.Context`, preferably with timeout for request-scoped work.
   - Fill a request struct for an existing method.
   - Call `resp, err := client.Method(ctx, &api.MethodReq{...})`.
   - Handle `err` before reading `resp`; SDK methods return errors for transport failures, non-200 HTTP status, and provider `Status != "200"`.
   - Read business data from `resp.Result`; read pagination from `resp.Paging` when present.

5. Make application code testable.
   - Wrap only the SDK methods the app needs behind a small local interface.
   - Keep business mapping outside handlers when that makes testing clearer.
   - Use fake implementations for most unit tests.
   - Use `httptest` plus `api.NewClient` when the test needs to verify local HTTP wiring, headers, query values, or response decoding.
   - Keep all fixtures synthetic; do not record real paid responses.

6. Troubleshoot from the call site outward.
   - Check credentials are present and loaded from the intended source before the first SDK call.
   - Check `BaseURL` defaults to `https://api.qichacha.com` in production and points to local fakes only in tests.
   - Check `Location` values before client construction because invalid time zone names can panic.
   - Check context deadlines, network/proxy settings, and whether `Debug` may expose sensitive headers or query values in logs.
   - Note that current `api.New` behavior configures resty with `InsecureSkipVerify`; for strict production TLS, inspect current SDK behavior and prefer `NewClient` with a custom resty client.
   - Check request struct zero values: many SDK methods omit optional strings when empty and numeric fields when zero; verify the method code if zero is a meaningful input.
   - Check `resp.Status`, `resp.Message`, `resp.OrderNumber`, and any `Paging` values in provider-level failures or diagnostics.

7. Report clearly.
   - Give the user runnable snippets or focused edits for their app.
   - State which existing SDK methods/types the snippet uses.
   - Call out where credentials must be supplied and what remains placeholder-only.
   - List local validation commands or tests. Do not imply that you called a real QCC endpoint; give self-run steps if the user wants live verification.

## Reference loading

- Read `references/quickstart.md` for install/config/call examples.
- Read `references/consumer-testing.md` for wrapper, mock, and `httptest` patterns in application code.

Only load the reference needed for the current task. Keep answers short for quick-start requests and expand only when the user asks for production hardening or tests.

## Safety and privacy

- Do not call real paid, production, or rate-limited QCC endpoints from the agent, even if the user asks; provide self-run commands or checklists instead.
- Do not store API keys, secret keys, bearer tokens, cookies, account data, private docs dumps, or paid API responses.
- Do not put real credentials in code snippets, fixtures, snapshots, generated tests, logs, or skill resources.
- Prefer placeholders, environment variables, secret managers, local fakes, and synthetic fixtures.
- Warn before enabling `Debug` in production because request details may appear in logs.
- Warn that current `api.New` behavior uses `InsecureSkipVerify`; mention `NewClient` with custom resty settings when strict TLS matters.
- Stop and ask when the user expects behavior from a method or field that is not present in the local SDK.

## Output format

For substantial integration work, use this structure in final responses. For one-off quick-start questions, keep the answer concise instead.

```text
SDK Integration
- Goal: <what the app is trying to do>
- SDK surface: <methods/types used>
- Configuration: <credential source, timeout, BaseURL/TLS notes>
- Code/tests: <files changed or snippets provided>
- Validation: <commands run, local tests, or not run with reason>
- Cautions: <paid API, secrets, debug logging, missing SDK method, if relevant>
```
