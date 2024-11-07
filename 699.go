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

type IPOGetIPODetailReq struct {
	StockCode string
}

type IPOGetIPODetailResp struct {
	Response[IPOGetIPODetailRespResult]
}

type IPOGetIPODetailRespResult struct {
	BaseInfo struct {
		YesterdayClosePrice  string `json:"YesterdayClosePrice"`
		TodayOpenPrice       string `json:"TodayOpenPrice"`
		MarketDate           string `json:"MarketDate"`
		InternetReleasesDate string `json:"InternetReleasesDate"`
		StockCategory        string `json:"StockCategory"`
		Companyname          string `json:"Companyname"`
		EnglishName          string `json:"EnglishName"`
		HistoryName          string `json:"HistoryName"`
		RegCapital           string `json:"RegCapital"`
		Industry             string `json:"Industry"`
		RegNo                string `json:"RegNo"`
		PersonNumber         string `json:"PersonNumber"`
		ManagerNumber        string `json:"ManagerNumber"`
		ACode                string `json:"ACode"`
		AShortName           string `json:"AShortName"`
		BCode                string `json:"BCode"`
		BShortName           string `json:"BShortName"`
		HCode                string `json:"HCode"`
		HShortName           string `json:"HShortName"`
		LawFirm              string `json:"LawFirm"`
		AccountingFirm       string `json:"AccountingFirm"`
		OperName             string `json:"OperName"`
		ChairmanSecretary    string `json:"ChairmanSecretary"`
		Director             string `json:"Director"`
		GeneralManager       string `json:"GeneralManager"`
		StockAgent           string `json:"StockAgent"`
		IndependentDirector  string `json:"IndependentDirector"`
		TelNumber            string `json:"TelNumber"`
		Email                string `json:"Email"`
		Fax                  string `json:"Fax"`
		Website              string `json:"Website"`
		Area                 string `json:"Area"`
		ZipCode              string `json:"ZipCode"`
		OfficeAddress        string `json:"OfficeAddress"`
		RegAddress           string `json:"RegAddress"`
		PBR                  string `json:"PBR"`
		PER                  string `json:"PER"`
	} `json:"BaseInfo"`
	IPOPublishInfo struct {
		EstablishDate      string `json:"EstablishDate"`
		ReleasesRate       string `json:"ReleasesRate"`
		ReleasesType       string `json:"ReleasesType"`
		ParValue           string `json:"ParValue"`
		ReleasesAmount     string `json:"ReleasesAmount"`
		ParReleasesValue   string `json:"ParReleasesValue"`
		ReleasesCost       string `json:"ReleasesCost"`
		ReleasesTotal      string `json:"ReleasesTotal"`
		NetRaisedFunds     string `json:"NetRaisedFunds"`
		FirstDayOpenPrice  string `json:"FirstDayOpenPrice"`
		FirstDayClosePrice string `json:"FirstDayClosePrice"`
		FirstDayChangeRate string `json:"FirstDayChangeRate"`
		DownRate           string `json:"DownRate"`
		PricingSuccRate    string `json:"PricingSuccRate"`
	} `json:"IPOPublishInfo"`
	IpoTheme string `json:"IpoTheme"`
}

// IPOGetIPODetail 上市企业-企业简介 https://openapi.qcc.com/dataApi/699
func (a *Api) IPOGetIPODetail(ctx context.Context, req *IPOGetIPODetailReq) (*IPOGetIPODetailResp, error) {
	var resp IPOGetIPODetailResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("stockCode", req.StockCode).
		SetResult(&resp).
		Get("https://api.qichacha.com/IPO/GetIPODetail")
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

type IPOGetMainIndicatorReq struct {
	SearchKey string
}

type IPOGetMainIndicatorResp struct {
	Response[IPOGetIPODetailRespResult]
}

type IPOGetMainIndicatorResult struct {
	VerifyResult int `json:"VerifyResult"`
	Data         struct {
		ReportDate  []string `json:"ReportDate"`
		PrimaryList []struct {
			PrimaryDes    string `json:"PrimaryDes"`
			SecondaryList []struct {
				SecondaryDes       string   `json:"SecondaryDes"`
				SecondaryValueList []string `json:"SecondaryValueList"`
			} `json:"SecondaryList"`
		} `json:"PrimaryList"`
	} `json:"Data"`
}

// IPOGetMainIndicator 上市企业-主要指标 https://openapi.qcc.com/dataApi/699
func (a *Api) IPOGetMainIndicator(ctx context.Context, req *IPOGetMainIndicatorReq) (*IPOGetMainIndicatorResp, error) {
	var resp IPOGetMainIndicatorResp
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
		Get("https://api.qichacha.com/IPO/GetMainIndicator")
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

type IPOGetIPOToHolderReq struct {
	StockCode string
}

type IPOGetIPOToHolderResp struct {
	Response[IPOGetIPOToHolderResult]
}

type IPOGetIPOToHolderResult struct {
	PublicDate       string `json:"PublicDate"`
	IpoTopHolderList []struct {
		HolderName       string `json:"HolderName"`
		Type             string `json:"Type"`
		Count            string `json:"Count"`
		Proportion       string `json:"Proportion"`
		ChangeCount      string `json:"ChangeCount"`
		ChangeProportion string `json:"ChangeProportion"`
		CompanyKeyNo     string `json:"CompanyKeyNo"`
	} `json:"IpoTopHolderList"`
}

// IPOGetIPOToHolder 上市企业-十大股东 https://openapi.qcc.com/dataApi/699
func (a *Api) IPOGetIPOToHolder(ctx context.Context, req *IPOGetIPOToHolderReq) (*IPOGetIPOToHolderResp, error) {
	var resp IPOGetIPOToHolderResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("stockCode", req.StockCode).
		SetResult(&resp).
		Get("https://api.qichacha.com/IPO/GetIPOToHolder")
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

type IPOGetIPOExecutiveReq struct {
	StockCode string
}

type IPOGetIPOExecutiveResp struct {
	Response[[]IPOGetIPOToHolderResult]
}

type IPOGetIPOExecutiveResult struct {
	Name      string `json:"Name"`
	Sex       string `json:"Sex"`
	Age       string `json:"Age"`
	Education string `json:"Education"`
	Position  string `json:"Position"`
}

// IPOGetIPOExecutive 上市企业-企业高管 https://openapi.qcc.com/dataApi/699
func (a *Api) IPOGetIPOExecutive(ctx context.Context, req *IPOGetIPOExecutiveReq) (*IPOGetIPOExecutiveResp, error) {
	var resp IPOGetIPOExecutiveResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("stockCode", req.StockCode).
		SetResult(&resp).
		Get("https://api.qichacha.com/IPO/GetIPOExecutive")
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
