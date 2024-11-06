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

type ECISeniorPersonGetListReq struct {
	SearchKey  string
	PersonName string
	Type       int64 // 0：担任法定代表人；1：对外投资；2：在外任职
	PageIndex  int64
	PageSize   int64
}

type ECISeniorPersonGetListResp struct {
	Response[[]ECISeniorPersonGetListRespResult]
	GroupItems []struct {
		Key   string `json:"Key"`
		Items []struct {
			Value string `json:"Value"`
			Desc  string `json:"Desc"`
			Count string `json:"Count"`
		} `json:"Items"`
	} `json:"GroupItems"`
}

type ECISeniorPersonGetListRespResult struct {
	Area struct {
		City     string `json:"City"`
		County   string `json:"County"`
		Province string `json:"Province"`
	} `json:"Area"`
	Date     string `json:"Date"`
	EconKind string `json:"EconKind"`
	ImageUrl string `json:"ImageUrl"`
	Industry struct {
		Industry        string `json:"Industry"`
		IndustryCode    string `json:"IndustryCode"`
		SubIndustry     string `json:"SubIndustry"`
		SubIndustryCode string `json:"SubIndustryCode"`
	} `json:"Industry"`
	KeyNo        string `json:"KeyNo"`
	Name         string `json:"Name"`
	OperName     string `json:"OperName"`
	OperPersonId string `json:"OperPersonId"`
	OperType     string `json:"OperType"`
	RegCap       string `json:"RegCap"`
	RelationList []struct {
		Type     string `json:"Type"`
		TypeDesc string `json:"TypeDesc"`
		Value    string `json:"Value"`
	} `json:"RelationList"`
	SXCount    string `json:"SXCount"`
	Status     string `json:"Status"`
	ZXCount    string `json:"ZXCount"`
	CreditCode string `json:"CreditCode"`
}

// ECISeniorPersonGetList 企业人员董监高信息 https://openapi.qcc.com/dataApi/669
func (a *Api) ECISeniorPersonGetList(ctx context.Context, req *ECISeniorPersonGetListReq) (*ECISeniorPersonGetListResp, error) {
	var resp ECISeniorPersonGetListResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("searchKey", req.SearchKey).
		SetQueryParam("personName", req.PersonName).
		SetQueryParam("type", fmt.Sprintf("%d", req.Type))
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}

	reply, err := c.SetResult(&resp).Get("https://api.qichacha.com/ECISeniorPerson/GetList")
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
