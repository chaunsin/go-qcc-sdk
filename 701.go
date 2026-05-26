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

type CompanyNewsSearchNewsReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 情感状态（1-消极，2-中立，3-积极，默认为所有状态）
	EmotionType string
	// 开始日期（如“2019-01-01”）
	StartDate string
	// 结束日期（如“2019-01-01”）
	EndDate string
	// 每页数据条数，默认为10，最大50
	PageSize int64
	// 页码，默认第1页
	PageIndex int64
}

type CompanyNewsSearchNewsResp struct {
	Response[[]CompanyNewsSearchNewsRespResult]
}

type CompanyNewsSearchNewsRespResult struct {
	ID           string `json:"Id"`
	NewsID       string `json:"NewsId"`
	Source       string `json:"Source"`
	Title        string `json:"Title"`
	URL          string `json:"Url"`
	PublishTime  string `json:"PublishTime"`
	EmotionType  string `json:"EmotionType"`
	Category     string `json:"Category"`
	CategoryDesc string `json:"CategoryDesc"`
	NewsTags     string `json:"NewsTags"`
	Content      string `json:"Content"`
}

// CompanyNewsSearchNews 企业新闻搜索 https://openapi.qcc.com/dataApi/701
func (a *Api) CompanyNewsSearchNews(ctx context.Context, req *CompanyNewsSearchNewsReq) (*CompanyNewsSearchNewsResp, error) {
	var resp CompanyNewsSearchNewsResp
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
	if req.EmotionType != "" {
		c.SetQueryParam("emotionType", req.EmotionType)
	}
	if req.StartDate != "" {
		c.SetQueryParam("startDate", req.StartDate)
	}
	if req.EndDate != "" {
		c.SetQueryParam("endDate", req.EndDate)
	}
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}

	reply, err := c.SetResult(&resp).Get("/CompanyNews/SearchNews")
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

type CompanyNewsGetNewsDetailReq struct {
	// 新闻 Id（NewsId）
	ID string
}

type CompanyNewsGetNewsDetailResp struct {
	Response[CompanyNewsGetNewsDetailRespResult]
}

type CompanyNewsGetNewsDetailRespResult struct {
	ID      string `json:"Id"`
	Content string `json:"Content"`
}

// CompanyNewsGetNewsDetail 获取新闻详情 https://openapi.qcc.com/dataApi/701
func (a *Api) CompanyNewsGetNewsDetail(ctx context.Context, req *CompanyNewsGetNewsDetailReq) (*CompanyNewsGetNewsDetailResp, error) {
	var resp CompanyNewsGetNewsDetailResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("id", req.ID)

	reply, err := c.SetResult(&resp).Get("/CompanyNews/GetNewsDetail")
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
