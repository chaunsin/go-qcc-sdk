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

type LandMergeCheckGetListReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 页码，默认第1页
	PageIndex string
	// 每页数据条数，默认为10，最大20
	PageSize string
}

type LandMergeCheckGetListResp struct {
	Response[LandMergeCheckGetListRespResult]
}

type LandMergeCheckGetListRespResult struct {
	VerifyResult int64                                     `json:"VerifyResult"`
	Data         []LandMergeCheckGetListRespResultDataItem `json:"Data"`
}

type LandMergeCheckGetListRespResultDataItem struct {
	LandPurID   string `json:"LandPurId"`
	Address     string `json:"Address"`
	Area        string `json:"Area"`
	TradePrice  string `json:"TradePrice"`
	LandUse     string `json:"LandUse"`
	PublishGov  string `json:"PublishGov"`
	Stage       string `json:"Stage"`
	PublishDate string `json:"PublishDate"`
	LandPubID   string `json:"LandPubId"`
}

// LandMergeCheckGetList 国有土地受让 https://openapi.qcc.com/dataApi/989
func (a *Api) LandMergeCheckGetList(ctx context.Context, req *LandMergeCheckGetListReq) (*LandMergeCheckGetListResp, error) {
	var resp LandMergeCheckGetListResp
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

	reply, err := c.SetResult(&resp).Get("/LandMergeCheck/GetList")
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

type LandMergeCheckGetPurchaseDetailReq struct {
	// 购地信息 Id
	LandPurID string
}

type LandMergeCheckGetPurchaseDetailResp struct {
	Response[LandMergeCheckGetPurchaseDetailRespResult]
}

type LandMergeCheckGetPurchaseDetailRespResult struct {
	ID              string `json:"Id"`
	AdminArea       string `json:"AdminArea"`
	ElecSuNum       string `json:"ElecSuNum"`
	ProjectName     string `json:"ProjectName"`
	Address         string `json:"Address"`
	Area            string `json:"Area"`
	LandSource      string `json:"LandSource"`
	LandUse         string `json:"LandUse"`
	SupplyWay       string `json:"SupplyWay"`
	LandUseYears    string `json:"LandUseYears"`
	Industry        string `json:"Industry"`
	LandLevel       string `json:"LandLevel"`
	TransAmt        string `json:"TransAmt"`
	LandHolder      any    `json:"LandHolder"`
	AgreeLandDate   string `json:"AgreeLandDate"`
	AgreeStartDate  string `json:"AgreeStartDate"`
	AgreeEndDate    string `json:"AgreeEndDate"`
	ActualStartDate string `json:"ActualStartDate"`
	ActualEndDate   string `json:"ActualEndDate"`
	ApprovalUnit    string `json:"ApprovalUnit"`
	SignDate        string `json:"SignDate"`
	PaymentRound    string `json:"PaymentRound"`
	AgreePayDate    string `json:"AgreePayDate"`
	AgreePayAmt     string `json:"AgreePayAmt"`
	Remarks         string `json:"Remarks"`
	AgreeRateMin    string `json:"AgreeRateMin"`
	AgreeRateMax    string `json:"AgreeRateMax"`
}

// LandMergeCheckGetPurchaseDetail 购地信息详情 https://openapi.qcc.com/dataApi/989
func (a *Api) LandMergeCheckGetPurchaseDetail(ctx context.Context, req *LandMergeCheckGetPurchaseDetailReq) (*LandMergeCheckGetPurchaseDetailResp, error) {
	var resp LandMergeCheckGetPurchaseDetailResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("landPurId", req.LandPurID)

	reply, err := c.SetResult(&resp).Get("/LandMergeCheck/GetPurchaseDetail")
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

type LandMergeCheckGetPublishDetailReq struct {
	// 地块公示 ID
	LandPubID string
}

type LandMergeCheckGetPublishDetailResp struct {
	Response[LandMergeCheckGetPublishDetailRespResult]
}

type LandMergeCheckGetPublishDetailRespResult struct {
	ID            string `json:"Id"`
	LandNo        string `json:"LandNo"`
	Address       string `json:"Address"`
	Acreage       string `json:"Acreage"`
	Purpose       string `json:"Purpose"`
	TradePrice    string `json:"TradePrice"`
	SellYears     string `json:"SellYears"`
	ProjectName   string `json:"ProjectName"`
	AssigneeUnit  any    `json:"AssigneeUnit"`
	Remarks       string `json:"Remarks"`
	PublishTerm   string `json:"PublishTerm"`
	PublishGov    string `json:"PublishGov"`
	PublishDate   string `json:"PublishDate"`
	Explains      string `json:"Explains"`
	ContactUnit   string `json:"ContactUnit"`
	UnitAddress   string `json:"UnitAddress"`
	PostCode      string `json:"PostCode"`
	ContactNumber string `json:"ContactNumber"`
	ContactPerson string `json:"ContactPerson"`
	Email         string `json:"Email"`
}

// LandMergeCheckGetPublishDetail 地块公示详情 https://openapi.qcc.com/dataApi/989
func (a *Api) LandMergeCheckGetPublishDetail(ctx context.Context, req *LandMergeCheckGetPublishDetailReq) (*LandMergeCheckGetPublishDetailResp, error) {
	var resp LandMergeCheckGetPublishDetailResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("landPubId", req.LandPubID)

	reply, err := c.SetResult(&resp).Get("/LandMergeCheck/GetPublishDetail")
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
