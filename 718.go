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

type RecruitmentGetListReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 页码
	PageIndex string
	// 每页数据条数，默认为10，最大50
	PageSize string
}

type RecruitmentGetListResp struct {
	Response[[]RecruitmentGetListRespResult]
}

type RecruitmentGetListRespResult struct {
	ID           string `json:"Id"`
	KeyNo        string `json:"KeyNo"`
	CompanyName  string `json:"CompanyName"`
	Title        string `json:"Title"`
	PublishDate  string `json:"PublishDate"`
	Salary       string `json:"Salary"`
	Province     string `json:"Province"`
	ProvinceDesc string `json:"ProvinceDesc"`
	City         string `json:"City"`
	Experience   string `json:"Experience"`
	Education    string `json:"Education"`
}

// RecruitmentGetList 企业招聘列表 https://openapi.qcc.com/dataApi/718
func (a *Api) RecruitmentGetList(ctx context.Context, req *RecruitmentGetListReq) (*RecruitmentGetListResp, error) {
	var resp RecruitmentGetListResp
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
	if req.PageIndex != "" {
		c.SetQueryParam("pageIndex", req.PageIndex)
	}
	if req.PageSize != "" {
		c.SetQueryParam("pageSize", req.PageSize)
	}

	reply, err := c.SetResult(&resp).Get("/Recruitment/GetList")
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

type RecruitmentGetDetailReq struct {
	// 详情 Id
	ID string
}

type RecruitmentGetDetailResp struct {
	Response[RecruitmentGetDetailRespResult]
}

type RecruitmentGetDetailRespResult struct {
	CompanyName string `json:"CompanyName"`
	KeyNo       string `json:"KeyNo"`
	Title       string `json:"Title"`
	Source      string `json:"Source"`
	URL         string `json:"Url"`
	Area        string `json:"Area"`
	Salary      string `json:"Salary"`
	Description string `json:"Description"`
	Education   string `json:"Education"`
	Experience  string `json:"Experience"`
}

// RecruitmentGetDetail 企业招聘详情 https://openapi.qcc.com/dataApi/718
func (a *Api) RecruitmentGetDetail(ctx context.Context, req *RecruitmentGetDetailReq) (*RecruitmentGetDetailResp, error) {
	var resp RecruitmentGetDetailResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("id", req.ID)

	reply, err := c.SetResult(&resp).Get("/Recruitment/GetDetail")
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
