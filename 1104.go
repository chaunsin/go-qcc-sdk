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

type CompanySumSaleGetInfoReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 月份（如“2021-01”）
	DataMonth string
}

type CompanySumSaleGetInfoResp struct {
	Response[CompanySumSaleGetInfoRespResult]
}

type CompanySumSaleGetInfoRespResult struct {
	VerifyResult int64                               `json:"VerifyResult"`
	Data         CompanySumSaleGetInfoRespResultData `json:"Data"`
}

type CompanySumSaleGetInfoRespResultData struct {
	DataMonth                   string `json:"DataMonth"`
	Creditcode                  string `json:"Creditcode"`
	CompanyName                 string `json:"CompanyName"`
	CompanySalesVolumeSumSix    string `json:"CompanySalesVolumeSumSix"`
	CompanySalesAmountSumSix    string `json:"CompanySalesAmountSumSix"`
	CompanySalesVolumeSumTwelve string `json:"CompanySalesVolumeSumTwelve"`
	CompanySalesAmountSumTwelve string `json:"CompanySalesAmountSumTwelve"`
	CompanySalesVolumeSumYoy    string `json:"CompanySalesVolumeSumYoy"`
	CompanySalesAmountSumYoy    string `json:"CompanySalesAmountSumYoy"`
	CompanySalesVolumeSum       string `json:"CompanySalesVolumeSum"`
	CompanySalesAmountSum       string `json:"CompanySalesAmountSum"`
}

// CompanySumSaleGetInfo 企业累计销售 https://openapi.qcc.com/dataApi/1104
func (a *Api) CompanySumSaleGetInfo(ctx context.Context, req *CompanySumSaleGetInfoReq) (*CompanySumSaleGetInfoResp, error) {
	var resp CompanySumSaleGetInfoResp
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
		SetQueryParam("dataMonth", req.DataMonth)

	reply, err := c.SetResult(&resp).Get("/CompanySumSale/GetInfo")
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
