# go-qcc-sdk Interface Patterns

Use this reference after identifying the target ApiCode and official QCC documentation facts.

## Repository Layout

- Package name is `api`; implementation files live at the repository root as `{apicode}.go`.
- Shared response envelope is `Response[T]` in `types.go` with `Status`, `Message`, `OrderNumber`, `Paging`, and `Result`.
- `api.go` owns config, resty client setup, base URL resolution, and `auth()` token generation.
- Tests use `httptest` and `NewClient` with a test base URL.

## ApiCode and Interface Count

- Treat each QCC ApiCode page as a documentation container, not necessarily a single endpoint.
- First inventory all interfaces on the page, including Chinese name, request path, HTTP method, parameters, response structure, and example.
- Most ApiCodes have one interface, but some have multiple interfaces. For example `dataApi/213` includes `查询公司年报` and `查询公司年报概况`, which should both be represented in `213.go` unless the user scoped the task to one interface.
- When auditing existing code, compare the documented interface count with the implemented method count and report missing, extra, renamed, or stale methods.

## Batched and Parallel Work

- Multiple ApiCodes or docs URLs are good parallel candidates only when each task owns a different `{apicode}.go` file.
- Multiple methods under the same ApiCode should usually be serial or owned by one agent because they share one file and public naming decisions.
- In parallel mode, give every subagent a single ApiCode or explicit file boundary plus the same repository style rules; tell it not to modify unrelated files.
- The coordinator should compare each subagent result against official docs, run formatting/tests once, and produce one combined summary by ApiCode.

## File and Method Style

- Keep the MIT license header and `package api` used by existing interface files.
- Import `context` and `fmt` for endpoint methods unless no formatting is needed.
- Define request, response, and response-result structs before methods.
- Name files by ApiCode, but allow multiple methods in one file when QCC documents multiple endpoints under the same ApiCode.
- Method comments should follow GoDoc shape: `// MethodName Chinese interface name https://openapi.qcc.com/dataApi/{apicode}`.

## Request Pattern

- Use `func (a *Api) Method(ctx context.Context, req *MethodReq) (*MethodResp, error)`.
- Start each call with:
  - `var resp MethodResp`
  - `token, unix, err := a.auth()`
  - wrap auth errors as `fmt.Errorf("auth: %w", err)`.
- Build requests with resty:
  - `SetContext(ctx)`
  - `SetHeader("Token", token)`
  - `SetHeader("Timespan", unix)`
  - `SetQueryParam("key", a.cfg.Key)`
  - endpoint-specific query parameters
  - `SetResult(&resp)`
  - `Get("/Module/Action")` unless official docs explicitly require another HTTP method.
- Required parameters are set directly; optional string parameters are set only when non-empty; optional numeric parameters are set only when greater than zero unless docs define zero as meaningful.
- Convert integers with `fmt.Sprintf("%d", value)` when passing query parameters.

## Type and Field Mapping

- Request struct fields are exported Go names and usually do not need JSON tags because existing code sends query parameters explicitly.
- Response struct fields must include exact official JSON tags, preserving QCC field case such as `json:"KeyNo"` or `json:"VerifyResult"`.
- Use `int64` for documented integer-like fields in existing style; use `string` for dates, IDs, codes, names, percentages, and values unless docs/examples prove a numeric JSON type.
- Preserve official spelling in JSON tags even if Go field names are normalized.
- Add concise comments for fields with official descriptions, enum meanings, units, or special constraints.

## Error Handling and Safety

- Preserve the existing response checks:
  - return transport errors directly
  - non-200 HTTP status returns `fmt.Errorf("request status code [%v] body: %s", ...)`
  - non-`"200"` QCC status returns `fmt.Errorf("err: %+v", resp)`
- Do not call real paid QCC endpoints during implementation or tests.
- Do not commit API keys, secret keys, generated tokens, cookies, or private official docs dumps.
- If official docs cannot be verified, stop and ask for pasted docs content rather than guessing.
