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

type InternationalPatentCheckGetListReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 页码，默认第1页
	PageIndex string
	// 每页数据条数，默认为10，最大50
	PageSize string
}

type InternationalPatentCheckGetListResp struct {
	Response[InternationalPatentCheckGetListRespResult]
}

type InternationalPatentCheckGetListRespResult struct {
	VerifyResult int64                                               `json:"VerifyResult"`
	Data         []InternationalPatentCheckGetListRespResultDataItem `json:"Data"`
}

type InternationalPatentCheckGetListRespResultDataItem struct {
	ID                string `json:"Id"`
	Title             string `json:"Title"`
	ApplicationNumber string `json:"ApplicationNumber"`
	ApplicationDate   string `json:"ApplicationDate"`
	PublicationNumber string `json:"PublicationNumber"`
	PublicationDate   string `json:"PublicationDate"`
	Inventors         string `json:"Inventors"`
	LegalStatus       string `json:"LegalStatus"`
}

// InternationalPatentCheckGetList 专利列表 https://openapi.qcc.com/dataApi/997
func (a *Api) InternationalPatentCheckGetList(ctx context.Context, req *InternationalPatentCheckGetListReq) (*InternationalPatentCheckGetListResp, error) {
	var resp InternationalPatentCheckGetListResp
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
	if req.PageIndex != "" {
		c.SetQueryParam("pageIndex", req.PageIndex)
	}
	if req.PageSize != "" {
		c.SetQueryParam("pageSize", req.PageSize)
	}

	reply, err := c.SetResult(&resp).Get("/InternationalPatentCheck/GetList")
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

type InternationalPatentCheckGetDetailReq struct {
	// 专利 id
	ID string
}

type InternationalPatentCheckGetDetailResp struct {
	Response[InternationalPatentCheckGetDetailRespResult]
}

type InternationalPatentCheckGetDetailRespResult struct {
	BasicInfo          InternationalPatentCheckGetDetailRespResultBasicInfo `json:"BasicInfo"`
	Requirement        string                                               `json:"Requirement"`
	Instructions       string                                               `json:"Instructions"`
	AbstractImage      string                                               `json:"AbstractImage"`
	InstructionImgList []string                                             `json:"InstructionImgList"`
}

type InternationalPatentCheckGetDetailRespResultBasicInfo struct {
	ApplicationNumber   string                                                                 `json:"ApplicationNumber"`
	ApplicationDate     string                                                                 `json:"ApplicationDate"`
	PublicationNumber   string                                                                 `json:"PublicationNumber"`
	PublicationDate     string                                                                 `json:"PublicationDate"`
	PriorityCode        string                                                                 `json:"PriorityCode"`
	PriorityDate        string                                                                 `json:"PriorityDate"`
	IPCList             []string                                                               `json:"IPCList"`
	CPCList             []string                                                               `json:"CPCList"`
	PatenteeList        []InternationalPatentCheckGetDetailRespResultBasicInfoPatenteeListItem `json:"PatenteeList"`
	InventorList        []string                                                               `json:"InventorList"`
	PatenteeAddressList []string                                                               `json:"PatenteeAddressList"`
	AgencyList          []string                                                               `json:"AgencyList"`
	Agent               string                                                                 `json:"Agent"`
	Abstract            string                                                                 `json:"Abstract"`
}

type InternationalPatentCheckGetDetailRespResultBasicInfoPatenteeListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

// InternationalPatentCheckGetDetail 专利详情 https://openapi.qcc.com/dataApi/997
func (a *Api) InternationalPatentCheckGetDetail(ctx context.Context, req *InternationalPatentCheckGetDetailReq) (*InternationalPatentCheckGetDetailResp, error) {
	var resp InternationalPatentCheckGetDetailResp
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

	reply, err := c.SetResult(&resp).Get("/InternationalPatentCheck/GetDetail")
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
