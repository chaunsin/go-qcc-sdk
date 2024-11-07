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

type ECIInvestmentThroughGetInfoReq struct {
	SearchKey          string
	Percent            string
	FilterAbnormalFlag int64 // 是否过滤已吊销、注销企业，0:不过滤；1:过滤 默认0
	PageSize           int64
	PageIndex          int64
}

type ECIInvestmentThroughGetInfoResp struct {
	Response[ECIInvestmentThroughGetInfoRespResult]
}

type ECIInvestmentThroughGetInfoRespResult struct {
	KeyNo            string      `json:"KeyNo"`
	CompanyName      string      `json:"CompanyName"`
	FindMatched      string      `json:"FindMatched"`
	Remark           interface{} `json:"Remark"`
	OrderNo          string      `json:"OrderNo"`
	BreakThroughList []struct {
		KeyNo             string `json:"KeyNo"`
		Name              string `json:"Name"`
		CreditCode        string `json:"CreditCode"`
		CorpStatus        string `json:"CorpStatus"`
		TotalStockPercent string `json:"TotalStockPercent"`
		DetailInfoList    []struct {
			Level                    string `json:"Level"`
			ShouldCapi               string `json:"ShouldCapi"`
			CapitalType              string `json:"CapitalType"`
			BreakThroughStockPercent string `json:"BreakThroughStockPercent"`
			StockType                string `json:"StockType"`
			Path                     string `json:"Path"`
			StockPercent             string `json:"StockPercent"`
		} `json:"DetailInfoList"`
	} `json:"BreakThroughList"`
}

// ECIInvestmentThroughGetInfo 企业对外投资穿透(十层) https://openapi.qcc.com/dataApi/663
func (a *Api) ECIInvestmentThroughGetInfo(ctx context.Context, req *ECIInvestmentThroughGetInfoReq) (*ECIInvestmentThroughGetInfoResp, error) {
	var resp ECIInvestmentThroughGetInfoResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("searchKey", req.SearchKey).
		SetQueryParam("percent", req.Percent)
	if req.FilterAbnormalFlag > 0 {
		c.SetQueryParam("filterAbnormalFlag", fmt.Sprintf("%d", req.FilterAbnormalFlag))
	}
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}

	reply, err := c.SetResult(&resp).Get("https://api.qichacha.com/ECIInvestmentThrough/GetInfo")
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
