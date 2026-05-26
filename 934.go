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

type HistoryChattelMortgageCheckGetListReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 页码，默认第1页
	PageIndex string
	// 每页数据条数，默认为10，最大20
	PageSize string
}

type HistoryChattelMortgageCheckGetListResp struct {
	Response[HistoryChattelMortgageCheckGetListRespResult]
}

type HistoryChattelMortgageCheckGetListRespResult struct {
	VerifyResult int64                                                  `json:"VerifyResult"`
	Data         []HistoryChattelMortgageCheckGetListRespResultDataItem `json:"Data"`
}

type HistoryChattelMortgageCheckGetListRespResultDataItem struct {
	ID                string                                                                  `json:"Id"`
	RegisterNo        string                                                                  `json:"RegisterNo"`
	RegisterDate      string                                                                  `json:"RegisterDate"`
	DebtSecuredAmount string                                                                  `json:"DebtSecuredAmount"`
	Status            string                                                                  `json:"Status"`
	PledgeeList       []HistoryChattelMortgageCheckGetListRespResultDataItemPledgeeListItem   `json:"PledgeeList"`
	Pledgor           HistoryChattelMortgageCheckGetListRespResultDataItemPledgor             `json:"Pledgor"`
	FulfillObligation string                                                                  `json:"FulfillObligation"`
	OwnershipList     []HistoryChattelMortgageCheckGetListRespResultDataItemOwnershipListItem `json:"OwnershipList"`
}

type HistoryChattelMortgageCheckGetListRespResultDataItemPledgeeListItem struct {
	Name  string `json:"Name"`
	KeyNo string `json:"KeyNo"`
}

type HistoryChattelMortgageCheckGetListRespResultDataItemPledgor struct {
	Name  string `json:"Name"`
	KeyNo string `json:"KeyNo"`
}

type HistoryChattelMortgageCheckGetListRespResultDataItemOwnershipListItem struct {
	Name  string `json:"Name"`
	KeyNo string `json:"KeyNo"`
}

// HistoryChattelMortgageCheckGetList 历史动产抵押核查 https://openapi.qcc.com/dataApi/934
func (a *Api) HistoryChattelMortgageCheckGetList(ctx context.Context, req *HistoryChattelMortgageCheckGetListReq) (*HistoryChattelMortgageCheckGetListResp, error) {
	var resp HistoryChattelMortgageCheckGetListResp
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

	reply, err := c.SetResult(&resp).Get("/HistoryChattelMortgageCheck/GetList")
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

type HistoryChattelMortgageCheckGetDetailReq struct {
	// 企业的主键
	KeyNo string
	// 动产抵押 Id
	ID string
}

type HistoryChattelMortgageCheckGetDetailResp struct {
	Response[HistoryChattelMortgageCheckGetDetailRespResult]
}

type HistoryChattelMortgageCheckGetDetailRespResult struct {
	Pledge        HistoryChattelMortgageCheckGetDetailRespResultPledge              `json:"Pledge"`
	PledgeeList   []HistoryChattelMortgageCheckGetDetailRespResultPledgeeListItem   `json:"PledgeeList"`
	SecuredClaim  HistoryChattelMortgageCheckGetDetailRespResultSecuredClaim        `json:"SecuredClaim"`
	GuaranteeList []HistoryChattelMortgageCheckGetDetailRespResultGuaranteeListItem `json:"GuaranteeList"`
	Pledgor       HistoryChattelMortgageCheckGetDetailRespResultPledgor             `json:"Pledgor"`
}

type HistoryChattelMortgageCheckGetDetailRespResultPledge struct {
	RegistNo     string `json:"RegistNo"`
	RegistDate   string `json:"RegistDate"`
	RegistOffice string `json:"RegistOffice"`
}

type HistoryChattelMortgageCheckGetDetailRespResultPledgeeListItem struct {
	Name         string `json:"Name"`
	IdentityType string `json:"IdentityType"`
	IdentityNo   string `json:"IdentityNo"`
	KeyNo        string `json:"KeyNo"`
}

type HistoryChattelMortgageCheckGetDetailRespResultSecuredClaim struct {
	Kind              string `json:"Kind"`
	Amount            string `json:"Amount"`
	AssuranceScope    string `json:"AssuranceScope"`
	FulfillObligation string `json:"FulfillObligation"`
	Remark            string `json:"Remark"`
}

type HistoryChattelMortgageCheckGetDetailRespResultGuaranteeListItem struct {
	Name          string                                                                             `json:"Name"`
	Other         string                                                                             `json:"Other"`
	Remark        string                                                                             `json:"Remark"`
	OwnershipList []HistoryChattelMortgageCheckGetDetailRespResultGuaranteeListItemOwnershipListItem `json:"OwnershipList"`
}

type HistoryChattelMortgageCheckGetDetailRespResultGuaranteeListItemOwnershipListItem struct {
	Name  string `json:"Name"`
	KeyNo string `json:"KeyNo"`
}

type HistoryChattelMortgageCheckGetDetailRespResultPledgor struct {
	Name  string `json:"Name"`
	KeyNo string `json:"KeyNo"`
}

// HistoryChattelMortgageCheckGetDetail 动产抵押详情 https://openapi.qcc.com/dataApi/934
func (a *Api) HistoryChattelMortgageCheckGetDetail(ctx context.Context, req *HistoryChattelMortgageCheckGetDetailReq) (*HistoryChattelMortgageCheckGetDetailResp, error) {
	var resp HistoryChattelMortgageCheckGetDetailResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("keyNo", req.KeyNo)
	c.SetQueryParam("id", req.ID)

	reply, err := c.SetResult(&resp).Get("/HistoryChattelMortgageCheck/GetDetail")
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
