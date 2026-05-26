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

type CourtNoticeDetailGetDetailReq struct {
	// 法院公告Id
	ID string
}

type CourtNoticeDetailGetDetailResp struct {
	Response[CourtNoticeDetailGetDetailRespResult]
}

type CourtNoticeDetailGetDetailRespResult struct {
	Court          string                                                   `json:"Court"`
	CaseNo         string                                                   `json:"CaseNo"`
	Category       string                                                   `json:"Category"`
	PublishDate    string                                                   `json:"PublishDate"`
	PublishPage    string                                                   `json:"PublishPage"`
	Content        string                                                   `json:"Content"`
	ProsecutorList []CourtNoticeDetailGetDetailRespResultProsecutorListItem `json:"ProsecutorList"`
	DefendantList  []CourtNoticeDetailGetDetailRespResultDefendantListItem  `json:"DefendantList"`
	RoleList       []CourtNoticeDetailGetDetailRespResultRoleListItem       `json:"RoleList"`
}

type CourtNoticeDetailGetDetailRespResultProsecutorListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

type CourtNoticeDetailGetDetailRespResultDefendantListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

type CourtNoticeDetailGetDetailRespResultRoleListItem struct {
	RoleType     string                                                             `json:"RoleType"`
	RoleName     string                                                             `json:"RoleName"`
	RoleItemList []CourtNoticeDetailGetDetailRespResultRoleListItemRoleItemListItem `json:"RoleItemList"`
}

type CourtNoticeDetailGetDetailRespResultRoleListItemRoleItemListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

// CourtNoticeDetailGetDetail 法院公告详情 https://openapi.qcc.com/dataApi/1056
func (a *Api) CourtNoticeDetailGetDetail(ctx context.Context, req *CourtNoticeDetailGetDetailReq) (*CourtNoticeDetailGetDetailResp, error) {
	var resp CourtNoticeDetailGetDetailResp
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

	reply, err := c.SetResult(&resp).Get("/CourtNoticeDetail/GetDetail")
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
