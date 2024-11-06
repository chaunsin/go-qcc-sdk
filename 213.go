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
	"time"
)

type ARGetAnnualReportReq struct {
	KeyNo string
}

type ARGetAnnualReportResp struct {
	Response[[]ARGetAnnualReportRespResult]
}

type ARGetAnnualReportRespResult struct {
	BasicInfoData struct {
		RegNo                 string      `json:"RegNo"`
		CompanyName           string      `json:"CompanyName"`
		CreditCode            string      `json:"CreditCode"`
		OperatorName          interface{} `json:"OperatorName"`
		ContactNo             string      `json:"ContactNo"`
		PostCode              string      `json:"PostCode"`
		Address               string      `json:"Address"`
		EmailAddress          string      `json:"EmailAddress"`
		IsStockRightTransfer  string      `json:"IsStockRightTransfer"`
		Status                string      `json:"Status"`
		HasWebSite            string      `json:"HasWebSite"`
		HasNewStockOrByStock  string      `json:"HasNewStockOrByStock"`
		EmployeeCount         string      `json:"EmployeeCount"`
		BelongTo              string      `json:"BelongTo"`
		CapitalAmount         interface{} `json:"CapitalAmount"`
		HasProvideAssurance   string      `json:"HasProvideAssurance"`
		OperationPlaces       interface{} `json:"OperationPlaces"`
		MainType              interface{} `json:"MainType"`
		OperationDuration     interface{} `json:"OperationDuration"`
		IfContentSame         interface{} `json:"IfContentSame"`
		DifferentContent      interface{} `json:"DifferentContent"`
		GeneralOperationItem  interface{} `json:"GeneralOperationItem"`
		ApprovedOperationItem interface{} `json:"ApprovedOperationItem"`
	} `json:"BasicInfoData"`
	AssetsData struct {
		TotalAssets        string      `json:"TotalAssets"`
		TotalOwnersEquity  string      `json:"TotalOwnersEquity"`
		GrossTradingIncome string      `json:"GrossTradingIncome"`
		TotalProfit        string      `json:"TotalProfit"`
		MainBusinessIncome string      `json:"MainBusinessIncome"`
		NetProfit          string      `json:"NetProfit"`
		TotalTaxAmount     string      `json:"TotalTaxAmount"`
		TotalLiabilities   string      `json:"TotalLiabilities"`
		BankingCredit      interface{} `json:"BankingCredit"`
		GovernmentSubsidy  interface{} `json:"GovernmentSubsidy"`
	} `json:"AssetsData"`
	ChangeList []struct {
		No         int    `json:"No"`
		ChangeName string `json:"ChangeName"`
		Before     string `json:"Before"`
		After      string `json:"After"`
		ChangeDate string `json:"ChangeDate"`
	} `json:"ChangeList"`
	InvestInfoList []interface{} `json:"InvestInfoList"`
	PartnerList    []struct {
		No              int    `json:"No"`
		Name            string `json:"Name"`
		ShouldCapi      string `json:"ShouldCapi"`
		ShouldDate      string `json:"ShouldDate"`
		ShouldType      string `json:"ShouldType"`
		RealCapi        string `json:"RealCapi"`
		RealDate        string `json:"RealDate"`
		RealType        string `json:"RealType"`
		Form            string `json:"Form"`
		InvestmentRatio string `json:"InvestmentRatio"`
		KeyNo           string `json:"KeyNo"`
	} `json:"PartnerList"`
	ProvideAssuranceList      []interface{} `json:"ProvideAssuranceList"`
	StockChangeList           []interface{} `json:"StockChangeList"`
	WebSiteList               []interface{} `json:"WebSiteList"`
	AdministrationLicenseList []interface{} `json:"AdministrationLicenseList"`
	BranchList                []interface{} `json:"BranchList"`
	EmployeeList              []interface{} `json:"EmployeeList"`
	SocialInsurance           struct {
		UrbanBasicIns                string `json:"UrbanBasicIns"`
		EmployeeBasicIns             string `json:"EmployeeBasicIns"`
		MaternityIns                 string `json:"MaternityIns"`
		UnemploymentIns              string `json:"UnemploymentIns"`
		IndustrialInjuryIns          string `json:"IndustrialInjuryIns"`
		EntUrbanBasicInsNum          string `json:"EntUrbanBasicInsNum"`
		EntUnemploymentInsNum        string `json:"EntUnemploymentInsNum"`
		EntEmployeeBasicInsNum       string `json:"EntEmployeeBasicInsNum"`
		EntIndustrialInjuryInsNum    string `json:"EntIndustrialInjuryInsNum"`
		EntMaternityInsNum           string `json:"EntMaternityInsNum"`
		ActUrbanBasicInsAmount       string `json:"ActUrbanBasicInsAmount"`
		ActUnemploymentInsAmount     string `json:"ActUnemploymentInsAmount"`
		ActEmployeeBasicInsAmount    string `json:"ActEmployeeBasicInsAmount"`
		ActIndustrialInjuryInsAmount string `json:"ActIndustrialInjuryInsAmount"`
		ActMaternityInsAmount        string `json:"ActMaternityInsAmount"`
		OweUrbanBasicInsAmount       string `json:"OweUrbanBasicInsAmount"`
		OweUnemploymentInsAmount     string `json:"OweUnemploymentInsAmount"`
		OweEmployeeBasicInsAmount    string `json:"OweEmployeeBasicInsAmount"`
		OweIndustrialInjuryInsAmount string `json:"OweIndustrialInjuryInsAmount"`
		OweMaternityInsAmount        string `json:"OweMaternityInsAmount"`
	} `json:"SocialInsurance"`
	No            int         `json:"No"`
	Year          string      `json:"Year"`
	Remarks       interface{} `json:"Remarks"`
	HasDetailInfo string      `json:"HasDetailInfo"`
	PublishDate   string      `json:"PublishDate"`
}

// ARGetAnnualReport 企业年报信息-查询公司年报信息 https://openapi.qcc.com/dataApi/213
func (a *Api) ARGetAnnualReport(ctx context.Context, req *ARGetAnnualReportReq) (*ARGetAnnualReportResp, error) {
	var resp ARGetAnnualReportResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("keyNo", req.KeyNo).
		SetResult(&resp).
		Get("https://api.qichacha.com/AR/GetAnnualReport")
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

type ARGetAnnualReportSummaryReq struct {
	KeyNo string
}

type ARGetAnnualReportSummaryResp struct {
	Response[[]ARGetAnnualReportSummaryResult]
}

type ARGetAnnualReportSummaryResult struct {
	No            int         `json:"No"`
	Year          string      `json:"Year"`
	Remarks       interface{} `json:"Remarks"`
	HasDetailInfo bool        `json:"HasDetailInfo"`
	PublishDate   time.Time   `json:"PublishDate"`
}

// ARGetAnnualReportSummary 企业年报信息-查询公司年报概况 https://openapi.qcc.com/dataApi/213
func (a *Api) ARGetAnnualReportSummary(ctx context.Context, req *ARGetAnnualReportSummaryReq) (*ARGetAnnualReportSummaryResp, error) {
	var resp ARGetAnnualReportSummaryResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("keyNo", req.KeyNo).
		SetResult(&resp).
		Get("https://api.qichacha.com/AR/GetAnnualReportSummary")
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
