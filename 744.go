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

type JudicialSaleCheckGetListReq struct {
	SearchKey string
	PageSize  int64
	PageIndex int64
}

type JudicialSaleCheckGetListResp struct {
	Response[JudicialSaleCheckGetListRespResult]
}

type JudicialSaleCheckGetListRespResult struct {
	VerifyResult int `json:"VerifyResult"`
	Data         []struct {
		Id              string `json:"Id"`
		Name            string `json:"Name"`
		Executegov      string `json:"Executegov"`
		ActionRemark    string `json:"ActionRemark"`
		YiWu            string `json:"YiWu"`
		EvaluationPrice string `json:"EvaluationPrice"`
	} `json:"Data"`
}

// JudicialSaleCheckGetList 司法拍卖核查 https://openapi.qcc.com/dataApi/744
func (a *Api) JudicialSaleCheckGetList(ctx context.Context, req *JudicialSaleCheckGetListReq) (*JudicialSaleCheckGetListResp, error) {
	var resp JudicialSaleCheckGetListResp
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
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}

	reply, err := c.SetResult(&resp).Get("https://api.qichacha.com/JudicialSaleCheck/GetList")
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
