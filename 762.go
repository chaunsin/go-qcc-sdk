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

type OffFilingCheckGetListReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
}

type OffFilingCheckGetListResp struct {
	Response[OffFilingCheckGetListRespResult]
}

type OffFilingCheckGetListRespResult struct {
	VerifyResult int64                                     `json:"VerifyResult"`
	Data         []OffFilingCheckGetListRespResultDataItem `json:"Data"`
}

type OffFilingCheckGetListRespResultDataItem struct {
	ID           string                                        `json:"Id"`
	NoticeStatus string                                        `json:"NoticeStatus"`
	Detail       OffFilingCheckGetListRespResultDataItemDetail `json:"Detail"`
}

type OffFilingCheckGetListRespResultDataItemDetail struct {
	LiqReferenceInfo OffFilingCheckGetListRespResultDataItemDetailLiqReferenceInfo `json:"LiqReferenceInfo"`
	CreditorAnnoInfo OffFilingCheckGetListRespResultDataItemDetailCreditorAnnoInfo `json:"CreditorAnnoInfo"`
}

type OffFilingCheckGetListRespResultDataItemDetailLiqReferenceInfo struct {
	CompanyName   string `json:"CompanyName"`
	CreditCode    string `json:"CreditCode"`
	BelongOrg     string `json:"BelongOrg"`
	ReferenceDate string `json:"ReferenceDate"`
	StartDate     string `json:"StartDate"`
	CancelReason  string `json:"CancelReason"`
	LiqAddress    string `json:"LiqAddress"`
	LiqTelNo      string `json:"LiqTelNo"`
	LiqLeader     string `json:"LiqLeader"`
	LiqMember     string `json:"LiqMember"`
}

type OffFilingCheckGetListRespResultDataItemDetailCreditorAnnoInfo struct {
	CompanyName    string `json:"CompanyName"`
	CreditCode     string `json:"CreditCode"`
	BelongOrg      string `json:"BelongOrg"`
	AnnoDate       string `json:"AnnoDate"`
	AnnoContent    string `json:"AnnoContent"`
	DeclareMember  string `json:"DeclareMember"`
	DeclareTelNo   string `json:"DeclareTelNo"`
	DeclareAddress string `json:"DeclareAddress"`
}

// OffFilingCheckGetList 注销备案核查 https://openapi.qcc.com/dataApi/762
func (a *Api) OffFilingCheckGetList(ctx context.Context, req *OffFilingCheckGetListReq) (*OffFilingCheckGetListResp, error) {
	var resp OffFilingCheckGetListResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("searchKey", req.SearchKey)

	reply, err := c.SetResult(&resp).Get("/OffFilingCheck/GetList")
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
