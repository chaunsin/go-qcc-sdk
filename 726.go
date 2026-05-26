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

type TelecomLicenseGetListReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 页码
	PageIndex string
	// 每页数据条数，默认为10，最大50
	PageSize string
}

type TelecomLicenseGetListResp struct {
	Response[[]TelecomLicenseGetListRespResult]
}

type TelecomLicenseGetListRespResult struct {
	ID          string `json:"Id"`
	KeyNo       string `json:"KeyNo"`
	CompanyName string `json:"CompanyName"`
	Scope       string `json:"Scope"`
	LicenseNo   string `json:"LicenseNo"`
	IsValid     string `json:"IsValid"`
}

// TelecomLicenseGetList 电信许可列表 https://openapi.qcc.com/dataApi/726
func (a *Api) TelecomLicenseGetList(ctx context.Context, req *TelecomLicenseGetListReq) (*TelecomLicenseGetListResp, error) {
	var resp TelecomLicenseGetListResp
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

	reply, err := c.SetResult(&resp).Get("/TelecomLicense/GetList")
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

type TelecomLicenseGetDetailReq struct {
	// 详情 Id
	ID string
}

type TelecomLicenseGetDetailResp struct {
	Response[TelecomLicenseGetDetailRespResult]
}

type TelecomLicenseGetDetailRespResult struct {
	ID           string                                        `json:"Id"`
	EntInfo      TelecomLicenseGetDetailRespResultEntInfo      `json:"EntInfo"`
	LicenseNo    string                                        `json:"LicenseNo"`
	IsValid      string                                        `json:"IsValid"`
	Scope        string                                        `json:"Scope"`
	AnnualReport TelecomLicenseGetDetailRespResultAnnualReport `json:"AnnualReport"`
}

type TelecomLicenseGetDetailRespResultEntInfo struct {
	KeyNo       string `json:"KeyNo"`
	CompanyName string `json:"CompanyName"`
}

type TelecomLicenseGetDetailRespResultAnnualReport struct {
	OperInfo       TelecomLicenseGetDetailRespResultAnnualReportOperInfo `json:"OperInfo"`
	CreditCode     string                                                `json:"CreditCode"`
	EntInfo        TelecomLicenseGetDetailRespResultEntInfo              `json:"EntInfo"`
	ReComplaintPer string                                                `json:"ReComplaintPer"`
	Address        string                                                `json:"Address"`
	Province       string                                                `json:"Province"`
	LicenseNo      string                                                `json:"LicenseNo"`
	Type           string                                                `json:"Type"`
	ComplaintNum   string                                                `json:"ComplaintNum"`
	RegistCapi     string                                                `json:"RegistCapi"`
	EconKind       string                                                `json:"EconKind"`
	IPOStatus      string                                                `json:"IPOStatus"`
	Tel            string                                                `json:"Tel"`
	StockCode      string                                                `json:"StockCode"`
}

type TelecomLicenseGetDetailRespResultAnnualReportOperInfo struct {
	OperID   string `json:"OperId"`
	OperName string `json:"OperName"`
}

// TelecomLicenseGetDetail 电信许可详情 https://openapi.qcc.com/dataApi/726
func (a *Api) TelecomLicenseGetDetail(ctx context.Context, req *TelecomLicenseGetDetailReq) (*TelecomLicenseGetDetailResp, error) {
	var resp TelecomLicenseGetDetailResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("id", req.ID)

	reply, err := c.SetResult(&resp).Get("/TelecomLicense/GetDetail")
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
