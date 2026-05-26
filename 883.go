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

type GroupMemberGetListReq struct {
	// 集团Id （可通过ApiCode 880接口获取）
	GroupID string
	// 页码，默认第1页
	PageIndex string
	// 每页数据条数，默认为10，最大20
	PageSize string
}

type GroupMemberGetListResp struct {
	Response[GroupMemberGetListRespResult]
}

type GroupMemberGetListRespResult struct {
	VerifyResult int64                                  `json:"VerifyResult"`
	Data         []GroupMemberGetListRespResultDataItem `json:"Data"`
}

type GroupMemberGetListRespResultDataItem struct {
	CompanyName           string                                   `json:"CompanyName"`
	CompanyKeyNo          string                                   `json:"CompanyKeyNo"`
	Oper                  GroupMemberGetListRespResultDataItemOper `json:"Oper"`
	RegistCapi            string                                   `json:"RegistCapi"`
	RegisteredCapital     string                                   `json:"RegisteredCapital"`
	RegisteredCapitalUnit string                                   `json:"RegisteredCapitalUnit"`
	RegisteredCapitalCCY  string                                   `json:"RegisteredCapitalCCY"`
	StartDate             string                                   `json:"StartDate"`
	Status                string                                   `json:"Status"`
	CreditCode            string                                   `json:"CreditCode"`
	CompanyLevel          string                                   `json:"CompanyLevel"`
}

type GroupMemberGetListRespResultDataItemOper struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

// GroupMemberGetList 集团成员 https://openapi.qcc.com/dataApi/883
func (a *Api) GroupMemberGetList(ctx context.Context, req *GroupMemberGetListReq) (*GroupMemberGetListResp, error) {
	var resp GroupMemberGetListResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("groupId", req.GroupID)
	if req.PageIndex != "" {
		c.SetQueryParam("pageIndex", req.PageIndex)
	}
	if req.PageSize != "" {
		c.SetQueryParam("pageSize", req.PageSize)
	}

	reply, err := c.SetResult(&resp).Get("/GroupMember/GetList")
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
