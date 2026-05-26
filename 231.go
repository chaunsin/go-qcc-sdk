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
	"encoding/json"
	"fmt"
)

type TmSearchByApplicantReq struct {
	// 申请人名称
	Keyword string
	// 每页数据条数，默认为10，最大50
	PageSize int64
	// 页码，默认第1页
	PageIndex int64
}

type TmSearchByApplicantResp struct {
	Response[[]TmSearchByApplicantRespResult]
}

type TmSearchByApplicantRespResult struct {
	ID             string `json:"ID"`
	RegNo          string `json:"RegNo"`
	Name           string `json:"Name"`
	CategoryID     int64  `json:"CategoryId"`
	Category       string `json:"Category"`
	Person         string `json:"Person"`
	HasImage       bool   `json:"HasImage"`
	Flow           string `json:"Flow"`
	ImageURL       string `json:"ImageUrl"`
	FlowStatus     string `json:"FlowStatus"`
	FlowStatusDesc string `json:"FlowStatusDesc"`
	AppDate        string `json:"AppDate"`
	Status         string `json:"Status"`
}

// TmSearchByApplicant 商标搜索 https://openapi.qcc.com/dataApi/231
func (a *Api) TmSearchByApplicant(ctx context.Context, req *TmSearchByApplicantReq) (*TmSearchByApplicantResp, error) {
	var resp TmSearchByApplicantResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("keyword", req.Keyword)
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}

	reply, err := c.SetResult(&resp).Get("/tm/SearchByApplicant")
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

type TmGetDetailsReq struct {
	// 商标 Id
	ID string
}

type TmGetDetailsResp struct {
	Response[TmGetDetailsRespResult]
}

type TmGetDetailsRespResult struct {
	AddressCn         string            `json:"AddressCn"`
	AddressEn         string            `json:"AddressEn"`
	AnnouncementIssue string            `json:"AnnouncementIssue"`
	AnnouncementDate  string            `json:"AnnouncementDate"`
	Applicant1        string            `json:"Applicant1"`
	Applicant2        string            `json:"Applicant2"`
	Color             string            `json:"Color"`
	RegIssue          string            `json:"RegIssue"`
	RegDate           string            `json:"RegDate"`
	HouQiZhiDingDate  string            `json:"HouQiZhiDingDate"`
	GuoJiZhuCeDate    string            `json:"GuoJiZhuCeDate"`
	YouXianQuanDate   string            `json:"YouXianQuanDate"`
	ValidPeriod       string            `json:"ValidPeriod"`
	FlowItems         []json.RawMessage `json:"FlowItems"`
	ListGroupItems    []string          `json:"ListGroupItems"`
	ID                string            `json:"Id"`
	RegNo             string            `json:"RegNo"`
	IntCls            int64             `json:"IntCls"`
	Name              string            `json:"Name"`
	AppDate           string            `json:"AppDate"`
	ApplicantCn       string            `json:"ApplicantCn"`
	ApplicantEn       string            `json:"ApplicantEn"`
	Agent             string            `json:"Agent"`
	Status            int64             `json:"Status"`
	FlowStatus        string            `json:"FlowStatus"`
	FlowStatusDesc    string            `json:"FlowStatusDesc"`
	HasImage          bool              `json:"HasImage"`
	ImageURL          string            `json:"ImageUrl"`
	IsShare           string            `json:"IsShare"`
	TmType            string            `json:"TmType"`
	TmStyle           string            `json:"TmStyle"`
	SimilarGroups     string            `json:"SimilarGroups"`
}

// TmGetDetails 商标详情 https://openapi.qcc.com/dataApi/231
func (a *Api) TmGetDetails(ctx context.Context, req *TmGetDetailsReq) (*TmGetDetailsResp, error) {
	var resp TmGetDetailsResp
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

	reply, err := c.SetResult(&resp).Get("/tm/GetDetails")
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
