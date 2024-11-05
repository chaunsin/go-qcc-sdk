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

type GetCompanyGraphReq struct {
	SearchKey string
}

type GetCompanyGraphResp struct {
	Response[GetCompanyGraphRespData]
}

type GetCompanyGraphRespData struct {
	VerifyResult int `json:"VerifyResult"`
	Data         struct {
		KeyNo      string `json:"KeyNo"`
		Name       string `json:"Name"`
		RelateList []struct {
			NodeName     string `json:"NodeName"`
			NodeDataList []struct {
				Id      string  `json:"Id"`
				Name    string  `json:"Name"`
				Percent *string `json:"Percent"`
				Job     *string `json:"Job"`
			} `json:"NodeDataList"`
		} `json:"RelateList"`
	} `json:"Data"`
}

// GetCompanyGraph 企业图谱 https://openapi.qcc.com/dataApi/962
func (a *Api) GetCompanyGraph(ctx context.Context, req *GetCompanyGraphReq) (*GetCompanyGraphResp, error) {
	var resp GetCompanyGraphResp
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
		Get("https://api.qichacha.com/CompanyGraphCheck/GetInfo")
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
