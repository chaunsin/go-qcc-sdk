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

type ECIInfoOverviewGetInfoReq struct {
	SearchKey string
}

type ECIInfoOverviewGetInfoResp struct {
	PermissionInfo []struct {
		Name     string `json:"Name"`
		Province string `json:"Province"`
		Liandate string `json:"Liandate"`
		CaseNo   string `json:"CaseNo"`
	} `json:"PermissionInfo"`
	Penalty []struct {
		DocNo       string      `json:"DocNo"`
		PenaltyType string      `json:"PenaltyType"`
		OfficeName  string      `json:"OfficeName"`
		Content     string      `json:"Content"`
		PenaltyDate string      `json:"PenaltyDate"`
		PublicDate  string      `json:"PublicDate"`
		Remark      interface{} `json:"Remark"`
	} `json:"Penalty"`
	Exceptions   []interface{} `json:"Exceptions"`
	ShiXinItems  []interface{} `json:"ShiXinItems"`
	ZhiXingItems []interface{} `json:"ZhiXingItems"`
	MPledge      []interface{} `json:"MPledge"`
	Liquidation  interface{}   `json:"Liquidation"`
	Pledge       []struct {
		CompanyId     string `json:"CompanyId"`
		RegistNo      string `json:"RegistNo"`
		Pledgor       string `json:"Pledgor"`
		PledgorNo     string `json:"PledgorNo"`
		Pledgee       string `json:"Pledgee"`
		PledgeeNo     string `json:"PledgeeNo"`
		PledgedAmount string `json:"PledgedAmount"`
		RegDate       string `json:"RegDate"`
		PublicDate    string `json:"PublicDate"`
		Status        string `json:"Status"`
	} `json:"Pledge"`
	SpotCheck []struct {
		No           interface{} `json:"No"`
		ExecutiveOrg string      `json:"ExecutiveOrg"`
		Type         string      `json:"Type"`
		Date         string      `json:"Date"`
		Consequence  string      `json:"Consequence"`
		Remark       interface{} `json:"Remark"`
	} `json:"SpotCheck"`
	CompanyTaxCreditItems []struct {
		No    string `json:"No"`
		Name  string `json:"Name"`
		Year  string `json:"Year"`
		Level string `json:"Level"`
	} `json:"CompanyTaxCreditItems"`
	CompanyProducts []struct {
		CompanyName string `json:"CompanyName"`
		Link        string `json:"Link"`
		ImageUrl    string `json:"ImageUrl"`
		Name        string `json:"Name"`
		Domain      string `json:"Domain"`
		Tags        string `json:"Tags"`
		Description string `json:"Description"`
		Category    string `json:"Category"`
	} `json:"CompanyProducts"`
	PermissionEciInfo []struct {
		LicensDocNo   string `json:"LicensDocNo"`
		LicensDocName string `json:"LicensDocName"`
		ValidityFrom  string `json:"ValidityFrom"`
		ValidityTo    string `json:"ValidityTo"`
		LicensOffice  string `json:"LicensOffice"`
		LicensContent string `json:"LicensContent"`
	} `json:"PermissionEciInfo"`
	PenaltyCreditChina []interface{} `json:"PenaltyCreditChina"`
	Partners           []struct {
		KeyNo               string      `json:"KeyNo"`
		StockName           string      `json:"StockName"`
		StockType           string      `json:"StockType"`
		StockPercent        string      `json:"StockPercent"`
		ShouldCapi          string      `json:"ShouldCapi"`
		ShoudDate           string      `json:"ShoudDate"`
		InvestType          string      `json:"InvestType"`
		InvestName          interface{} `json:"InvestName"`
		RealCapi            interface{} `json:"RealCapi"`
		CapiDate            interface{} `json:"CapiDate"`
		TagsList            []string    `json:"TagsList"`
		FinalBenefitPercent string      `json:"FinalBenefitPercent"`
		RelatedProduct      *struct {
			Id             string `json:"Id"`
			Name           string `json:"Name"`
			Round          string `json:"Round"`
			FinancingCount string `json:"FinancingCount"`
		} `json:"RelatedProduct"`
		RelatedOrg *struct {
			Id   string `json:"Id"`
			Name string `json:"Name"`
		} `json:"RelatedOrg"`
		Area string `json:"Area"`
	} `json:"Partners"`
	Employees []struct {
		KeyNo string `json:"KeyNo"`
		Name  string `json:"Name"`
		Job   string `json:"Job"`
	} `json:"Employees"`
	Branches      []interface{} `json:"Branches"`
	ChangeRecords []struct {
		ProjectName   string `json:"ProjectName"`
		BeforeContent string `json:"BeforeContent"`
		AfterContent  string `json:"AfterContent"`
		ChangeDate    string `json:"ChangeDate"`
	} `json:"ChangeRecords"`
	ContactInfo struct {
		WebSite []struct {
			Name interface{} `json:"Name"`
			Url  string      `json:"Url"`
		} `json:"WebSite"`
		PhoneNumber string `json:"PhoneNumber"`
		Email       string `json:"Email"`
	} `json:"ContactInfo"`
	Industry struct {
		IndustryCode       string      `json:"IndustryCode"`
		Industry           string      `json:"Industry"`
		SubIndustryCode    string      `json:"SubIndustryCode"`
		SubIndustry        string      `json:"SubIndustry"`
		MiddleCategoryCode interface{} `json:"MiddleCategoryCode"`
		MiddleCategory     interface{} `json:"MiddleCategory"`
		SmallCategoryCode  interface{} `json:"SmallCategoryCode"`
		SmallCategory      interface{} `json:"SmallCategory"`
	} `json:"Industry"`
	Area struct {
		Province string `json:"Province"`
		City     string `json:"City"`
		County   string `json:"County"`
	} `json:"Area"`
	AreaCode            string `json:"AreaCode"`
	EmergingIndustyList []struct {
		PrimaryCode   string `json:"PrimaryCode"`
		PrimaryDes    string `json:"PrimaryDes"`
		SecondaryList []struct {
			SecondaryCode string `json:"SecondaryCode"`
			SecondaryDes  string `json:"SecondaryDes"`
			TertiaryList  []struct {
				TertiaryCode string `json:"TertiaryCode"`
				TertiaryDes  string `json:"TertiaryDes"`
			} `json:"TertiaryList"`
		} `json:"SecondaryList"`
	} `json:"EmergingIndustyList"`
	RevokeInfo   interface{} `json:"RevokeInfo"`
	InsuredCount string      `json:"InsuredCount"`
	EnglishName  string      `json:"EnglishName"`
	PersonScope  string      `json:"PersonScope"`
	IXCode       interface{} `json:"IXCode"`
	TagList      []struct {
		Type string `json:"Type"`
		Name string `json:"Name"`
	} `json:"TagList"`
	ARContactList []struct {
		ContactNo    string `json:"ContactNo"`
		EmailAddress string `json:"EmailAddress"`
		Address      string `json:"Address"`
	} `json:"ARContactList"`
	EconKindCodeList []string    `json:"EconKindCodeList"`
	KeyNo            string      `json:"KeyNo"`
	Name             string      `json:"Name"`
	No               string      `json:"No"`
	BelongOrg        string      `json:"BelongOrg"`
	OperId           string      `json:"OperId"`
	OperName         string      `json:"OperName"`
	StartDate        string      `json:"StartDate"`
	EndDate          string      `json:"EndDate"`
	Status           string      `json:"Status"`
	Province         string      `json:"Province"`
	UpdatedDate      string      `json:"UpdatedDate"`
	CreditCode       string      `json:"CreditCode"`
	RegistCapi       string      `json:"RegistCapi"`
	EconKind         string      `json:"EconKind"`
	Address          string      `json:"Address"`
	Scope            string      `json:"Scope"`
	TermStart        string      `json:"TermStart"`
	TermEnd          string      `json:"TermEnd"`
	CheckDate        string      `json:"CheckDate"`
	OrgNo            string      `json:"OrgNo"`
	IsOnStock        string      `json:"IsOnStock"`
	StockNumber      interface{} `json:"StockNumber"`
	StockType        interface{} `json:"StockType"`
	OriginalName     []struct {
		Name       string `json:"Name"`
		ChangeDate string `json:"ChangeDate"`
	} `json:"OriginalName"`
	ImageUrl string `json:"ImageUrl"`
	EntType  string `json:"EntType"`
	RecCap   string `json:"RecCap"`
}

// ECIInfoOverviewGetInfo 企业风险扫描 https://openapi.qcc.com/dataApi/736
func (a *Api) ECIInfoOverviewGetInfo(ctx context.Context, req *ECIInfoOverviewGetInfoReq) (*ECIInfoOverviewGetInfoResp, error) {
	var resp ECIInfoOverviewGetInfoResp
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
		Get("https://api.qichacha.com/ECIInfoOverview/GetInfo")
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
