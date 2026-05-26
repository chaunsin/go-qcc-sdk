# Consumer Testing Patterns

Use this reference when the user wants to test application code that depends on `go-qcc-sdk`.

## Prefer a narrow app interface

Wrap only the SDK methods the application needs:

```go
type CompanySearcher interface {
	FuzzySearchGetList(context.Context, *api.FuzzySearchGetListReq) (*api.FuzzySearchGetListResp, error)
}

type CompanyService struct {
	qcc CompanySearcher
}
```

This keeps most business tests independent from QCC credentials, network access, paid quota, and SDK internals.

## Fake for business tests

```go
type fakeCompanySearcher struct {
	resp *api.FuzzySearchGetListResp
	err  error
}

func (f fakeCompanySearcher) FuzzySearchGetList(ctx context.Context, req *api.FuzzySearchGetListReq) (*api.FuzzySearchGetListResp, error) {
	return f.resp, f.err
}
```

Use fake responses to test request validation, business mapping, not-found behavior, retries owned by your app, and logging decisions.

## `httptest` for local integration wiring

Use `httptest` when you need to verify the SDK is wired to a local fake server. Assert synthetic headers, query values, and decoded fields; never target `https://api.qichacha.com` from tests.

```go
func TestFuzzySearchWiring(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/FuzzySearch/GetList" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if r.Header.Get("Token") == "" || r.Header.Get("Timespan") == "" {
			t.Fatal("missing auth headers")
		}
		if got := r.URL.Query().Get("key"); got != "test-key" {
			t.Fatalf("unexpected key: %s", got)
		}
		if got := r.URL.Query().Get("searchKey"); got != "企查查科技股份有限公司" {
			t.Fatalf("unexpected searchKey: %s", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"200","Message":"OK","Result":{"Name":"企查查科技股份有限公司","CreditCode":"91320594758983201R"}}`))
	}))
	defer ts.Close()

	client := api.NewClient(&api.Config{
		Key:       "test-key",
		SecretKey: "test-secret",
		BaseURL:   ts.URL,
	}, resty.New())

	resp, err := client.FuzzySearchGetList(context.Background(), &api.FuzzySearchGetListReq{
		SearchKey: "企查查科技股份有限公司",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Result.Name != "企查查科技股份有限公司" {
		t.Fatalf("unexpected name: %s", resp.Result.Name)
	}
}
```

Keep these tests local and deterministic. Assert only synthetic keys, paths, query values, headers, and decoded fields that matter to the application.

## Safety checklist

- Do not run unit tests against `https://api.qichacha.com`.
- Do not record real paid responses as fixtures.
- Use placeholders like `test-key` and `test-secret`.
- Avoid enabling SDK `Debug` in CI logs unless logs are private and scrubbed.
- Make context deadlines explicit so tests and production requests cannot hang indefinitely.
