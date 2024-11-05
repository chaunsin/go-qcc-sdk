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

type EnterpriseInfoVerifyReq struct {
	SearchKey string
}

type EnterpriseInfoVerifyResp struct {
	Response[EnterpriseInfoVerifyRespResult]
}

type EnterpriseInfoVerifyRespResult struct {
	VerifyResult int64                                `json:"VerifyResult"`
	Data         []EnterpriseInfoVerifyRespResultData `json:"Data"`
}

type EnterpriseInfoVerifyRespResultData struct {
	KeyNo        string `json:"KeyNo"`
	Name         string `json:"Name"`
	CreditCode   string `json:"CreditCode"`
	OperName     string `json:"OperName"`
	Status       string `json:"Status"`
	StartDate    string `json:"StartDate"`
	RegistCapi   string `json:"RegistCapi"`
	RealCapi     string `json:"RealCapi"`
	OrgNo        string `json:"OrgNo"`
	No           string `json:"No"`
	TaxNo        string `json:"TaxNo"`
	EconKind     string `json:"EconKind"`
	TermStart    string `json:"TermStart"`
	TermEnd      string `json:"TermEnd"`
	TaxpayerType string `json:"TaxpayerType"`
	PersonScope  string `json:"PersonScope"`
	InsuredCount string `json:"InsuredCount"`
	CheckDate    string `json:"CheckDate"`
	AreaCode     string `json:"AreaCode"`
	Area         struct {
		Province string `json:"Province"`
		City     string `json:"City"`
		County   string `json:"County"`
	} `json:"Area"`
	BelongOrg string `json:"BelongOrg"`
	ImExCode  string `json:"ImExCode"`
	Industry  struct {
		IndustryCode       string `json:"IndustryCode"`
		Industry           string `json:"Industry"`
		SubIndustryCode    string `json:"SubIndustryCode"`
		SubIndustry        string `json:"SubIndustry"`
		MiddleCategoryCode string `json:"MiddleCategoryCode"`
		MiddleCategory     string `json:"MiddleCategory"`
		SmallCategoryCode  string `json:"SmallCategoryCode"`
		SmallCategory      string `json:"SmallCategory"`
	} `json:"Industry"`
	EnglishName   string `json:"EnglishName"`
	Address       string `json:"Address"`
	AnnualAddress string `json:"AnnualAddress"`
	Scope         string `json:"Scope"`
	EntType       string `json:"EntType"`
	OrgCodeList   []struct {
		PrimaryCode   string `json:"PrimaryCode"`
		SecondaryCode string `json:"SecondaryCode"`
	} `json:"OrgCodeList"`
	ImageUrl   string `json:"ImageUrl"`
	RevokeInfo struct {
		CancelDate   string `json:"CancelDate"`
		CancelReason string `json:"CancelReason"`
		RevokeDate   string `json:"RevokeDate"`
		RevokeReason string `json:"RevokeReason"`
	} `json:"RevokeInfo"`
	OriginalName []struct {
		Name       string `json:"Name"`
		ChangeDate string `json:"ChangeDate"`
	} `json:"OriginalName"`
	StockInfo struct {
		StockNumber string `json:"StockNumber"`
		StockType   string `json:"StockType"`
	} `json:"StockInfo"`
	ContactInfo struct {
		WebSiteList   []string `json:"WebSiteList"`
		Email         string   `json:"Email"`
		MoreEmailList []struct {
			Email  string `json:"Email"`
			Source string `json:"Source"`
		} `json:"MoreEmailList"`
		Tel         string `json:"Tel"`
		MoreTelList []struct {
			Tel    string `json:"Tel"`
			Source string `json:"Source"`
		} `json:"MoreTelList"`
	} `json:"ContactInfo"`
	LongLat struct {
		Longitude string `json:"Longitude"`
		Latitude  string `json:"Latitude"`
	} `json:"LongLat"`
	BankInfo struct {
		Bank        string `json:"Bank"`
		BankAccount string `json:"BankAccount"`
		Name        string `json:"Name"`
		CreditCode  string `json:"CreditCode"`
		Address     string `json:"Address"`
		Tel         string `json:"Tel"`
	} `json:"BankInfo"`
	IsSmall     string `json:"IsSmall"`
	Scale       string `json:"Scale"`
	QccIndustry struct {
		AName string `json:"AName"`
		BName string `json:"BName"`
		CName string `json:"CName"`
		DName string `json:"DName"`
	} `json:"QccIndustry"`
	IsOfficialEnglish string `json:"IsOfficialEnglish"`
}

// EnterpriseInfoVerify 企业信息核验 https://openapi.qcc.com/dataApi/2001
func (a *Api) EnterpriseInfoVerify(ctx context.Context, req *EnterpriseInfoVerifyReq) (*EnterpriseInfoVerifyResp, error) {
	var resp EnterpriseInfoVerifyResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("searchKey", req.SearchKey).
		SetResult(&resp).
		Get("https://api.qichacha.com/EnterpriseInfo/Verify")
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
