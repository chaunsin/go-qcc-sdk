# Quickstart Integration

Use this reference when the user wants a fast, correct SDK setup example.

## Install and import

```bash
go get github.com/chaunsin/go-qcc-sdk
```

```go
import api "github.com/chaunsin/go-qcc-sdk"
```

## Minimal client setup

Prefer environment variables or a secret manager for credentials. Validate them before constructing the client because the SDK does not reject empty keys at startup.

```go
key := os.Getenv("QCC_KEY")
secretKey := os.Getenv("QCC_SECRET_KEY")
if key == "" || secretKey == "" {
	return nil, fmt.Errorf("QCC_KEY and QCC_SECRET_KEY must be set")
}

client := api.New(&api.Config{
	Key:       key,
	SecretKey: secretKey,
	Timeout:   10 * time.Second,
})
```

Important `Config` fields:

- `Key` and `SecretKey`: required for authenticated QCC calls; validate before constructing the client.
- `BaseURL`: defaults to `https://api.qichacha.com`; override mainly for tests, local fakes, or controlled custom clients.
- `Location`: optional time zone name used when generating auth timestamps; invalid values can panic during client construction.
- `Timeout`: resty request timeout. A zero value means no timeout.
- `Debug`: useful during local diagnosis, but avoid in production logs because request details may be emitted.
- Cache-related fields exist in `Config`, but do not promise cache behavior unless current SDK code shows an initialized cache path for the method being used.

## Production cautions

- Current `api.New` behavior configures resty with `InsecureSkipVerify`; if strict TLS certificate verification matters, inspect the current SDK and consider `NewClient` with a custom resty client.
- `NewClient` is useful for local fake `BaseURL` tests and for advanced production transport, proxy, logging, or TLS customization.
- Do not ask the agent to make live QCC calls. Provide commands for the user to run locally if they choose to spend API quota.

## Config file setup

For production services, prefer `LoadConfig` so the app can handle file errors, validate secrets, and decide how to expand environment variables before calling `api.New`:

```go
cfg, err := api.LoadConfig("qcc.yaml")
if err != nil {
	return nil, err
}
if cfg.Key == "" || cfg.SecretKey == "" {
	return nil, fmt.Errorf("qcc.yaml must include key and secretKey")
}
client := api.New(cfg)
```

`NewFromFile` also exists, but it panics on load failure, so reserve it for demos or simple scripts.

Example YAML with placeholders, not real credentials:

```yaml
key: "replace-with-qcc-key"
secretKey: "replace-with-qcc-secret-key"
timeout: 10s
debug: false
```

`LoadConfig` decodes the file as-is; it does not automatically expand `${QCC_KEY}` placeholders. If the app needs env expansion, do it in application code before assigning `Config` fields.

## Example call

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

resp, err := client.FuzzySearchGetList(ctx, &api.FuzzySearchGetListReq{
	SearchKey: "企查查科技股份有限公司",
	PageIndex: 1,
})
if err != nil {
	return err
}

fmt.Println(resp.Result.Name, resp.Result.CreditCode)
```

Check the local method's response type before assuming `Result` is a slice. Some SDK methods return a struct, some return a slice, and some include nested data.

## Error handling notes

Most endpoint methods return an error when:

- SDK auth token generation fails.
- The HTTP request fails.
- The HTTP status is not 200.
- QCC returns a provider envelope with `Status != "200"`.

Handle `err` before reading `resp`. For diagnostics, log sanitized context such as method name, request ID/order number when available, and provider message; do not log secrets or full paid responses.
