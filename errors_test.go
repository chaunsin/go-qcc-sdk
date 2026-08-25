// MIT License
//
// Copyright (c) 2026 chaunsin
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
//

package api

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAPIError_Error(t *testing.T) {
	e := &APIError{Status: StatusKeyArrears, Message: "当前KEY已欠费", OrderNumber: "ord-1"}
	assert.Contains(t, e.Error(), `status="102"`)
	assert.Contains(t, e.Error(), "当前KEY已欠费")
	assert.Contains(t, e.Error(), "ord-1")
}

func TestResponse_err(t *testing.T) {
	// 成功状态返回 nil
	r := Response[string]{Status: StatusSuccess}
	assert.NoError(t, r.err())

	// 无结果(201)也返回 nil
	r201 := Response[string]{Status: StatusNoResult, Message: "查询无结果"}
	assert.NoError(t, r201.err())

	// 业务错误返回 *APIError，字段完整
	r2 := Response[string]{Status: StatusKeyArrears, Message: "当前KEY已欠费", OrderNumber: "ord-1"}
	err := r2.err()
	require.Error(t, err)
	var apiErr *APIError
	require.ErrorAs(t, err, &apiErr)
	assert.Equal(t, StatusKeyArrears, apiErr.Status)
	assert.Equal(t, "当前KEY已欠费", apiErr.Message)
	assert.Equal(t, "ord-1", apiErr.OrderNumber)
}

func TestEndpoint_NoResult(t *testing.T) {
	api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"201","Message":"查询无结果","OrderNumber":"n1"}`))
	})
	defer closeServer()

	resp, err := api.FuzzySearchGetList(ctx, &FuzzySearchGetListReq{SearchKey: "不存在的企业"})
	require.NoError(t, err)
	assert.Nil(t, resp)
}

func TestEndpoint_BusinessError(t *testing.T) {
	api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"202","Message":"查询参数错误，请检查","OrderNumber":"o202"}`))
	})
	defer closeServer()

	resp, err := api.FuzzySearchGetList(ctx, &FuzzySearchGetListReq{SearchKey: "x"})
	require.Error(t, err)
	assert.Nil(t, resp)
	var apiErr *APIError
	require.ErrorAs(t, err, &apiErr)
	assert.Equal(t, StatusParamError, apiErr.Status)
	assert.Equal(t, "查询参数错误，请检查", apiErr.Message)
	assert.Equal(t, "o202", apiErr.OrderNumber)
}

func TestEndpoint_HTTPError(t *testing.T) {
	api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`oops`))
	})
	defer closeServer()

	resp, err := api.FuzzySearchGetList(ctx, &FuzzySearchGetListReq{SearchKey: "x"})
	require.Error(t, err)
	assert.Nil(t, resp)
	var apiErr *APIError
	assert.False(t, errors.As(err, &apiErr), "HTTP 错误不应是 APIError")
}

func TestEndpoint_Success(t *testing.T) {
	api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"Status":"200","Message":"OK","Result":{"KeyNo":"k1","Name":"企查查科技股份有限公司"}}`))
	})
	defer closeServer()

	resp, err := api.FuzzySearchGetList(ctx, &FuzzySearchGetListReq{SearchKey: "企查查科技股份有限公司"})
	assert.NoError(t, err)
	require.NotNil(t, resp)
	// Status 枚举类型经 JSON 反序列化后应正确解析
	assert.Equal(t, StatusSuccess, resp.Status)
	assert.Equal(t, "企查查科技股份有限公司", resp.Result.Name)
}

// TestStatusConstants 校验 42 个状态码常量与官方文档
// https://openapi.qcc.com/services/after/status 一致，且值互不重复。
func TestStatusConstants(t *testing.T) {
	cases := []struct {
		name     string
		status   Status
		expected string
	}{
		// 有效请求状态码
		{"StatusSuccess", StatusSuccess, "200"},
		{"StatusNoResult", StatusNoResult, "201"},
		{"StatusParamError", StatusParamError, "202"},
		{"StatusProcessing", StatusProcessing, "205"},
		{"StatusItemsExceedLimit", StatusItemsExceedLimit, "207"},
		{"StatusUnsupportedCompanyType", StatusUnsupportedCompanyType, "208"},
		{"StatusCompaniesExceedLimit", StatusCompaniesExceedLimit, "209"},
		{"StatusParamTooShort", StatusParamTooShort, "213"},
		{"StatusUnsupportedKeyword", StatusUnsupportedKeyword, "215"},
		{"StatusShellScanUnsupported", StatusShellScanUnsupported, "218"},
		{"StatusDueDiligenceUnsupported", StatusDueDiligenceUnsupported, "219"},
		{"StatusApiOffline", StatusApiOffline, "105"},
		{"StatusRetryLater", StatusRetryLater, "110"},
		// 无效请求状态码
		{"StatusKeyInvalid", StatusKeyInvalid, "101"},
		{"StatusKeyArrears", StatusKeyArrears, "102"},
		{"StatusKeySuspended", StatusKeySuspended, "103"},
		{"StatusKeyAbnormal", StatusKeyAbnormal, "104"},
		{"StatusTooManyIllegalRequests", StatusTooManyIllegalRequests, "106"},
		{"StatusIpForbiddenOrSignError", StatusIpForbiddenOrSignError, "107"},
		{"StatusTooManyAbnormalRequests", StatusTooManyAbnormalRequests, "108"},
		{"StatusDailyLimitExceeded", StatusDailyLimitExceeded, "109"},
		{"StatusPermissionNotGranted", StatusPermissionNotGranted, "111"},
		{"StatusQuotaExhausted", StatusQuotaExhausted, "112"},
		{"StatusApiDeleted", StatusApiDeleted, "113"},
		{"StatusApiDisabled", StatusApiDisabled, "114"},
		{"StatusAuthInvalidOrExpired", StatusAuthInvalidOrExpired, "115"},
		{"StatusDailyTotalExceeded", StatusDailyTotalExceeded, "116"},
		{"StatusUnsupportedParamCalls", StatusUnsupportedParamCalls, "117"},
		{"StatusUnsupportedCallMethod", StatusUnsupportedCallMethod, "118"},
		{"StatusAccountAbnormal", StatusAccountAbnormal, "119"},
		{"StatusSystemTrafficAbnormal", StatusSystemTrafficAbnormal, "120"},
		{"StatusDataExportRestricted", StatusDataExportRestricted, "121"},
		{"StatusConcurrencyLimitExceeded", StatusConcurrencyLimitExceeded, "122"},
		{"StatusUnverifiedQuotaLimit", StatusUnverifiedQuotaLimit, "123"},
		{"StatusSystemMaintenance", StatusSystemMaintenance, "124"},
		{"StatusQueryParamError", StatusQueryParamError, "125"},
		{"StatusParamTypeInvalid", StatusParamTypeInvalid, "126"},
		{"StatusUnknownError", StatusUnknownError, "199"},
		{"StatusSystemQueryError", StatusSystemQueryError, "203"},
		{"StatusApiNotPurchased", StatusApiNotPurchased, "214"},
		{"StatusQueryParamFuzzy", StatusQueryParamFuzzy, "223"},
		{"StatusQueryParamInvalid", StatusQueryParamInvalid, "224"},
	}
	require.Len(t, cases, 42, "官方状态码应为 42 个（13 个有效请求 + 29 个无效请求）")

	seen := make(map[Status]string, len(cases))
	for _, c := range cases {
		assert.Equal(t, Status(c.expected), c.status, "%s 值与官方文档不一致", c.name)
		if prev, dup := seen[c.status]; dup {
			t.Errorf("状态码 %q 被 %s 与 %s 重复定义", c.status, prev, c.name)
		}
		seen[c.status] = c.name
	}
}

func TestStatus_IsSuccess(t *testing.T) {
	assert.True(t, StatusSuccess.IsSuccess())
	assert.False(t, StatusNoResult.IsSuccess())
	assert.False(t, StatusProcessing.IsSuccess())
	assert.False(t, StatusKeyArrears.IsSuccess())
}

func TestStatus_IsValid(t *testing.T) {
	// 有效请求状态码返回 true
	for _, s := range []Status{
		StatusApiOffline, StatusRetryLater, StatusSuccess, StatusNoResult,
		StatusParamError, StatusProcessing, StatusItemsExceedLimit,
		StatusUnsupportedCompanyType, StatusCompaniesExceedLimit, StatusParamTooShort,
		StatusUnsupportedKeyword, StatusShellScanUnsupported, StatusDueDiligenceUnsupported,
	} {
		assert.Truef(t, s.IsValid(), "有效请求状态 %q 应返回 true", s)
	}
	// 无效请求状态码返回 false
	for _, s := range []Status{
		StatusKeyInvalid, StatusKeyArrears, StatusKeySuspended, StatusKeyAbnormal,
		StatusUnknownError, StatusSystemQueryError, StatusQueryParamInvalid,
	} {
		assert.Falsef(t, s.IsValid(), "无效请求状态 %q 应返回 false", s)
	}
	// 未知状态码返回 false
	assert.False(t, Status("999").IsValid())
}
