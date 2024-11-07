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

type PersonRiskCountCheckGetInfoReq struct {
	SearchKey  string
	PersonName string
}

type PersonRiskCountCheckGetInfoResp struct {
	Response[PersonRiskCountCheckGetInfoRespResult]
}

type PersonRiskCountCheckGetInfoRespResult struct {
	VerifyResult int `json:"VerifyResult"`
	Data         struct {
		Id       string `json:"Id"`
		SXCount  string `json:"SXCount"`
		ZXCount  string `json:"ZXCount"`
		STCount  string `json:"STCount"`
		ENCCount string `json:"ENCCount"`
		JUDCount string `json:"JUDCount"`
		CTACount string `json:"CTACount"`
		UNLCount string `json:"UNLCount"`
		STFCount string `json:"STFCount"`
		STPCount string `json:"STPCount"`
		PLECount string `json:"PLECount"`
		LACount  string `json:"LACount"`
		CSACount string `json:"CSACount"`
		DLNCount string `json:"DLNCount"`
		INECount string `json:"INECount"`
		LOCount  string `json:"LOCount"`
	} `json:"Data"`
}

// PersonRiskCountCheckGetInfo 董监高风险扫描 https://openapi.qcc.com/dataApi/763
func (a *Api) PersonRiskCountCheckGetInfo(ctx context.Context, req *PersonRiskCountCheckGetInfoReq) (*PersonRiskCountCheckGetInfoResp, error) {
	var resp PersonRiskCountCheckGetInfoResp
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
		SetQueryParam("personName", req.PersonName).
		SetResult(&resp).
		Get("https://api.qichacha.com/PersonRiskCountCheck/GetInfo")
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
