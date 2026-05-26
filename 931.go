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

type HistoryDNoticeCheckGetListReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 页码，默认第1页
	PageIndex string
	// 每页数据条数，默认为10，最大20
	PageSize string
}

type HistoryDNoticeCheckGetListResp struct {
	Response[HistoryDNoticeCheckGetListRespResult]
}

type HistoryDNoticeCheckGetListRespResult struct {
	VerifyResult int64                                          `json:"VerifyResult"`
	Data         []HistoryDNoticeCheckGetListRespResultDataItem `json:"Data"`
}

type HistoryDNoticeCheckGetListRespResultDataItem struct {
	ID             string                                                           `json:"Id"`
	Title          string                                                           `json:"Title"`
	CaseNo         string                                                           `json:"CaseNo"`
	CaseReason     string                                                           `json:"CaseReason"`
	CourtName      string                                                           `json:"CourtName"`
	PublishDate    string                                                           `json:"PublishDate"`
	ProsecutorList []HistoryDNoticeCheckGetListRespResultDataItemProsecutorListItem `json:"ProsecutorList"`
	DefendantList  []HistoryDNoticeCheckGetListRespResultDataItemDefendantListItem  `json:"DefendantList"`
	RoleList       []HistoryDNoticeCheckGetListRespResultDataItemRoleListItem       `json:"RoleList"`
}

type HistoryDNoticeCheckGetListRespResultDataItemProsecutorListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

type HistoryDNoticeCheckGetListRespResultDataItemDefendantListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

type HistoryDNoticeCheckGetListRespResultDataItemRoleListItem struct {
	RoleType     string                                                                     `json:"RoleType"`
	RoleName     string                                                                     `json:"RoleName"`
	RoleItemList []HistoryDNoticeCheckGetListRespResultDataItemRoleListItemRoleItemListItem `json:"RoleItemList"`
}

type HistoryDNoticeCheckGetListRespResultDataItemRoleListItemRoleItemListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

// HistoryDNoticeCheckGetList 历史送达公告核查 https://openapi.qcc.com/dataApi/931
func (a *Api) HistoryDNoticeCheckGetList(ctx context.Context, req *HistoryDNoticeCheckGetListReq) (*HistoryDNoticeCheckGetListResp, error) {
	var resp HistoryDNoticeCheckGetListResp
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

	reply, err := c.SetResult(&resp).Get("/HistoryDNoticeCheck/GetList")
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
