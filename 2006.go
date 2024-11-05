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

type RiskControlScanReq struct {
	SearchKey string
}

type RiskControlScanResp struct {
	Response[RiskControlScanRespResult]
}

type RiskControlScanRespResult struct {
	VerifyResult int64                           `json:"VerifyResult"`
	Data         []RiskControlScanRespResultData `json:"Data"`
}

type RiskControlScanRespResultData struct {
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
	PartnerList       []struct {
		KeyNo        string `json:"KeyNo"`
		StockName    string `json:"StockName"`
		StockType    string `json:"StockType"`
		StockPercent string `json:"StockPercent"`
		ShouldCapi   string `json:"ShouldCapi"`
		ShoudDate    string `json:"ShoudDate"`
		StakeDate    string `json:"StakeDate"`
		CreditCode   string `json:"CreditCode"`
		Area         string `json:"Area"`
	} `json:"PartnerList"`
	PubPartnerList []struct {
		StockName    string `json:"StockName"`
		StockPercent string `json:"StockPercent"`
		HoldType     string `json:"HoldType"`
		Amount       string `json:"Amount"`
		CreditCode   string `json:"CreditCode"`
		Area         string `json:"Area"`
	} `json:"PubPartnerList"`
	EmployeeList []struct {
		KeyNo string `json:"KeyNo"`
		Name  string `json:"Name"`
		Job   string `json:"Job"`
	} `json:"EmployeeList"`
	PubEmployeeList []struct {
		Name string `json:"Name"`
		Job  string `json:"Job"`
	} `json:"PubEmployeeList"`
	BranchList []struct {
		KeyNo     string `json:"KeyNo"`
		Name      string `json:"Name"`
		OperName  string `json:"OperName"`
		StartDate string `json:"StartDate"`
		Status    string `json:"Status"`
		Area      struct {
			Province string `json:"Province"`
			City     string `json:"City"`
			County   string `json:"County"`
		} `json:"Area"`
	} `json:"BranchList"`
	ChangeList []struct {
		ProjectName string   `json:"ProjectName"`
		ChangeDate  string   `json:"ChangeDate"`
		BeforeList  []string `json:"BeforeList"`
		AfterList   []string `json:"AfterList"`
	} `json:"ChangeList"`
	TagList []struct {
		Type string `json:"Type"`
		Name string `json:"Name"`
	} `json:"TagList"`
	Parent struct {
		KeyNo      string `json:"KeyNo"`
		Name       string `json:"Name"`
		OperName   string `json:"OperName"`
		StartDate  string `json:"StartDate"`
		Status     string `json:"Status"`
		RegistCapi string `json:"RegistCapi"`
	} `json:"Parent"`
	Beneficiary struct {
		KeyNo               string `json:"KeyNo"`
		Name                string `json:"Name"`
		FinalBenefitPercent string `json:"FinalBenefitPercent"`
		Reason              string `json:"Reason"`
	} `json:"Beneficiary"`
	ActualControllerList []struct {
		KeyNo               string `json:"KeyNo"`
		Name                string `json:"Name"`
		FinalBenefitPercent string `json:"FinalBenefitPercent"`
		ControlPercent      string `json:"ControlPercent"`
		IsActual            string `json:"IsActual"`
	} `json:"ActualControllerList"`
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
	GroupInfo struct {
		GroupId string `json:"GroupId"`
		Name    string `json:"Name"`
		Logo    string `json:"Logo"`
	} `json:"GroupInfo"`
	InvestmentList []struct {
		KeyNo       string `json:"KeyNo"`
		Name        string `json:"Name"`
		StartDate   string `json:"StartDate"`
		Status      string `json:"Status"`
		FundedRatio string `json:"FundedRatio"`
		ShouldCapi  string `json:"ShouldCapi"`
		Industry    struct {
			IndustryCode       string `json:"IndustryCode"`
			Industry           string `json:"Industry"`
			SubIndustryCode    string `json:"SubIndustryCode"`
			SubIndustry        string `json:"SubIndustry"`
			MiddleCategoryCode string `json:"MiddleCategoryCode"`
			MiddleCategory     string `json:"MiddleCategory"`
			SmallCategoryCode  string `json:"SmallCategoryCode"`
			SmallCategory      string `json:"SmallCategory"`
		} `json:"Industry"`
		Area struct {
			Province string `json:"Province"`
			City     string `json:"City"`
			County   string `json:"County"`
		} `json:"Area"`
	} `json:"InvestmentList"`
	ProductList []struct {
		Name        string `json:"Name"`
		StartDate   string `json:"StartDate"`
		RoundDesc   string `json:"RoundDesc"`
		Location    string `json:"Location"`
		Description string `json:"Description"`
	} `json:"ProductList"`
	AdminLicenseList []struct {
		LicensDocNo   string `json:"LicensDocNo"`
		LicensDocName string `json:"LicensDocName"`
		ValidityFrom  string `json:"ValidityFrom"`
		ValidityTo    string `json:"ValidityTo"`
		LicensOffice  string `json:"LicensOffice"`
		LicensContent string `json:"LicensContent"`
		Source        string `json:"Source"`
	} `json:"AdminLicenseList"`
	ApproveSiteList []struct {
		Name       string `json:"Name"`
		WebAddress string `json:"WebAddress"`
		DomainName string `json:"DomainName"`
		LesenceNo  string `json:"LesenceNo"`
		AuditDate  string `json:"AuditDate"`
	} `json:"ApproveSiteList"`
	SpotCheckList []struct {
		ExecutiveOrg string `json:"ExecutiveOrg"`
		Type         string `json:"Type"`
		Date         string `json:"Date"`
		Consequence  string `json:"Consequence"`
	} `json:"SpotCheckList"`
	TaxCreditList []struct {
		TaxNo string `json:"TaxNo"`
		Year  string `json:"Year"`
		Level string `json:"Level"`
		Org   string `json:"Org"`
	} `json:"TaxCreditList"`
	ShiXin struct {
		TotalAmount string `json:"TotalAmount"`
		TotalCount  string `json:"TotalCount"`
		DataList    []struct {
			Id            string `json:"Id"`
			CaseNo        string `json:"CaseNo"`
			ExecuteCourt  string `json:"ExecuteCourt"`
			ExecuteNo     string `json:"ExecuteNo"`
			Amount        string `json:"Amount"`
			ExecuteStatus string `json:"ExecuteStatus"`
			ActionRemark  string `json:"ActionRemark"`
			RegisterDate  string `json:"RegisterDate"`
			PublicDate    string `json:"PublicDate"`
		} `json:"DataList"`
	} `json:"ShiXin"`
	ZhiXing struct {
		TotalAmount string `json:"TotalAmount"`
		TotalCount  string `json:"TotalCount"`
		DataList    []struct {
			Id           string `json:"Id"`
			CaseNo       string `json:"CaseNo"`
			BiaoDi       string `json:"BiaoDi"`
			ExecuteCourt string `json:"ExecuteCourt"`
			RegisterDate string `json:"RegisterDate"`
		} `json:"DataList"`
	} `json:"ZhiXing"`
	AdminPenalty struct {
		TotalAmount string `json:"TotalAmount"`
		TotalCount  string `json:"TotalCount"`
		DataList    []struct {
			Id           string `json:"Id"`
			DocNo        string `json:"DocNo"`
			PunishReason string `json:"PunishReason"`
			PunishResult string `json:"PunishResult"`
			PunishAmt    string `json:"PunishAmt"`
			PunishOffice string `json:"PunishOffice"`
			PunishDate   string `json:"PunishDate"`
		} `json:"DataList"`
	} `json:"AdminPenalty"`
	Exception struct {
		TotalCount string `json:"TotalCount"`
		DataList   []struct {
			AddDate   string `json:"AddDate"`
			AddOffice string `json:"AddOffice"`
			AddReason string `json:"AddReason"`
		} `json:"DataList"`
	} `json:"Exception"`
	ChattelMortgage struct {
		TotalCount string `json:"TotalCount"`
		DataList   []struct {
			RegisterNo         string   `json:"RegisterNo"`
			Status             string   `json:"Status"`
			RegisterDate       string   `json:"RegisterDate"`
			SecureClaimsAmount string   `json:"SecureClaimsAmount"`
			Pledger            []string `json:"Pledger"`
			Pledgee            []string `json:"Pledgee"`
			DebtTerm           string   `json:"DebtTerm"`
		} `json:"DataList"`
	} `json:"ChattelMortgage"`
	Liquidation struct {
		Leader string `json:"Leader"`
		Member string `json:"Member"`
	} `json:"Liquidation"`
	EquityPledge struct {
		TotalCount string `json:"TotalCount"`
		DataList   []struct {
			Id             string   `json:"Id"`
			RegisterNo     string   `json:"RegisterNo"`
			PledgorList    []string `json:"PledgorList"`
			PledgeeList    []string `json:"PledgeeList"`
			RelatedCompany string   `json:"RelatedCompany"`
			PledgedAmount  string   `json:"PledgedAmount"`
			RegisterDate   string   `json:"RegisterDate"`
			Status         string   `json:"Status"`
		} `json:"DataList"`
	} `json:"EquityPledge"`
	SeriousIllegal struct {
		TotalCount string `json:"TotalCount"`
		DataList   []struct {
			AddDate   string `json:"AddDate"`
			AddOffice string `json:"AddOffice"`
			AddReason string `json:"AddReason"`
		} `json:"DataList"`
	} `json:"SeriousIllegal"`
	EquityFreeze struct {
		TotalCount string `json:"TotalCount"`
		DataList   []struct {
			Id              string `json:"Id"`
			DocNo           string `json:"DocNo"`
			BeExecuted      string `json:"BeExecuted"`
			FreezeCompany   string `json:"FreezeCompany"`
			EquityAmount    string `json:"EquityAmount"`
			ExecuteCourt    string `json:"ExecuteCourt"`
			Status          string `json:"Status"`
			FreezeStartDate string `json:"FreezeStartDate"`
			FreezeEndDate   string `json:"FreezeEndDate"`
		} `json:"DataList"`
	} `json:"EquityFreeze"`
	JudicialSale struct {
		TotalCount string `json:"TotalCount"`
		DataList   []struct {
			Id          string `json:"Id"`
			Name        string `json:"Name"`
			CaseNo      string `json:"CaseNo"`
			AuctionTime string `json:"AuctionTime"`
			ExecuteGov  string `json:"ExecuteGov"`
		} `json:"DataList"`
	} `json:"JudicialSale"`
	Bankruptcy struct {
		TotalCount string `json:"TotalCount"`
		DataList   []struct {
			Id             string   `json:"Id"`
			CaseNo         string   `json:"CaseNo"`
			PublicDate     string   `json:"PublicDate"`
			ApplicantList  []string `json:"ApplicantList"`
			RespondentList []string `json:"RespondentList"`
		} `json:"DataList"`
	} `json:"Bankruptcy"`
	Sumptuary struct {
		TotalAmount string `json:"TotalAmount"`
		TotalCount  string `json:"TotalCount"`
		DataList    []struct {
			Id           string `json:"Id"`
			CaseNo       string `json:"CaseNo"`
			CompanyName  string `json:"CompanyName"`
			RelatedName  string `json:"RelatedName"`
			Applicant    string `json:"Applicant"`
			Amount       string `json:"Amount"`
			RegisterDate string `json:"RegisterDate"`
			PublicDate   string `json:"PublicDate"`
			ExecuteCourt string `json:"ExecuteCourt"`
		} `json:"DataList"`
	} `json:"Sumptuary"`
	EnvPunishment struct {
		TotalAmount string `json:"TotalAmount"`
		TotalCount  string `json:"TotalCount"`
		DataList    []struct {
			Id           string `json:"Id"`
			DocNo        string `json:"DocNo"`
			PunishReason string `json:"PunishReason"`
			PunishResult string `json:"PunishResult,omitempty"`
			PunishAmt    string `json:"PunishAmt"`
			PunishOffice string `json:"PunishOffice"`
			PunishDate   string `json:"PunishDate"`
		} `json:"DataList"`
	} `json:"EnvPunishment"`
	TaxOweNotice struct {
		TotalAmount string `json:"TotalAmount"`
		TotalCount  string `json:"TotalCount"`
		DataList    []struct {
			Id            string `json:"Id"`
			Title         string `json:"Title"`
			Amount        string `json:"Amount"`
			NewAmount     string `json:"NewAmount"`
			PublishDate   string `json:"PublishDate"`
			PublishOffice string `json:"PublishOffice"`
		} `json:"DataList"`
	} `json:"TaxOweNotice"`
	TaxIllegal struct {
		TotalCount string `json:"TotalCount"`
		DataList   []struct {
			Id             string `json:"Id"`
			PublishDate    string `json:"PublishDate"`
			CaseNature     string `json:"CaseNature"`
			TaxGov         string `json:"TaxGov"`
			IllegalContent string `json:"IllegalContent"`
			PunishContent  string `json:"PunishContent"`
		} `json:"DataList"`
	} `json:"TaxIllegal"`
	TaxAbnormal struct {
		TotalCount string `json:"TotalCount"`
		DataList   []struct {
			TaxNo     string `json:"TaxNo"`
			AddOffice string `json:"AddOffice"`
			AddDate   string `json:"AddDate"`
		} `json:"DataList"`
	} `json:"TaxAbnormal"`
}

// RiskControlScan 综合风险排查 https://openapi.qcc.com/dataApi/2006
func (a *Api) RiskControlScan(ctx context.Context, req *RiskControlScanReq) (*RiskControlScanResp, error) {
	var resp RiskControlScanResp
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
		Get("https://api.qichacha.com/RiskControl/Scan")
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
