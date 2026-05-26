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

type CaseFilingDetailGetDetailReq struct {
	// 详情Id
	ID string
}

type CaseFilingDetailGetDetailResp struct {
	Response[CaseFilingDetailGetDetailRespResult]
}

type CaseFilingDetailGetDetailRespResult struct {
	ID             string                                                  `json:"Id"`
	CaseReason     string                                                  `json:"CaseReason"`
	CaseNo         string                                                  `json:"CaseNo"`
	LianDate       string                                                  `json:"LianDate"`
	HoldDate       string                                                  `json:"HoldDate"`
	FinishDate     string                                                  `json:"FinishDate"`
	UndertakeDep   string                                                  `json:"UndertakeDep"`
	Court          string                                                  `json:"Court"`
	Judger         string                                                  `json:"Judger"`
	Assistant      string                                                  `json:"Assistant"`
	CaseType       string                                                  `json:"CaseType"`
	CaseStatus     string                                                  `json:"CaseStatus"`
	ProsecutorList []CaseFilingDetailGetDetailRespResultProsecutorListItem `json:"ProsecutorList"`
	DefendantList  []CaseFilingDetailGetDetailRespResultDefendantListItem  `json:"DefendantList"`
	OutSiderList   []CaseFilingDetailGetDetailRespResultOutSiderListItem   `json:"OutSiderList"`
	RoleList       []CaseFilingDetailGetDetailRespResultRoleListItem       `json:"RoleList"`
}

type CaseFilingDetailGetDetailRespResultProsecutorListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

type CaseFilingDetailGetDetailRespResultDefendantListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

type CaseFilingDetailGetDetailRespResultOutSiderListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

type CaseFilingDetailGetDetailRespResultRoleListItem struct {
	RoleType     string                                                            `json:"RoleType"`
	RoleName     string                                                            `json:"RoleName"`
	RoleItemList []CaseFilingDetailGetDetailRespResultRoleListItemRoleItemListItem `json:"RoleItemList"`
}

type CaseFilingDetailGetDetailRespResultRoleListItemRoleItemListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

// CaseFilingDetailGetDetail 立案信息详情 https://openapi.qcc.com/dataApi/1061
func (a *Api) CaseFilingDetailGetDetail(ctx context.Context, req *CaseFilingDetailGetDetailReq) (*CaseFilingDetailGetDetailResp, error) {
	var resp CaseFilingDetailGetDetailResp
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

	reply, err := c.SetResult(&resp).Get("/CaseFilingDetail/GetDetail")
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
