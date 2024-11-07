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

type ECIThreeElVerifyGetInfoReq struct {
	CreditCode  string
	CompanyName string
	OperName    string
}

type ECIThreeElVerifyGetInfoResp struct {
	Response[ECIThreeElVerifyGetInfoRespResult]
}

type ECIThreeElVerifyGetInfoRespResult struct {
	VerifyResult int64 `json:"VerifyResult"` // 0:公司编号有误，1:一致，2:企业名称不一致，3:法定代表人名称不一致
}

// ECIThreeElVerifyGetInfo 企业三要素核验 https://openapi.qcc.com/dataApi/856
func (a *Api) ECIThreeElVerifyGetInfo(ctx context.Context, req *ECIThreeElVerifyGetInfoReq) (*ECIThreeElVerifyGetInfoResp, error) {
	var resp ECIThreeElVerifyGetInfoResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("creditCode", req.CreditCode).
		SetQueryParam("companyName", req.CompanyName).
		SetQueryParam("operName", req.OperName).
		SetResult(&resp).
		Get("https://api.qichacha.com/AR/GetAnnualReport")
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
