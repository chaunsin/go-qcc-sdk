package api

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var ctx = context.Background()

func newTestAPI(t *testing.T, handler http.HandlerFunc) (*Api, func()) {
	t.Helper()

	ts := httptest.NewServer(handler)
	api := NewClient(&Config{
		Key:       "test-key",
		SecretKey: "test-secret",
		BaseURL:   ts.URL,
	}, resty.New())

	return api, ts.Close
}

func TestNewInitializesRestyClient(t *testing.T) {
	srv := New(&Config{})
	assert.NotNil(t, srv.GetClient())
}

func TestNewSetsDefaultBaseURL(t *testing.T) {
	api := New(&Config{})
	assert.Equal(t, defaultBaseURL, api.GetClient().BaseURL)
}

func TestNewClientUsesConfiguredBaseURL(t *testing.T) {
	cli := resty.New().SetBaseURL("https://example.invalid")
	api := NewClient(&Config{BaseURL: "http://127.0.0.1:18080"}, cli)
	assert.Equal(t, "http://127.0.0.1:18080", api.GetClient().BaseURL)
}

func TestFuzzySearchGetListUsesExpectedRequestAndParsesResponse(t *testing.T) {
	api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/FuzzySearch/GetList", r.URL.Path)
		assert.NotEmpty(t, r.Header.Get("Token"))
		assert.NotEmpty(t, r.Header.Get("Timespan"))
		assert.Equal(t, "test-key", r.URL.Query().Get("key"))
		assert.Equal(t, "企查查科技股份有限公司", r.URL.Query().Get("searchKey"))
		assert.Equal(t, "1", r.URL.Query().Get("pageIndex"))
		assert.Empty(t, r.URL.Query().Get("pageSize"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"200","Message":"OK","Result":{"KeyNo":"k1","Name":"企查查科技股份有限公司"}}`))
	})
	defer closeServer()

	resp, err := api.FuzzySearchGetList(ctx, &FuzzySearchGetListReq{
		SearchKey: "企查查科技股份有限公司",
		PageIndex: 1,
	})

	assert.NoError(t, err)
	assert.Equal(t, "企查查科技股份有限公司", resp.Result.Name)
}

func TestECIThreeElVerifyGetInfoUsesExpectedPath(t *testing.T) {
	api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/ECIThreeElVerify/GetInfo", r.URL.Path)
		assert.Equal(t, "91320594758983201R", r.URL.Query().Get("creditCode"))
		assert.Equal(t, "企查查科技股份有限公司", r.URL.Query().Get("companyName"))
		assert.Equal(t, "陈德强", r.URL.Query().Get("operName"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"200","Message":"OK","Result":{"VerifyResult":1}}`))
	})
	defer closeServer()

	resp, err := api.ECIThreeElVerifyGetInfo(ctx, &ECIThreeElVerifyGetInfoReq{
		CreditCode:  "91320594758983201R",
		CompanyName: "企查查科技股份有限公司",
		OperName:    "陈德强",
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(1), resp.Result.VerifyResult)
}

func TestECIEmployeeGetListUsesExpectedPath(t *testing.T) {
	api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/ECIEmployee/GetList", r.URL.Path)
		assert.Equal(t, "企查查科技股份有限公司", r.URL.Query().Get("searchKey"))
		assert.Equal(t, "1", r.URL.Query().Get("pageIndex"))
		assert.Equal(t, "10", r.URL.Query().Get("pageSize"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"200","Message":"OK","Result":[{"Name":"张三","Job":"董事"}]}`))
	})
	defer closeServer()

	resp, err := api.ECIEmployeeGetList(ctx, &ECIEmployeeGetListReq{
		SearchKey: "企查查科技股份有限公司",
		PageIndex: 1,
		PageSize:  10,
	})

	assert.NoError(t, err)
	assert.Len(t, resp.Result, 1)
	assert.Equal(t, "张三", resp.Result[0].Name)
}

func TestIPOGetMainIndicatorParsesMainIndicatorResult(t *testing.T) {
	api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/IPO/GetMainIndicator", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"200","Message":"OK","Result":{"VerifyResult":1,"Data":{"ReportDate":["2024"],"PrimaryList":[{"PrimaryDes":"盈利能力","SecondaryList":[{"SecondaryDes":"净利润","SecondaryValueList":["100"]}]}]}}}`))
	})
	defer closeServer()

	resp, err := api.IPOGetMainIndicator(ctx, &IPOGetMainIndicatorReq{SearchKey: "企查查科技股份有限公司"})

	assert.NoError(t, err)
	assert.Equal(t, []string{"2024"}, resp.Result.Data.ReportDate)
	assert.Equal(t, "盈利能力", resp.Result.Data.PrimaryList[0].PrimaryDes)
}

func TestIPOGetIPOExecutiveParsesExecutiveResult(t *testing.T) {
	api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/IPO/GetIPOExecutive", r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"200","Message":"OK","Result":[{"Name":"李四","Sex":"男","Age":"45","Education":"本科","Position":"总经理"}]}`))
	})
	defer closeServer()

	resp, err := api.IPOGetIPOExecutive(ctx, &IPOGetIPOExecutiveReq{StockCode: "688001"})

	assert.NoError(t, err)
	assert.Len(t, resp.Result, 1)
	assert.Equal(t, "总经理", resp.Result[0].Position)
}

func TestCompanyShopInfoGetInfoUsesExpectedPathAndParsesResponse(t *testing.T) {
	api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "/CompanyShopInfo/GetInfo", r.URL.Path)
		assert.NotEmpty(t, r.Header.Get("Token"))
		assert.NotEmpty(t, r.Header.Get("Timespan"))
		assert.Equal(t, "test-key", r.URL.Query().Get("key"))
		assert.Equal(t, "企查查科技股份有限公司", r.URL.Query().Get("searchKey"))
		assert.Equal(t, "2026-05", r.URL.Query().Get("dataMonth"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"200","Message":"OK","Result":{"VerifyResult":1,"Data":{"CompanyShopNum":"1","ShopList":[{"DataMonth":"2026-05","PlatformId":"1","Platform":"淘宝","ShopId":"s1","ShopName":"企查查旗舰店","CreditCode":"91320594758983201R","CompanyName":"企查查科技股份有限公司"}]}}}`))
	})
	defer closeServer()

	resp, err := api.CompanyShopInfoGetInfo(ctx, &CompanyShopInfoGetInfoReq{
		SearchKey: "企查查科技股份有限公司",
		DataMonth: "2026-05",
	})

	assert.NoError(t, err)
	assert.Equal(t, int64(1), resp.Result.VerifyResult)
	assert.Equal(t, "1", resp.Result.Data.CompanyShopNum)
	assert.Equal(t, "淘宝", resp.Result.Data.ShopList[0].Platform)
	assert.Equal(t, "企查查旗舰店", resp.Result.Data.ShopList[0].ShopName)
}
