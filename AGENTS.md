# Repository Guide

## Repository Shape

- Go module: `github.com/chaunsin/go-qcc-sdk`; package name is `api`.
- Most SDK code lives at the repository root.
- `api.go` contains client setup, config loading, base URL handling, and QCC auth token generation.
- `types.go` contains shared response envelope types such as `Response[T]` and `Paging`.
- Each QCC ApiCode is implemented in a root-level `{ApiCode}.go` file. One ApiCode file may contain multiple endpoint methods.
- `skills/qcc-check` and `skills/qcc-create` are repo-local AI-agent skills. `.agents/skills` and `.claude/skills` symlink to `skills`.

## Agent Skill Triggers

- Use `qcc-check` for `/qcc:check`, `qcc:check`, `/qcc-check`, `$qcc-check`, or requests to compare local ApiCode implementations with official QCC OpenAPI docs.
- Use `qcc-create` for `/qcc:create`, `qcc:create`, `/qcc-create`, or requests to create, audit, repair, or refresh one or more QCC ApiCodes.
- For QCC interface work, official docs under `https://openapi.qcc.com/dataApi` are the source of truth. If docs cannot be accessed, ask the user for the relevant docs content instead of guessing.
- Never call real paid QCC API endpoints. Never store real QCC keys, secret keys, generated tokens, cookies, account data, private docs dumps, or paid API responses.

## QCC Endpoint Implementation Rules

- Keep endpoint implementations in root-level `{ApiCode}.go` files with the MIT license header and `package api`.
- Keep request structs, response structs, result structs, and endpoint methods together in the same ApiCode file.
- Preserve existing exported method and type names unless official docs require a change or the current name is clearly wrong.
- Endpoint methods should use `func (a *Api) Method(ctx context.Context, req *MethodReq) (*MethodResp, error)`.
- Use `a.auth()` and wrap auth failures as `fmt.Errorf("auth: %w", err)`.
- Build resty requests with `SetContext(ctx)`, `Token` and `Timespan` headers, shared `key` query parameter, documented endpoint parameters, and `SetResult(&resp)`.
- Prefer fluent resty chaining for request setup: include all unconditional endpoint query parameters in the initial `a.cli.R()` chain instead of assigning the request and then calling unconditional `c.SetQueryParam(...)` on later lines. Use separate calls only when a parameter is conditional, such as omitting an optional empty value.
- Preserve current error handling: return transport errors directly, report non-200 HTTP status with response body, and return an error when `resp.Status != "200"`.
- Request parameter wire names and response JSON tags must exactly match official docs, including casing and spelling.
- Send every documented request parameter unless official docs mark it inapplicable or the user explicitly scopes it out.
- Omit optional strings when empty. For optional numeric fields, use pointer or explicit unset marker unless zero is confirmed not to be a meaningful documented value.
- Prefer `int64` for integer-like response fields in this repository's existing style; prefer `string` for dates, IDs, codes, names, percentages, money-like values, and ambiguous examples unless docs prove a numeric JSON type.
- Do not use `json.RawMessage` as an SDK interface response/result field. Model documented response payloads with concrete structs, slices, and typed fields instead.
- Method comments should be GoDoc-style and include the official docs URL, for example `https://openapi.qcc.com/dataApi/{ApiCode}`.

## Tests And Validation

- Run `gofmt -w <touched .go files>` after editing Go files.
- Run `go test ./...` before finalizing Go changes when possible.
- Tests should use `httptest` and `NewClient` with a configured `BaseURL`; do not hit real QCC services.
- When endpoint paths or parameters change, add or update focused assertions for the exact request path, method, headers, shared `key` parameter, and every documented endpoint-specific parameter.
- Use `api_response_schema_test.go` for response JSON tag/schema regression coverage.
- Use `api_generated_coverage_test.go` for broad generated endpoint request coverage.
- For coverage audits, the local scanner command is `python3 skills/qcc-check/scripts/qcc_check.py --repo . --local-json`.

## Working Practices

- Before SDK interface changes, inspect the current ApiCode file, same-area tests, and relevant repo-local skill docs.
- Treat each QCC ApiCode page as a container that may document multiple interfaces; verify the documented interface count before deciding a file is complete.
- Keep patches minimal and behavior-preserving unless the user asks for a broader refactor.
- Do not overwrite unrelated local changes. Check `git status --short` before and after substantial edits.
- If the user asks whether an ApiCode is missing, stale, or safe to remove, compare official docs, local numeric files, and implemented `*Api` method counts before answering.
