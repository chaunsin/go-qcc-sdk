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

type TenderCheckGetListReq struct {
	// 搜索关键词（支持项目名称、编号、关键字，可以添加多个关键词，多关键词用空格隔离，查询时默认同时满足，单个关键词字符长度不超过100）
	Keyword string
	// 行政区域编码（仅支持省、市查询，如"320500"）
	AreaCode string
	// 信息类型代码（3-招标公告，4-中标公告）
	MsgType string
	// 发布日期起始日期（如“2022-01-01”，最小日期为当前日期-2年，最大日期为当前日期）
	PubDateStart string
	// 发布日期结束日期（如“2022-01-01”，最小日期为当前日期-2年，最大日期为当前日期）
	PubDateEnd string
	// 页码，默认第1页
	PageIndex string
	// 每页数据条数，默认为10，最大20
	PageSize string
}

type TenderCheckGetListResp struct {
	Response[TenderCheckGetListRespResult]
}

type TenderCheckGetListRespResult struct {
	VerifyResult int64                                  `json:"VerifyResult"`
	Data         []TenderCheckGetListRespResultDataItem `json:"Data"`
}

type TenderCheckGetListRespResultDataItem struct {
	ID              string                                                    `json:"Id"`
	Title           string                                                    `json:"Title"`
	ProjectNo       string                                                    `json:"ProjectNo"`
	ChannelName     string                                                    `json:"ChannelName"`
	Province        string                                                    `json:"Province"`
	City            string                                                    `json:"City"`
	IndustryDesc    string                                                    `json:"IndustryDesc"`
	BudgetAmt       string                                                    `json:"BudgetAmt"`
	PublishDate     string                                                    `json:"PublishDate"`
	OpenDate        string                                                    `json:"OpenDate"`
	ObtainEndDate   string                                                    `json:"ObtainEndDate"`
	BidInviUnitList []TenderCheckGetListRespResultDataItemBidInviUnitListItem `json:"BidInviUnitList"`
	WinBidUnitList  []TenderCheckGetListRespResultDataItemWinBidUnitListItem  `json:"WinBidUnitList"`
	AgentUnitList   []TenderCheckGetListRespResultDataItemAgentUnitListItem   `json:"AgentUnitList"`
	BidProgressList []string                                                  `json:"BidProgressList"`
	BidEndDate      string                                                    `json:"BidEndDate"`
	ContentURL      string                                                    `json:"ContentUrl"`
	ContractEndTime string                                                    `json:"ContractEndTime"`
}

type TenderCheckGetListRespResultDataItemBidInviUnitListItem struct {
	KeyNo   string `json:"KeyNo"`
	Name    string `json:"Name"`
	Contact string `json:"Contact"`
	TelNo   string `json:"TelNo"`
}

type TenderCheckGetListRespResultDataItemWinBidUnitListItem struct {
	WinBidAmt string `json:"WinBidAmt"`
	KeyNo     string `json:"KeyNo"`
	Name      string `json:"Name"`
	Contact   string `json:"Contact"`
	TelNo     string `json:"TelNo"`
}

type TenderCheckGetListRespResultDataItemAgentUnitListItem struct {
	KeyNo   string `json:"KeyNo"`
	Name    string `json:"Name"`
	Contact string `json:"Contact"`
	TelNo   string `json:"TelNo"`
}

// TenderCheckGetList 招投标搜索 https://openapi.qcc.com/dataApi/958
func (a *Api) TenderCheckGetList(ctx context.Context, req *TenderCheckGetListReq) (*TenderCheckGetListResp, error) {
	var resp TenderCheckGetListResp
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
	if req.AreaCode != "" {
		c.SetQueryParam("areaCode", req.AreaCode)
	}
	if req.MsgType != "" {
		c.SetQueryParam("msgType", req.MsgType)
	}
	if req.PubDateStart != "" {
		c.SetQueryParam("pubDateStart", req.PubDateStart)
	}
	if req.PubDateEnd != "" {
		c.SetQueryParam("pubDateEnd", req.PubDateEnd)
	}
	if req.PageIndex != "" {
		c.SetQueryParam("pageIndex", req.PageIndex)
	}
	if req.PageSize != "" {
		c.SetQueryParam("pageSize", req.PageSize)
	}

	reply, err := c.SetResult(&resp).Get("/TenderCheck/GetList")
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

type TenderCheckGetDetailReq struct {
	// 招投标详情 id
	ID string
}

type TenderCheckGetDetailResp struct {
	Response[TenderCheckGetDetailRespResult]
}

type TenderCheckGetDetailRespResult struct {
	ContentURL  string                             `json:"ContentUrl"`
	PublishDate string                             `json:"PublishDate"`
	Title       string                             `json:"Title"`
	Data        TenderCheckGetDetailRespResultData `json:"Data"`
}

type TenderCheckGetDetailRespResultData struct {
	Content string `json:"Content"`
}

// TenderCheckGetDetail 招投标详情 https://openapi.qcc.com/dataApi/958
func (a *Api) TenderCheckGetDetail(ctx context.Context, req *TenderCheckGetDetailReq) (*TenderCheckGetDetailResp, error) {
	var resp TenderCheckGetDetailResp
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

	reply, err := c.SetResult(&resp).Get("/TenderCheck/GetDetail")
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
