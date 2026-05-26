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

type PatentV4SearchReq struct {
	// 查询关键字
	SearchKey string
	// 国际专利分类号
	Ipc string
	// 发布开始时间（如“20220101”）
	PubDateBegin string
	// 发布结束时间（如“20220101”）
	PubDateEnd string
	// 每页数据条数，默认为10，最大50
	PageSize string
	// 页码，默认第1页
	PageIndex string
}

type PatentV4SearchResp struct {
	Response[[]PatentV4SearchRespResult]
}

type PatentV4SearchRespResult struct {
	ID                 string   `json:"Id"`
	IPCList            []string `json:"IPCList"`
	ApplicationNumber  string   `json:"ApplicationNumber"`
	ApplicationDate    string   `json:"ApplicationDate"`
	PublicationNumber  string   `json:"PublicationNumber"`
	PublicationDate    string   `json:"PublicationDate"`
	LegalStatusDesc    string   `json:"LegalStatusDesc"`
	Title              string   `json:"Title"`
	Agency             []string `json:"Agency"`
	KindCodeDesc       string   `json:"KindCodeDesc"`
	IPCDesc            []string `json:"IPCDesc"`
	InventorStringList []string `json:"InventorStringList"`
	AssigneestringList []string `json:"AssigneestringList"`
}

// PatentV4Search 专利多重查询 https://openapi.qcc.com/dataApi/514
func (a *Api) PatentV4Search(ctx context.Context, req *PatentV4SearchReq) (*PatentV4SearchResp, error) {
	var resp PatentV4SearchResp
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
	if req.Ipc != "" {
		c.SetQueryParam("ipc", req.Ipc)
	}
	if req.PubDateBegin != "" {
		c.SetQueryParam("pubDateBegin", req.PubDateBegin)
	}
	if req.PubDateEnd != "" {
		c.SetQueryParam("pubDateEnd", req.PubDateEnd)
	}
	if req.PageSize != "" {
		c.SetQueryParam("pageSize", req.PageSize)
	}
	if req.PageIndex != "" {
		c.SetQueryParam("pageIndex", req.PageIndex)
	}

	reply, err := c.SetResult(&resp).Get("/PatentV4/Search")
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

type PatentV4GetDetailsReq struct {
	// 专利主键
	ID string
}

type PatentV4GetDetailsResp struct {
	Response[PatentV4GetDetailsRespResult]
}

type PatentV4GetDetailsRespResult struct {
	DocumentTypes      string                                               `json:"DocumentTypes"`
	Agent              string                                               `json:"Agent"`
	LegalStatusDate    string                                               `json:"LegalStatusDate"`
	PrimaryExaminer    string                                               `json:"PrimaryExaminer"`
	AssiantExaminer    string                                               `json:"AssiantExaminer"`
	Cites              string                                               `json:"Cites"`
	OtherReferences    string                                               `json:"OtherReferences"`
	PatentImage        string                                               `json:"PatentImage"`
	IPCList            []string                                             `json:"IPCList"`
	ApplicationNumber  string                                               `json:"ApplicationNumber"`
	ApplicationDate    string                                               `json:"ApplicationDate"`
	PublicationNumber  string                                               `json:"PublicationNumber"`
	PublicationDate    string                                               `json:"PublicationDate"`
	LegalStatusDesc    string                                               `json:"LegalStatusDesc"`
	Title              string                                               `json:"Title"`
	Abstract           string                                               `json:"Abstract"`
	Agency             []string                                             `json:"Agency"`
	KindCodeDesc       string                                               `json:"KindCodeDesc"`
	IPCDesc            []string                                             `json:"IPCDesc"`
	InventorStringList []string                                             `json:"InventorStringList"`
	AssigneestringList []string                                             `json:"AssigneestringList"`
	PatentLegalHistory []PatentV4GetDetailsRespResultPatentLegalHistoryItem `json:"PatentLegalHistory"`
}

type PatentV4GetDetailsRespResultPatentLegalHistoryItem struct {
	Desc            string `json:"Desc"`
	LegalStatus     string `json:"LegalStatus"`
	LegalStatusDate string `json:"LegalStatusDate"`
}

// PatentV4GetDetails 专利详情查询 https://openapi.qcc.com/dataApi/514
func (a *Api) PatentV4GetDetails(ctx context.Context, req *PatentV4GetDetailsReq) (*PatentV4GetDetailsResp, error) {
	var resp PatentV4GetDetailsResp
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

	reply, err := c.SetResult(&resp).Get("/PatentV4/GetDetails")
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

type PatentV4SearchMultiPatentsReq struct {
	// 搜索关键词（统一社会信用代码、
	SearchKey string
	// 每页数据条数，默认为 10，最大
	PageSize int64
	// 页码，默认第 1 页
	PageIndex int64
	// 专利类型（1-发明公布，2-发明
	Kindcode string
	// 国际专利分类号
	Ipc string
	// 发布开始时间（如“20151020”、
	PubDateBegin string
	// 发布结束时间（如“20151020”、
	PubDateEnd string
	// 申请开始时间（如“20151020”、
	AppDateBegin string
	// 申请结束时间（如“20151020”、
	AppDateEnd string
}

type PatentV4SearchMultiPatentsResp struct {
	Response[PatentV4SearchMultiPatentsRespResult]
}

type PatentV4SearchMultiPatentsRespResult struct {
	ID                 string   `json:"Id"`
	LegalStatus        string   `json:"LegalStatus"`
	KindCode           string   `json:"KindCode"`
	IPCList            []string `json:"IPCList"`
	ApplicationNumber  string   `json:"ApplicationNumber"`
	ApplicationDate    string   `json:"ApplicationDate"`
	PublicationNumber  string   `json:"PublicationNumber"`
	PublicationDate    string   `json:"PublicationDate"`
	LegalStatusDesc    string   `json:"LegalStatusDesc"`
	Title              string   `json:"Title"`
	Agency             []string `json:"Agency"`
	KindCodeDesc       string   `json:"KindCodeDesc"`
	IPCDesc            []string `json:"IPCDesc"`
	InventorStringList []string `json:"InventorStringList"`
	AssigneestringList []string `json:"AssigneestringList"`
}

// PatentV4SearchMultiPatents 公司专利多重查询 https://openapi.qcc.com/dataApi/514
func (a *Api) PatentV4SearchMultiPatents(ctx context.Context, req *PatentV4SearchMultiPatentsReq) (*PatentV4SearchMultiPatentsResp, error) {
	var resp PatentV4SearchMultiPatentsResp
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
	if req.Kindcode != "" {
		c.SetQueryParam("kindcode", req.Kindcode)
	}
	if req.Ipc != "" {
		c.SetQueryParam("ipc", req.Ipc)
	}
	if req.PubDateBegin != "" {
		c.SetQueryParam("pubDateBegin", req.PubDateBegin)
	}
	if req.PubDateEnd != "" {
		c.SetQueryParam("pubDateEnd", req.PubDateEnd)
	}
	if req.AppDateBegin != "" {
		c.SetQueryParam("appDateBegin", req.AppDateBegin)
	}
	if req.AppDateEnd != "" {
		c.SetQueryParam("appDateEnd", req.AppDateEnd)
	}

	reply, err := c.SetResult(&resp).Get("/PatentV4/SearchMultiPatents")
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
