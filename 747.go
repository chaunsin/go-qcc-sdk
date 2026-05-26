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

type ChattelMortgageCheckGetListReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 页码，默认第1页
	PageIndex string
	// 每页数据条数，默认为10，最大20
	PageSize string
}

type ChattelMortgageCheckGetListResp struct {
	Response[ChattelMortgageCheckGetListRespResult]
}

type ChattelMortgageCheckGetListRespResult struct {
	VerifyResult int64                                           `json:"VerifyResult"`
	Data         []ChattelMortgageCheckGetListRespResultDataItem `json:"Data"`
}

type ChattelMortgageCheckGetListRespResultDataItem struct {
	RegisterNo        string                                                  `json:"RegisterNo"`
	RegisterDate      string                                                  `json:"RegisterDate"`
	RegisterOffice    string                                                  `json:"RegisterOffice"`
	DebtSecuredAmount string                                                  `json:"DebtSecuredAmount"`
	Status            string                                                  `json:"Status"`
	Detail            ChattelMortgageCheckGetListRespResultDataItemDetail     `json:"Detail"`
	DebtorInfo        ChattelMortgageCheckGetListRespResultDataItemDebtorInfo `json:"DebtorInfo"`
}

type ChattelMortgageCheckGetListRespResultDataItemDetail struct {
	Pledge        ChattelMortgageCheckGetListRespResultDataItemDetailPledge              `json:"Pledge"`
	PledgeeList   []ChattelMortgageCheckGetListRespResultDataItemDetailPledgeeListItem   `json:"PledgeeList"`
	SecuredClaim  ChattelMortgageCheckGetListRespResultDataItemDetailSecuredClaim        `json:"SecuredClaim"`
	GuaranteeList []ChattelMortgageCheckGetListRespResultDataItemDetailGuaranteeListItem `json:"GuaranteeList"`
	CancelInfo    ChattelMortgageCheckGetListRespResultDataItemDetailCancelInfo          `json:"CancelInfo"`
	ChangeList    []ChattelMortgageCheckGetListRespResultDataItemDetailChangeListItem    `json:"ChangeList"`
}

type ChattelMortgageCheckGetListRespResultDataItemDetailPledge struct {
	RegistNo     string `json:"RegistNo"`
	RegistDate   string `json:"RegistDate"`
	RegistOffice string `json:"RegistOffice"`
}

type ChattelMortgageCheckGetListRespResultDataItemDetailPledgeeListItem struct {
	Name         string `json:"Name"`
	IdentityType string `json:"IdentityType"`
	IdentityNo   string `json:"IdentityNo"`
}

type ChattelMortgageCheckGetListRespResultDataItemDetailSecuredClaim struct {
	Kind              string `json:"Kind"`
	Amount            string `json:"Amount"`
	AssuranceScope    string `json:"AssuranceScope"`
	FulfillObligation string `json:"FulfillObligation"`
	Remark            string `json:"Remark"`
}

type ChattelMortgageCheckGetListRespResultDataItemDetailGuaranteeListItem struct {
	Name      string                                                                              `json:"Name"`
	Ownership string                                                                              `json:"Ownership"`
	Other     string                                                                              `json:"Other"`
	Remark    string                                                                              `json:"Remark"`
	KeyNoList []ChattelMortgageCheckGetListRespResultDataItemDetailGuaranteeListItemKeyNoListItem `json:"KeyNoList"`
}

type ChattelMortgageCheckGetListRespResultDataItemDetailGuaranteeListItemKeyNoListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

type ChattelMortgageCheckGetListRespResultDataItemDetailCancelInfo struct {
	CancelDate   string `json:"CancelDate"`
	CancelReason string `json:"CancelReason"`
}

type ChattelMortgageCheckGetListRespResultDataItemDetailChangeListItem struct {
	ChangeDate    string `json:"ChangeDate"`
	ChangeContent string `json:"ChangeContent"`
}

type ChattelMortgageCheckGetListRespResultDataItemDebtorInfo struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

// ChattelMortgageCheckGetList 动产抵押核查 https://openapi.qcc.com/dataApi/747
func (a *Api) ChattelMortgageCheckGetList(ctx context.Context, req *ChattelMortgageCheckGetListReq) (*ChattelMortgageCheckGetListResp, error) {
	var resp ChattelMortgageCheckGetListResp
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
	if req.PageIndex != "" {
		c.SetQueryParam("pageIndex", req.PageIndex)
	}
	if req.PageSize != "" {
		c.SetQueryParam("pageSize", req.PageSize)
	}

	reply, err := c.SetResult(&resp).Get("/ChattelMortgageCheck/GetList")
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
