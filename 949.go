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

type AcctScanGetInfoReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
}

type AcctScanGetInfoResp struct {
	Response[AcctScanGetInfoRespResult]
}

type AcctScanGetInfoRespResult struct {
	CompanyName  string                              `json:"CompanyName"`
	KeyNo        string                              `json:"KeyNo"`
	VerifyResult string                              `json:"VerifyResult"`
	RiskLevel    string                              `json:"RiskLevel"`
	RiskScore    string                              `json:"RiskScore"`
	Data         []AcctScanGetInfoRespResultDataItem `json:"Data"`
}

type AcctScanGetInfoRespResultDataItem struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
	Passage     string `json:"Passage"`
	RiskType    string `json:"RiskType"`
}

// AcctScanGetInfo 准入尽职调查列表 https://openapi.qcc.com/dataApi/949
func (a *Api) AcctScanGetInfo(ctx context.Context, req *AcctScanGetInfoReq) (*AcctScanGetInfoResp, error) {
	var resp AcctScanGetInfoResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("searchKey", req.SearchKey)

	reply, err := c.SetResult(&resp).Get("/AcctScan/GetInfo")
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
