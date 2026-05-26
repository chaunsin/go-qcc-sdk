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

type BusinessRealTraceCheckGetInfoReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
}

type BusinessRealTraceCheckGetInfoResp struct {
	Response[BusinessRealTraceCheckGetInfoRespResult]
}

type BusinessRealTraceCheckGetInfoRespResult struct {
	VerifyResult int64                                       `json:"VerifyResult"`
	Data         BusinessRealTraceCheckGetInfoRespResultData `json:"Data"`
}

type BusinessRealTraceCheckGetInfoRespResultData struct {
	OperateFlag   string                                                   `json:"OperateFlag"`
	OperateDetail BusinessRealTraceCheckGetInfoRespResultDataOperateDetail `json:"OperateDetail"`
}

type BusinessRealTraceCheckGetInfoRespResultDataOperateDetail struct {
	TMFLag        string `json:"TMFLag"`
	OpusCopyFlag  string `json:"OpusCopyFlag"`
	CertFlag      string `json:"CertFlag"`
	WebSiteFlag   string `json:"WebSiteFlag"`
	TaxCreFlag    string `json:"TaxCreFlag"`
	TenderFlag    string `json:"TenderFlag"`
	RecruitFlag   string `json:"RecruitFlag"`
	CreRatingFlag string `json:"CreRatingFlag"`
	ImpExFlag     string `json:"ImpExFlag"`
	AdLicFlag     string `json:"AdLicFlag"`
	TaxPayerFlag  string `json:"TaxPayerFlag"`
	DobCheckFlag  string `json:"DobCheckFlag"`
}

// BusinessRealTraceCheckGetInfo 企业经营真实痕迹核查 https://openapi.qcc.com/dataApi/961
func (a *Api) BusinessRealTraceCheckGetInfo(ctx context.Context, req *BusinessRealTraceCheckGetInfoReq) (*BusinessRealTraceCheckGetInfoResp, error) {
	var resp BusinessRealTraceCheckGetInfoResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("searchKey", req.SearchKey)

	reply, err := c.SetResult(&resp).Get("/BusinessRealTraceCheck/GetInfo")
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
