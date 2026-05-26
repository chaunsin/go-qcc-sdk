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

type CompanyProductV4SearchCompanyCompanyProductsReq struct {
	// 搜索关键字（统一社会信用代码、企业名称）
	SearchKey string
	// 每页数据条数，默认为10，最大50
	PageSize string
	// 页码，默认第1页
	PageIndex string
}

type CompanyProductV4SearchCompanyCompanyProductsResp struct {
	Response[[]CompanyProductV4SearchCompanyCompanyProductsRespResult]
}

type CompanyProductV4SearchCompanyCompanyProductsRespResult struct {
	CompanyName string `json:"CompanyName"`
	Link        string `json:"Link"`
	ImageURL    string `json:"ImageUrl"`
	Name        string `json:"Name"`
	Domain      string `json:"Domain"`
	Tags        string `json:"Tags"`
	Description string `json:"Description"`
	Category    string `json:"Category"`
	ProductID   string `json:"ProductId"`
}

// CompanyProductV4SearchCompanyCompanyProducts 企业业务查询 https://openapi.qcc.com/dataApi/521
func (a *Api) CompanyProductV4SearchCompanyCompanyProducts(ctx context.Context, req *CompanyProductV4SearchCompanyCompanyProductsReq) (*CompanyProductV4SearchCompanyCompanyProductsResp, error) {
	var resp CompanyProductV4SearchCompanyCompanyProductsResp
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
	if req.PageSize != "" {
		c.SetQueryParam("pageSize", req.PageSize)
	}
	if req.PageIndex != "" {
		c.SetQueryParam("pageIndex", req.PageIndex)
	}

	reply, err := c.SetResult(&resp).Get("/CompanyProductV4/SearchCompanyCompanyProducts")
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
