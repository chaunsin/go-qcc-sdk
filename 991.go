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

type HKDataCreateOrderReq struct {
	// 搜索关键词（企业中文名称、企业英文名称、企业商业登记号码）
	HkEntityName string
}

type HKDataCreateOrderResp struct {
	Response[HKDataCreateOrderRespResult]
}

type HKDataCreateOrderRespResult struct {
	OrderNo string `json:"OrderNo"`
}

// HKDataCreateOrder 数据下单 https://openapi.qcc.com/dataApi/991
func (a *Api) HKDataCreateOrder(ctx context.Context, req *HKDataCreateOrderReq) (*HKDataCreateOrderResp, error) {
	var resp HKDataCreateOrderResp
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

	reply, err := c.SetResult(&resp).Get("/HKData/CreateOrder")
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

type HKDataGetDataReq struct {
	// 订单号
	OrderNo string
}

type HKDataGetDataResp struct {
	Response[HKDataGetDataRespResult]
}

type HKDataGetDataRespResult struct {
	DataStatus string                `json:"DataStatus"`
	Data       HKDataGetDataRespData `json:"Data"`
}

type HKDataGetDataRespData struct {
	Basic                          HKDataBasic                          `json:"basic"`
	CapitalStructureLiveSimplified HKDataCapitalStructureLiveSimplified `json:"capitalstructure_live_simplified"`
	DirectorsLive                  HKDataDirectorsLive                  `json:"directors_live"`
	Shareholders                   HKDataShareholders                   `json:"shareholders"`
	CompanySecretariesLive         HKDataCompanySecretariesLive         `json:"company_secretaries_live"`
}

type HKDataBasic struct {
	RetrievalTime string             `json:"RetrievalTime"`
	Details       HKDataBasicDetails `json:"Details"`
}

type HKDataBasicDetails struct {
	ID                         string               `json:"Id"`
	CompanyNumber              string               `json:"CompanyNumber"`
	BusinessRegistrationNumber string               `json:"BusinessRegistrationNumber"`
	CompanyNameChn             string               `json:"CompanyNameChn"`
	CompanyNameEng             string               `json:"CompanyNameEng"`
	CompanyType                string               `json:"CompanyType"`
	RegistrationDate           string               `json:"RegistrationDate"`
	Status                     string               `json:"Status"`
	Address                    string               `json:"Address"`
	AddressTrans               string               `json:"Address_Trans"`
	OfficeEffectiveDate        string               `json:"OfficeEffectiveDate"`
	WindingUpMode              string               `json:"WindingUpMode"`
	RegisterOfCharges          string               `json:"RegisterOfCharges"`
	DissolutionDate            string               `json:"DissolutionDate"`
	Important                  string               `json:"Important"`
	IncorporationPlace         string               `json:"IncorporationPlace"`
	Remarks                    string               `json:"Remarks"`
	OriginalNameList           []HKDataOriginalName `json:"OriginalNameList"`
}

type HKDataOriginalName struct {
	StartDate string `json:"StartDate"`
	EndDate   string `json:"EndDate"`
	Name      string `json:"Name"`
	EnName    string `json:"EnName"`
}

type HKDataCapitalStructureLiveSimplified struct {
	RetrievalTime string                            `json:"RetrievalTime"`
	Details       []HKDataCapitalStructureLiveEntry `json:"Details"`
}

type HKDataCapitalStructureLiveEntry struct {
	Currency                string `json:"Currency"`
	TotalAmount             string `json:"TotalAmount"`
	TotalAmountPaid         string `json:"TotalAmountPaid"`
	TotalAmountPaidCurrency string `json:"TotalAmountPaidCurrency"`
}

type HKDataDirectorsLive struct {
	RetrievalTime string               `json:"RetrievalTime"`
	Details       []HKDataDirectorLive `json:"Details"`
}

type HKDataDirectorLive struct {
	FullNameEng string `json:"FullNameEng"`
	FullNameChn string `json:"FullNameChn"`
	Position    string `json:"Position"`
	Type        string `json:"Type"`
}

type HKDataShareholders struct {
	Date     string              `json:"Date"`
	Details  []HKDataShareholder `json:"Details"`
	Verified string              `json:"verified"`
}

type HKDataShareholder struct {
	Address        string `json:"Address"`
	ClassOfShares  string `json:"ClassofShares"`
	FullName       string `json:"FullName"`
	FullNameChn    string `json:"FullNameChn"`
	FullNameEng    string `json:"FullNameEng"`
	NumberOfShares string `json:"NumberofShares"`
	PercentOfClass string `json:"PercentofClass"`
	TotalNumber    string `json:"TotalNumber"`
	Type           string `json:"Type"`
}

type HKDataCompanySecretariesLive struct {
	RetrievalTime string                       `json:"RetrievalTime"`
	Details       []HKDataCompanySecretaryLive `json:"Details"`
}

type HKDataCompanySecretaryLive struct {
	FullNameEng   string `json:"FullNameEng"`
	FullNameChn   string `json:"FullNameChn"`
	Type          string `json:"Type"`
	Address       string `json:"Address"`
	AppointedDate string `json:"AppointedDate"`
}

// HKDataGetData 数据获取 https://openapi.qcc.com/dataApi/991
func (a *Api) HKDataGetData(ctx context.Context, req *HKDataGetDataReq) (*HKDataGetDataResp, error) {
	var resp HKDataGetDataResp
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

	reply, err := c.SetResult(&resp).Get("/HKData/GetData")
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
