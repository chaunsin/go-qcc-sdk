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

type TaxDataCreateOrderReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 税务局用户名（SM4加密，密钥为用户key）
	UserName string
	// 税务局密码（SM4加密，密钥为用户key）
	Password string
}

type TaxDataCreateOrderResp struct {
	Response[TaxDataCreateOrderRespResult]
}

type TaxDataCreateOrderRespResult struct {
	DataStatus  string `json:"DataStatus"`
	OrderNo     string `json:"OrderNo"`
	OrderResult string `json:"OrderResult"`
}

// TaxDataCreateOrder 数据下单 https://openapi.qcc.com/dataApi/1124
func (a *Api) TaxDataCreateOrder(ctx context.Context, req *TaxDataCreateOrderReq) (*TaxDataCreateOrderResp, error) {
	var resp TaxDataCreateOrderResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	body := map[string]any{}
	body["searchKey"] = req.SearchKey
	body["userName"] = req.UserName
	body["password"] = req.Password
	c.SetBody(body)

	reply, err := c.SetResult(&resp).Post("/TaxData/CreateOrder")
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

type TaxDataSendCodeReq struct {
	// 订单号
	OrderNo string
	// 验证码
	VerifyCode string
}

type TaxDataSendCodeResp struct {
	Response[TaxDataSendCodeRespResult]
}

type TaxDataSendCodeRespResult struct {
	DataStatus  string `json:"DataStatus"`
	OrderNo     string `json:"OrderNo"`
	OrderResult string `json:"OrderResult"`
}

// TaxDataSendCode 验证码发送 https://openapi.qcc.com/dataApi/1124
func (a *Api) TaxDataSendCode(ctx context.Context, req *TaxDataSendCodeReq) (*TaxDataSendCodeResp, error) {
	var resp TaxDataSendCodeResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("orderNo", req.OrderNo).
		SetQueryParam("verifyCode", req.VerifyCode)

	reply, err := c.SetResult(&resp).Get("/TaxData/SendCode")
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

type TaxDataGetDataReq struct {
	// 订单号（有效期 7 天）
	OrderNo string
}

type TaxDataGetDataResp struct {
	Response[TaxDataGetDataRespResult]
}

type TaxDataGetDataRespResult struct {
	DataStatus string                       `json:"DataStatus"`
	Data       TaxDataGetDataRespResultData `json:"Data"`
}

type TaxDataGetDataRespResultData struct {
	FinancialIndexList   []TaxDataFinancialIndexItem   `json:"FinancialIndexList"`
	DeclarationDetail    TaxDataDeclarationDetail      `json:"DeclarationDetail"`
	CollectionDetail     TaxDataCollectionDetail       `json:"CollectionDetail"`
	SaleList             []TaxDataYearAmountList       `json:"SaleList"`
	TaxData              TaxDataTaxData                `json:"TaxData"`
	TaxBurdenRateList    []TaxDataTaxBurdenRateItem    `json:"TaxBurdenRateList"`
	FinancialList        []TaxDataFinancialItem        `json:"FinancialList"`
	SupplierCustomerList []TaxDataSupplierCustomerItem `json:"SupplierCustomerList"`
	TopCustomerList      []TaxDataTopCustomerItem      `json:"TopCustomerList"`
	TopSupplierList      []TaxDataTopSupplierItem      `json:"TopSupplierList"`
	BreakLawDetailList   []TaxDataBreakLawDetailItem   `json:"BreakLawDetailList"`
	BreakLawSummaryList  []TaxDataBreakLawSummaryItem  `json:"BreakLawSummaryList"`
	ExpenseDetail        TaxDataExpenseDetail          `json:"ExpenseDetail"`
	CashFlowList         []TaxDataCashFlowItem         `json:"CashFlowList"`
}

type TaxDataFinancialIndexItem struct {
	IndexName string                       `json:"IndexName"`
	ValueList []TaxDataFinancialIndexValue `json:"ValueList"`
}

type TaxDataFinancialIndexValue struct {
	Date  string `json:"Date"`
	Value string `json:"Value"`
}

type TaxDataDeclarationDetail struct {
	CorporateInTaxDeclareList []TaxDataCorporateInTaxDeclareItem `json:"CorporateInTaxDeclareList"`
	ValueAddedTaxDeclareList  []TaxDataValueAddedTaxDeclareItem  `json:"ValueAddedTaxDeclareList"`
	OtherTaxDeclareList       []TaxDataOtherTaxDeclareItem       `json:"OtherTaxDeclareList"`
}

type TaxDataCorporateInTaxDeclareItem struct {
	ThisYearSaleRevenue      string `json:"ThisYearSaleRevenue"`
	ThisYearCumulativeProfit string `json:"ThisYearCumulativeProfit"`
	StartDate                string `json:"StartDate"`
	EndDate                  string `json:"EndDate"`
	TaxPayable               string `json:"TaxPayable"`
	WithholdingTax           string `json:"WithholdingTax"`
	TaxCompensate            string `json:"TaxCompensate"`
	TaxDeduction             string `json:"TaxDeduction"`
}

type TaxDataValueAddedTaxDeclareItem struct {
	AllSaleRevenue       string `json:"AllSaleRevenue"`
	AllCumulativeRevenue string `json:"AllCumulativeRevenue"`
	StartDate            string `json:"StartDate"`
	EndDate              string `json:"EndDate"`
	TaxPayable           string `json:"TaxPayable"`
	WithholdingTax       string `json:"WithholdingTax"`
	TaxCompensate        string `json:"TaxCompensate"`
	TaxDeduction         string `json:"TaxDeduction"`
}

type TaxDataOtherTaxDeclareItem struct {
	LevyItemCode   string `json:"LevyItemCode"`
	LevyItemValue  string `json:"LevyItemValue"`
	TaxBasis       string `json:"TaxBasis"`
	StartDate      string `json:"StartDate"`
	EndDate        string `json:"EndDate"`
	TaxPayable     string `json:"TaxPayable"`
	WithholdingTax string `json:"WithholdingTax"`
	TaxCompensate  string `json:"TaxCompensate"`
	TaxDeduction   string `json:"TaxDeduction"`
}

type TaxDataCollectionDetail struct {
	CorporateInTaxCollectionList []TaxDataCorporateInTaxCollectionItem `json:"CorporateInTaxCollectionList"`
	ValueAddedTaxCollectionList  []TaxDataValueAddedTaxCollectionItem  `json:"ValueAddedTaxCollectionList"`
	OtherTaxCollectionList       []TaxDataOtherTaxCollectionItem       `json:"OtherTaxCollectionList"`
}

type TaxDataCorporateInTaxCollectionItem struct {
	ThisYearCumulativeProfit string `json:"ThisYearCumulativeProfit"`
	StartDate                string `json:"StartDate"`
	EndDate                  string `json:"EndDate"`
	PaymentLimitDate         string `json:"PaymentLimitDate"`
	PaymentDate              string `json:"PaymentDate"`
	TaxType                  string `json:"TaxType"`
	TaxRate                  string `json:"TaxRate"`
	ActualAmount             string `json:"ActualAmount"`
}

type TaxDataValueAddedTaxCollectionItem struct {
	SaleRevenue      string `json:"SaleRevenue"`
	StartDate        string `json:"StartDate"`
	EndDate          string `json:"EndDate"`
	PaymentLimitDate string `json:"PaymentLimitDate"`
	PaymentDate      string `json:"PaymentDate"`
	TaxType          string `json:"TaxType"`
	TaxRate          string `json:"TaxRate"`
	ActualAmount     string `json:"ActualAmount"`
}

type TaxDataOtherTaxCollectionItem struct {
	LevyItemCode     string `json:"LevyItemCode"`
	LevyItemValue    string `json:"LevyItemValue"`
	TaxBasis         string `json:"TaxBasis"`
	StartDate        string `json:"StartDate"`
	EndDate          string `json:"EndDate"`
	PaymentLimitDate string `json:"PaymentLimitDate"`
	PaymentDate      string `json:"PaymentDate"`
	TaxType          string `json:"TaxType"`
	TaxRate          string `json:"TaxRate"`
	ActualAmount     string `json:"ActualAmount"`
}

type TaxDataTaxData struct {
	TotalTaxList       []TaxDataYearAmountList     `json:"TotalTaxList"`
	CorporateInTaxList []TaxDataCorporateInTaxItem `json:"CorporateInTaxList"`
	ValueAddedTaxList  []TaxDataYearAmountList     `json:"ValueAddedTaxList"`
	OtherTaxList       []TaxDataOtherTaxItem       `json:"OtherTaxList"`
}

type TaxDataYearAmountList struct {
	Year     string                   `json:"Year"`
	DataList []TaxDataMonthAmountItem `json:"DataList"`
}

type TaxDataMonthAmountItem struct {
	Month  string `json:"Month"`
	Amount string `json:"Amount"`
}

type TaxDataCorporateInTaxItem struct {
	Year              string                     `json:"Year"`
	AnnualTax         string                     `json:"AnnualTax"`
	QuarterlyDataList []TaxDataQuarterAmountItem `json:"QuarterlyDataList"`
}

type TaxDataQuarterAmountItem struct {
	Quarter string `json:"Quarter"`
	Amount  string `json:"Amount"`
}

type TaxDataOtherTaxItem struct {
	LevyItemCode  string                  `json:"LevyItemCode"`
	LevyItemValue string                  `json:"LevyItemValue"`
	YearDataList  []TaxDataYearAmountList `json:"YearDataList"`
}

type TaxDataTaxBurdenRateItem struct {
	LevyItemCode  string                      `json:"LevyItemCode"`
	LevyItemValue string                      `json:"LevyItemValue"`
	DataList      []TaxDataTaxBurdenRateValue `json:"DataList"`
}

type TaxDataTaxBurdenRateValue struct {
	Year  string `json:"Year"`
	Ratio string `json:"Ratio"`
}

type TaxDataFinancialItem struct {
	Type        string                    `json:"Type"`
	TypeValue   string                    `json:"TypeValue"`
	SubjectList []TaxDataFinancialSubject `json:"SubjectList"`
}

type TaxDataFinancialSubject struct {
	Subject     string                  `json:"Subject"`
	RevenueList []TaxDataYearAmountItem `json:"RevenueList"`
}

type TaxDataYearAmountItem struct {
	Year   string `json:"Year"`
	Amount string `json:"Amount"`
}

type TaxDataBreakLawDetailItem struct {
	LimitChangeState string `json:"LimitChangeState"`
	RegistrationDate string `json:"RegistrationDate"`
	MainFact         string `json:"MainFact"`
	BreakType        string `json:"BreakType"`
	BreakStatus      string `json:"BreakStatus"`
}

type TaxDataBreakLawSummaryItem struct {
	GeneralCount    string `json:"GeneralCount"`
	Date            string `json:"Date"`
	SeriousCount    string `json:"SeriousCount"`
	InspectionCount string `json:"InspectionCount"`
}

type TaxDataSupplierCustomerItem struct {
	Type         string                            `json:"Type"`
	TypeValue    string                            `json:"TypeValue"`
	YearDataList []TaxDataSupplierCustomerYearData `json:"YearDataList"`
}

type TaxDataSupplierCustomerYearData struct {
	Year                     string                        `json:"Year"`
	SupplierCustomerInfoList []TaxDataSupplierCustomerInfo `json:"SupplierCustomerInfoList"`
}

type TaxDataSupplierCustomerInfo struct {
	Name       string `json:"Name"`
	Amount     string `json:"Amount"`
	Proportion string `json:"Proportion"`
}

type TaxDataTopCustomerItem struct {
	Year               string `json:"Year"`
	RepeatCount        string `json:"RepeatCount"`
	RepeatAmount       string `json:"RepeatAmont"`
	TotalAmount        string `json:"TotalAmount"`
	PurchaseProportion string `json:"PurchaseProportion"`
}

type TaxDataTopSupplierItem struct {
	Year               string `json:"Year"`
	RepeatCount        string `json:"RepeatCount"`
	TotalAmount        string `json:"TotalAmount"`
	SupplierProportion string `json:"SupplierProportion"`
}

type TaxDataExpenseDetail struct {
	ElectricityExpenseList         []TaxDataExpenseYearAmountList `json:"ElectricityExpenseList"`
	WaterExpenseList               []TaxDataExpenseYearAmountList `json:"WaterExpenseList"`
	GasExpenseList                 []TaxDataExpenseYearAmountList `json:"GasExpenseList"`
	HouseRentalExpenseList         []TaxDataExpenseYearAmountList `json:"HouseRentalExpenseList"`
	TransportAndStorageExpenseList []TaxDataExpenseYearAmountList `json:"TransportAndStorageExpenseList"`
}

type TaxDataExpenseYearAmountList struct {
	Year     string                   `json:"Year"`
	DataList []TaxDataMonthAmountItem `json:"DataList"`
}

type TaxDataCashFlowItem struct {
	Subject     string                  `json:"Subject"`
	RevenueList []TaxDataYearAmountItem `json:"RevenueList"`
}

// TaxDataGetData 数据获取 https://openapi.qcc.com/dataApi/1124
func (a *Api) TaxDataGetData(ctx context.Context, req *TaxDataGetDataReq) (*TaxDataGetDataResp, error) {
	var resp TaxDataGetDataResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("orderNo", req.OrderNo)

	reply, err := c.SetResult(&resp).Get("/TaxData/GetData")
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
