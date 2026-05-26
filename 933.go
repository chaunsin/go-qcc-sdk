// MIT License
//
// Copyright (c) 2024 chaunsin
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
	"context"
	"fmt"
)

type HistoryAdminPenaltyCheckGetListReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 页码，默认第1页
	PageIndex string
	// 每页数据条数，默认为10，最大20
	PageSize string
}

type HistoryAdminPenaltyCheckGetListResp struct {
	Response[HistoryAdminPenaltyCheckGetListRespResult]
}

type HistoryAdminPenaltyCheckGetListRespResult struct {
	VerifyResult int64                                               `json:"VerifyResult"`
	Data         []HistoryAdminPenaltyCheckGetListRespResultDataItem `json:"Data"`
}

type HistoryAdminPenaltyCheckGetListRespResultDataItem struct {
	ID           string `json:"Id"`
	DocNo        string `json:"DocNo"`
	PunishReason string `json:"PunishReason"`
	PunishResult string `json:"PunishResult"`
	PunishOffice string `json:"PunishOffice"`
	PunishDate   string `json:"PunishDate"`
	Source       string `json:"Source"`
	SourceCode   string `json:"SourceCode"`
	PunishAmt    string `json:"PunishAmt"`
}

// HistoryAdminPenaltyCheckGetList 历史行政处罚核查 https://openapi.qcc.com/dataApi/933
func (a *Api) HistoryAdminPenaltyCheckGetList(ctx context.Context, req *HistoryAdminPenaltyCheckGetListReq) (*HistoryAdminPenaltyCheckGetListResp, error) {
	var resp HistoryAdminPenaltyCheckGetListResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("searchKey", req.SearchKey)
	if req.PageIndex != "" {
		c.SetQueryParam("pageIndex", req.PageIndex)
	}
	if req.PageSize != "" {
		c.SetQueryParam("pageSize", req.PageSize)
	}

	reply, err := c.SetResult(&resp).Get("/HistoryAdminPenaltyCheck/GetList")
	if err != nil {
		return nil, err
	}
	if reply.StatusCode() != 200 {
		return nil, fmt.Errorf("request status code [%v] body: %s", reply.StatusCode(), string(reply.Body()))
	}
	if resp.Status != "200" {
		return nil, fmt.Errorf("err: %+v", resp)
	}
	return &resp, nil
}
