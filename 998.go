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

type HKNRTDataCreateOrderReq struct {
	// 搜索关键词（企业中文名称、企业英文名称、企业商业登记号码）
	HkEntityName string
}

type HKNRTDataCreateOrderResp struct {
	Response[HKNRTDataCreateOrderRespResult]
}

type HKNRTDataCreateOrderRespResult struct {
	OrderNo string `json:"OrderNo"`
}

// HKNRTDataCreateOrder 数据下单 https://openapi.qcc.com/dataApi/998
func (a *Api) HKNRTDataCreateOrder(ctx context.Context, req *HKNRTDataCreateOrderReq) (*HKNRTDataCreateOrderResp, error) {
	var resp HKNRTDataCreateOrderResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("hkEntityName", req.HkEntityName)

	reply, err := c.SetResult(&resp).Get("/HKNRTData/CreateOrder")
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

type HKNRTDataGetDataReq struct {
	// 订单号
	OrderNo string
}

type HKNRTDataGetDataResp struct {
	Response[HKNRTDataGetDataRespResult]
}

type HKNRTDataGetDataRespResult struct {
	DataStatus string                   `json:"DataStatus"`
	Data       HKNRTDataGetDataRespData `json:"Data"`
}

type HKNRTDataGetDataRespData struct {
	Basic                        HKNRTDataBasic                        `json:"basic"`
	CapitalStructureHistorical   HKNRTDataCapitalStructureHistorical   `json:"capitalstructure_historical"`
	DirectorHistorical           HKNRTDataDirectorHistorical           `json:"director_historical"`
	Shareholders                 HKNRTDataShareholders                 `json:"shareholders"`
	CompanySecretariesHistorical HKNRTDataCompanySecretariesHistorical `json:"company_secretaries_historical"`
	Verified                     string                                `json:"verified"`
	OriginalFile                 string                                `json:"OriginalFile"`
}

type HKNRTDataBasic struct {
	Details HKNRTDataBasicDetails `json:"Details"`
}

type HKNRTDataBasicDetails struct {
	ID                         string                  `json:"Id"`
	CompanyNumber              string                  `json:"CompanyNumber"`
	BusinessRegistrationNumber string                  `json:"BusinessRegistrationNumber"`
	CompanyNameChn             string                  `json:"CompanyNameChn"`
	CompanyNameEng             string                  `json:"CompanyNameEng"`
	CompanyType                string                  `json:"CompanyType"`
	RegistrationDate           string                  `json:"RegistrationDate"`
	Status                     string                  `json:"Status"`
	Address                    string                  `json:"Address"`
	AddressTrans               string                  `json:"Address_Trans"`
	OfficeEffectiveDate        string                  `json:"OfficeEffectiveDate"`
	WindingUpMode              string                  `json:"WindingUpMode"`
	RegisterOfCharges          string                  `json:"RegisterOfCharges"`
	DissolutionDate            string                  `json:"DissolutionDate"`
	Important                  string                  `json:"Important"`
	IncorporationPlace         string                  `json:"IncorporationPlace"`
	Remarks                    string                  `json:"Remarks"`
	OriginalNameList           []HKNRTDataOriginalName `json:"OriginalNameList"`
}

type HKNRTDataOriginalName struct {
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	Name      string `json:"Name"`
	EnName    string `json:"EnName"`
}

type HKNRTDataCapitalStructureHistorical struct {
	Date    string                                     `json:"Date"`
	Details []HKNRTDataCapitalStructureHistoricalEntry `json:"Details"`
}

type HKNRTDataCapitalStructureHistoricalEntry struct {
	ClassOfShares   string `json:"ClassofShares"`
	Currency        string `json:"Currency"`
	TotalAmount     string `json:"TotalAmount"`
	TotalAmountPaid string `json:"TotalAmountPaid"`
	TotalNumber     string `json:"TotalNumber"`
}

type HKNRTDataDirectorHistorical struct {
	Date    string                             `json:"Date"`
	Details []HKNRTDataDirectorHistoricalEntry `json:"Details"`
}

type HKNRTDataDirectorHistoricalEntry struct {
	FullNameEng     string `json:"FullNameEng"`
	FullNameChn     string `json:"FullNameChn"`
	Position        string `json:"Position"`
	Type            string `json:"Type"`
	ID              string `json:"Id"`
	PassportCountry string `json:"PassportCountry"`
	PassportNumber  string `json:"PassportNumber"`
}

type HKNRTDataShareholders struct {
	Date    string                 `json:"Date"`
	Details []HKNRTDataShareholder `json:"Details"`
}

type HKNRTDataShareholder struct {
	Address        string `json:"Address"`
	ClassOfShares  string `json:"ClassofShares"`
	FullName       string `json:"FullName"`
	FullNameChn    string `json:"FullNameChn"`
	FullNameEng    string `json:"FullNameEng"`
	NumberOfShares string `json:"NumberofShares"`
	PercentOfClass string `json:"PercentofClass"`
	TotalNumber    string `json:"TotalNumber"`
}

type HKNRTDataCompanySecretariesHistorical struct {
	Date    string                                `json:"Date"`
	Details []HKNRTDataCompanySecretaryHistorical `json:"Details"`
}

type HKNRTDataCompanySecretaryHistorical struct {
	FullNameChn string `json:"FullNameChn"`
	FullNameEng string `json:"FullNameEng"`
	Type        string `json:"Type"`
	Address     string `json:"Address"`
}

// HKNRTDataGetData 数据获取 https://openapi.qcc.com/dataApi/998
func (a *Api) HKNRTDataGetData(ctx context.Context, req *HKNRTDataGetDataReq) (*HKNRTDataGetDataResp, error) {
	var resp HKNRTDataGetDataResp
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

	reply, err := c.SetResult(&resp).Get("/HKNRTData/GetData")
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
