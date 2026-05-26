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

type CopyRightSearchCopyRightReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 每页数据条数，默认为10，最大50
	PageSize int64
	// 页码，默认第1页
	PageIndex int64
}

type CopyRightSearchCopyRightResp struct {
	Response[[]CopyRightSearchCopyRightRespResult]
}

type CopyRightSearchCopyRightRespResult struct {
	Category     string `json:"Category"`
	Name         string `json:"Name"`
	Owner        string `json:"Owner"`
	RegisterNo   string `json:"RegisterNo"`
	RegisterDate string `json:"RegisterDate"`
	FinishDate   string `json:"FinishDate"`
	PublishDate  string `json:"PublishDate"`
}

// CopyRightSearchCopyRight 著作权查询 https://openapi.qcc.com/dataApi/233
func (a *Api) CopyRightSearchCopyRight(ctx context.Context, req *CopyRightSearchCopyRightReq) (*CopyRightSearchCopyRightResp, error) {
	var resp CopyRightSearchCopyRightResp
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

	reply, err := c.SetResult(&resp).Get("/CopyRight/SearchCopyRight")
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

type CopyRightGetCopyRightReq struct {
	// 作品著作权人
	PersonName string
	// 作品名称
	ProductName string
	// 登记号
	RegisteNo string
	// 每页数据条数，默认为 10，最大
	PageSize int64
	// 页码，默认第 1 页
	PageIndex int64
}

type CopyRightGetCopyRightResp struct {
	Response[[]CopyRightGetCopyRightRespResult]
}

type CopyRightGetCopyRightRespResult struct {
	Category     string `json:"Category"`
	Name         string `json:"Name"`
	Owner        string `json:"Owner"`
	RegisterNo   string `json:"RegisterNo"`
	RegisterDate string `json:"RegisterDate"`
	FinishDate   string `json:"FinishDate"`
	PublishDate  string `json:"PublishDate"`
}

// CopyRightGetCopyRight 著作权多重查询 https://openapi.qcc.com/dataApi/233
func (a *Api) CopyRightGetCopyRight(ctx context.Context, req *CopyRightGetCopyRightReq) (*CopyRightGetCopyRightResp, error) {
	var resp CopyRightGetCopyRightResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	if req.PersonName != "" {
		c.SetQueryParam("personName", req.PersonName)
	}
	if req.ProductName != "" {
		c.SetQueryParam("productName", req.ProductName)
	}
	if req.RegisteNo != "" {
		c.SetQueryParam("registeNo", req.RegisteNo)
	}
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}

	reply, err := c.SetResult(&resp).Get("/CopyRight/GetCopyRight")
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

type CopyRightSearchSoftwareCrReq struct {
	// 搜索关键词（统一社会信用代码、
	SearchKey string
	// 每页数据条数，默认为 10，最大
	PageSize int64
	// 页码，默认第 1 页
	PageIndex int64
}

type CopyRightSearchSoftwareCrResp struct {
	Response[[]CopyRightSearchSoftwareCrRespResult]
}

type CopyRightSearchSoftwareCrRespResult struct {
	Category          string `json:"Category"`
	PublishDate       string `json:"PublishDate"`
	VersionNo         string `json:"VersionNo"`
	RegisterNo        string `json:"RegisterNo"`
	RegisterAperDate  string `json:"RegisterAperDate"`
	Name              string `json:"Name"`
	ShortName         string `json:"ShortName"`
	Owner             string `json:"Owner"`
	FinishDevelopDate string `json:"FinishDevelopDate"`
}

// CopyRightSearchSoftwareCr 软件著作权查询 https://openapi.qcc.com/dataApi/233
func (a *Api) CopyRightSearchSoftwareCr(ctx context.Context, req *CopyRightSearchSoftwareCrReq) (*CopyRightSearchSoftwareCrResp, error) {
	var resp CopyRightSearchSoftwareCrResp
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

	reply, err := c.SetResult(&resp).Get("/CopyRight/SearchSoftwareCr")
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

type CopyRightGetSoftwareCrReq struct {
	// 软件著作权人
	PersonName string
	// 软件全称
	FullName string
	// 软件简称
	ShortName string
	// 登记号
	RegisteNo string
	// 每页数据条数，默认为 10，最大
	PageSize int64
	// 页码，默认第 1 页
	PageIndex int64
}

type CopyRightGetSoftwareCrResp struct {
	Response[[]CopyRightGetSoftwareCrRespResult]
}

type CopyRightGetSoftwareCrRespResult struct {
	Category          string `json:"Category"`
	PublishDate       string `json:"PublishDate"`
	VersionNo         string `json:"VersionNo"`
	RegisterNo        string `json:"RegisterNo"`
	RegisterAperDate  string `json:"RegisterAperDate"`
	Name              string `json:"Name"`
	ShortName         string `json:"ShortName"`
	Owner             string `json:"Owner"`
	FinishDevelopDate string `json:"FinishDevelopDate"`
}

// CopyRightGetSoftwareCr 软件著作权多重查询 https://openapi.qcc.com/dataApi/233
func (a *Api) CopyRightGetSoftwareCr(ctx context.Context, req *CopyRightGetSoftwareCrReq) (*CopyRightGetSoftwareCrResp, error) {
	var resp CopyRightGetSoftwareCrResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	if req.PersonName != "" {
		c.SetQueryParam("personName", req.PersonName)
	}
	if req.FullName != "" {
		c.SetQueryParam("fullName", req.FullName)
	}
	if req.ShortName != "" {
		c.SetQueryParam("shortName", req.ShortName)
	}
	if req.RegisteNo != "" {
		c.SetQueryParam("registeNo", req.RegisteNo)
	}
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}

	reply, err := c.SetResult(&resp).Get("/CopyRight/GetSoftwareCr")
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
