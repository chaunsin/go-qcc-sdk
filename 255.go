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

type ECICertificationSearchCertificationReq struct {
	// 搜索关键字（公司名称）
	SearchKey string
	// 每页数据条数，默认为10，最大20
	PageSize int64
	// 页码，默认第1页
	PageIndex int64
	// 是否有效（0-无效，1-有效，2-未披露，默认为空）
	IsValid string
}

type ECICertificationSearchCertificationResp struct {
	Response[[]ECICertificationSearchCertificationRespResult]
}

type ECICertificationSearchCertificationRespResult struct {
	ID              string   `json:"Id"`
	Name            string   `json:"Name"`
	Type            string   `json:"Type"`
	StartDate       string   `json:"StartDate"`
	EndDate         string   `json:"EndDate"`
	No              string   `json:"No"`
	TypeDesc        string   `json:"TypeDesc"`
	InstitutionList []string `json:"InstitutionList"`
	Status          string   `json:"Status"`
}

// ECICertificationSearchCertification 资质证书 https://openapi.qcc.com/dataApi/255
func (a *Api) ECICertificationSearchCertification(ctx context.Context, req *ECICertificationSearchCertificationReq) (*ECICertificationSearchCertificationResp, error) {
	var resp ECICertificationSearchCertificationResp
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
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}
	if req.IsValid != "" {
		c.SetQueryParam("isValid", req.IsValid)
	}

	reply, err := c.SetResult(&resp).Get("/ECICertification/SearchCertification")
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

type ECICertificationGetCertificationDetailByIDReq struct {
	// 查询证书的 ID
	CertID string
}

type ECICertificationGetCertificationDetailByIDResp struct {
	Response[ECICertificationGetCertificationDetailByIDRespResult]
}

type ECICertificationGetCertificationDetailByIDRespResult struct {
	ID     string `json:"Id"`
	Data   any    `json:"Data"`
	Schema any    `json:"Schema"`
}

// ECICertificationGetCertificationDetailByID 资质证书详情 https://openapi.qcc.com/dataApi/255
func (a *Api) ECICertificationGetCertificationDetailByID(ctx context.Context, req *ECICertificationGetCertificationDetailByIDReq) (*ECICertificationGetCertificationDetailByIDResp, error) {
	var resp ECICertificationGetCertificationDetailByIDResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("certId", req.CertID)

	reply, err := c.SetResult(&resp).Get("/ECICertification/GetCertificationDetailById")
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
