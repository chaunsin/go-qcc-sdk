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

type ECIV4GetBasicDetailsByNameReq struct {
	Keyword string
}

type ECIV4GetBasicDetailsByNameResp struct {
	Response[ECIV4GetBasicDetailsByNameRespResult]
}

type ECIV4GetBasicDetailsByNameRespResult struct {
	KeyNo        string        `json:"KeyNo"`
	Name         string        `json:"Name"`
	No           string        `json:"No"`
	BelongOrg    string        `json:"BelongOrg"`
	OperId       string        `json:"OperId"`
	OperName     string        `json:"OperName"`
	StartDate    string        `json:"StartDate"`
	EndDate      string        `json:"EndDate"`
	Status       string        `json:"Status"`
	Province     string        `json:"Province"`
	UpdatedDate  string        `json:"UpdatedDate"`
	CreditCode   string        `json:"CreditCode"`
	RegistCapi   string        `json:"RegistCapi"`
	EconKind     string        `json:"EconKind"`
	Address      string        `json:"Address"`
	Scope        string        `json:"Scope"`
	TermStart    string        `json:"TermStart"`
	TermEnd      string        `json:"TermEnd"`
	CheckDate    string        `json:"CheckDate"`
	OrgNo        interface{}   `json:"OrgNo"`
	IsOnStock    string        `json:"IsOnStock"`
	StockNumber  interface{}   `json:"StockNumber"`
	StockType    interface{}   `json:"StockType"`
	OriginalName []interface{} `json:"OriginalName"`
	ImageUrl     string        `json:"ImageUrl"`
	EntType      string        `json:"EntType"`
	RecCap       string        `json:"RecCap"`
	RevokeInfo   struct {
		CancelDate   string `json:"CancelDate"`
		CancelReason string `json:"CancelReason"`
		RevokeDate   string `json:"RevokeDate"`
		RevokeReason string `json:"RevokeReason"`
	} `json:"RevokeInfo"`
	Area struct {
		Province string `json:"Province"`
		City     string `json:"City"`
		County   string `json:"County"`
	} `json:"Area"`
	AreaCode string `json:"AreaCode"`
}

// ECIV4GetBasicDetailsByName 企业工商照面 https://openapi.qcc.com/dataApi/410
func (a *Api) ECIV4GetBasicDetailsByName(ctx context.Context, req *ECIV4GetBasicDetailsByNameReq) (*ECIV4GetBasicDetailsByNameResp, error) {
	var resp ECIV4GetBasicDetailsByNameResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("keyword", req.Keyword).
		SetResult(&resp).
		Get("https://api.qichacha.com/ECIV4/GetBasicDetailsByName")
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
