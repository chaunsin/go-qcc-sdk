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

import "fmt"

// APIError 表示企查查开放平台业务错误（HTTP 200 但响应信封 Status != StatusSuccess）。
// auth 失败、网络传输错误、HTTP 非 200 不属于 APIError，保持原有 error 形式。
//
// 调用方可通过 errors.As 识别业务错误并按状态码分支处理：
//
//	var apiErr *api.APIError
//	if errors.As(err, &apiErr) {
//		switch apiErr.Status {
//		case api.StatusKeyArrears: // KEY 已欠费
//			// ... 业务处理
//		}
//	}
type APIError struct {
	// Status 企查查业务状态码，见 Status* 常量
	Status Status
	// Message 官方错误描述
	Message string
	// OrderNumber 请求流水号，用于向官方排查问题
	OrderNumber string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("qcc: status=%q message=%q orderNumber=%q", e.Status, e.Message, e.OrderNumber)
}

// Status 表示企查查开放平台请求状态码（响应信封中的 Status 字段）。
// 完整定义见 https://openapi.qcc.com/services/after/status
type Status string

// IsSuccess 判断是否为成功状态（StatusSuccess，200）。
func (s Status) IsSuccess() bool {
	return s == StatusSuccess
}

// IsValid 判断是否为"有效请求"状态码（官方归类，含 200/201/205 等），
// 用于与 KEY 欠费、参数错误等"无效请求"状态相区分。
func (s Status) IsValid() bool {
	switch s {
	case StatusApiOffline, StatusRetryLater,
		StatusSuccess, StatusNoResult, StatusParamError, StatusProcessing,
		StatusItemsExceedLimit, StatusUnsupportedCompanyType, StatusCompaniesExceedLimit,
		StatusParamTooShort, StatusUnsupportedKeyword, StatusShellScanUnsupported,
		StatusDueDiligenceUnsupported:
		return true
	default:
		return false
	}
}

// 有效请求状态码（按状态码数值升序）
const (
	StatusApiOffline              Status = "105" // 接口已下线停用
	StatusRetryLater              Status = "110" // 当前相同查询连续出错，请等2小时后重试
	StatusSuccess                 Status = "200" // 查询成功
	StatusNoResult                Status = "201" // 查询无结果
	StatusParamError              Status = "202" // 查询参数错误，请检查
	StatusProcessing              Status = "205" // 等待处理中
	StatusItemsExceedLimit        Status = "207" // 请求数据的条目数超过上限（5000）
	StatusUnsupportedCompanyType  Status = "208" // 此接口不支持此公司类型查询
	StatusCompaniesExceedLimit    Status = "209" // 企业数量超过上限
	StatusParamTooShort           Status = "213" // 参数长度不能小于2
	StatusUnsupportedKeyword      Status = "215" // 不支持的查询关键字
	StatusShellScanUnsupported    Status = "218" // 该企业暂不支持空壳扫描
	StatusDueDiligenceUnsupported Status = "219" // 该企业暂不支持准入尽调
)

// 无效请求状态码。
const (
	StatusKeyInvalid               Status = "101" // 当前的KEY无效或者还未生效中
	StatusKeyArrears               Status = "102" // 当前KEY已欠费
	StatusKeySuspended             Status = "103" // 当前KEY被暂停使用
	StatusKeyAbnormal              Status = "104" // 请求KEY异常，请联系管理员
	StatusTooManyIllegalRequests   Status = "106" // 非法请求过多，请联系管理员
	StatusIpForbiddenOrSignError   Status = "107" // 被禁止的IP或者签名错误
	StatusTooManyAbnormalRequests  Status = "108" // 异常请求过多，请联系管理员
	StatusDailyLimitExceeded       Status = "109" // 请求超过每日系统限制
	StatusPermissionNotGranted     Status = "111" // 接口权限未开通，请联系管理员
	StatusQuotaExhausted           Status = "112" // 您的账号剩余使用量已不足或已过期
	StatusApiDeleted               Status = "113" // 当前接口已被删除，请重新申请
	StatusApiDisabled              Status = "114" // 当前接口已被禁用，请联系管理员
	StatusAuthInvalidOrExpired     Status = "115" // 身份验证错误或者已过期
	StatusDailyTotalExceeded       Status = "116" // 请求超过每日（服务期）调用总量限制
	StatusUnsupportedParamCalls    Status = "117" // 当前不支持的请求参数调用量过多
	StatusUnsupportedCallMethod    Status = "118" // 当前接口不支持此方式的调用
	StatusAccountAbnormal          Status = "119" // 您的账号出现异常，请联系管理员
	StatusSystemTrafficAbnormal    Status = "120" // 系统流量异常，请稍后再试
	StatusDataExportRestricted     Status = "121" // 数据不能出境
	StatusConcurrencyLimitExceeded Status = "122" // 请求超过系统并发限制
	StatusUnverifiedQuotaLimit     Status = "123" // 您的请求已达未认证权益上限，请及时认证
	StatusSystemMaintenance        Status = "124" // 系统维护中，暂时无法下单
	StatusQueryParamError          Status = "125" // 查询参数错误
	StatusParamTypeInvalid         Status = "126" // 查询参数类型或结构有误
	StatusUnknownError             Status = "199" // 系统未知错误，请联系技术客服
	StatusSystemQueryError         Status = "203" // 系统查询有异常，请联系技术人员
	StatusApiNotPurchased          Status = "214" // 您还未购买过该接口，请先购买
	StatusQueryParamFuzzy          Status = "223" // 查询参数模糊，无法获取结果
	StatusQueryParamInvalid        Status = "224" // 查询参数无效
)
