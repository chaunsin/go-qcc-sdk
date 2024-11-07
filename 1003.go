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

type BeneficialOwnerGetInfoReq struct {
	CompanyName string // 企业名称或统一社会信用代码
}

type BeneficialOwnerGetInfoResp struct {
	Response[BeneficialOwnerGetInfoRespResult]
}

type BeneficialOwnerGetInfoRespResult struct {
	CompanyName      string      `json:"CompanyName"`
	KeyNo            string      `json:"KeyNo"`
	FindMatched      string      `json:"FindMatched"`
	OperName         interface{} `json:"OperName"`
	BenifitType      interface{} `json:"BenifitType"`
	Position         interface{} `json:"Position"`
	Remark           interface{} `json:"Remark"`
	BreakThroughList []struct {
		Name              string `json:"Name"`
		KeyNo             string `json:"KeyNo"`
		TotalStockPercent string `json:"TotalStockPercent"`
		BenifitType       string `json:"BenifitType"`
		Position          string `json:"Position"`
		Calculation       string `json:"Calculation"`
		DetailInfoList    []struct {
			Level                    int     `json:"Level"`
			ShouldCapi               float64 `json:"ShouldCapi"`
			CapitalType              string  `json:"CapitalType"`
			BreakthroughStockPercent string  `json:"BreakthroughStockPercent"`
			StockType                string  `json:"StockType"`
			Path                     string  `json:"Path"`
			StockPercent             string  `json:"StockPercent"`
		} `json:"DetailInfoList"`
	} `json:"BreakThroughList"`
	Executives       []interface{} `json:"Executives"`
	ActualExecutives []struct {
		Name         string `json:"Name"`
		KeyNo        string `json:"KeyNo"`
		StockPercent string `json:"StockPercent"`
		BenifitType  string `json:"BenifitType"`
		Position     string `json:"Position"`
	} `json:"ActualExecutives"`
}

// BeneficialOwnerGetInfo 受益所有人 https://openapi.qcc.com/dataApi/1003
func (a *Api) BeneficialOwnerGetInfo(ctx context.Context, req *BeneficialOwnerGetInfoReq) (*BeneficialOwnerGetInfoResp, error) {
	var resp BeneficialOwnerGetInfoResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("companyName", req.CompanyName).
		SetResult(&resp).
		Get("https://api.qichacha.com/BeneficialOwner/GetInfo")
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
