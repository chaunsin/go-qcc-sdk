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

type CourtAnnoDetailGetDetailReq struct {
	// 详情Id
	ID string
}

type CourtAnnoDetailGetDetailResp struct {
	Response[CourtAnnoDetailGetDetailRespResult]
}

type CourtAnnoDetailGetDetailRespResult struct {
	CaseReason   string                                            `json:"CaseReason"`
	CaseNo       string                                            `json:"CaseNo"`
	CourtTime    string                                            `json:"CourtTime"`
	Province     string                                            `json:"Province"`
	ScheduleTime string                                            `json:"ScheduleTime"`
	UndertakeDep string                                            `json:"UndertakeDep"`
	JusticeChief string                                            `json:"JusticeChief"`
	PartyList    []CourtAnnoDetailGetDetailRespResultPartyListItem `json:"PartyList"`
	Court        string                                            `json:"Court"`
	CourtRoom    string                                            `json:"CourtRoom"`
	Content      string                                            `json:"Content"`
	RoleList     []CourtAnnoDetailGetDetailRespResultRoleListItem  `json:"RoleList"`
}

type CourtAnnoDetailGetDetailRespResultPartyListItem struct {
	Name     string `json:"Name"`
	RoleType string `json:"RoleType"`
	Keyno    string `json:"Keyno"`
}

type CourtAnnoDetailGetDetailRespResultRoleListItem struct {
	RoleType     string                                                           `json:"RoleType"`
	RoleName     string                                                           `json:"RoleName"`
	RoleItemList []CourtAnnoDetailGetDetailRespResultRoleListItemRoleItemListItem `json:"RoleItemList"`
}

type CourtAnnoDetailGetDetailRespResultRoleListItemRoleItemListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

// CourtAnnoDetailGetDetail 开庭公告详情 https://openapi.qcc.com/dataApi/1057
func (a *Api) CourtAnnoDetailGetDetail(ctx context.Context, req *CourtAnnoDetailGetDetailReq) (*CourtAnnoDetailGetDetailResp, error) {
	var resp CourtAnnoDetailGetDetailResp
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

	reply, err := c.SetResult(&resp).Get("/CourtAnnoDetail/GetDetail")
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
