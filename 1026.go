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

type ECIRecentPartnerGetListReq struct {
	SearchKey  string
	SearchType int64 // 查询类型，1-最新公示；2-工商登记，默认返回最新公示的信息
	PageSize   int64
	PageIndex  int64
}

type ECIRecentPartnerGetListResp struct {
	Response[ECIRecentPartnerGetListRespResult]
}

type ECIRecentPartnerGetListRespResult struct {
	VerifyResult int `json:"VerifyResult"`
	Data         struct {
		RegStockInfo struct {
			Paging struct {
				PageSize     int `json:"PageSize"`
				PageIndex    int `json:"PageIndex"`
				TotalRecords int `json:"TotalRecords"`
			} `json:"Paging"`
			RegStockList []struct {
				KeyNo        string `json:"KeyNo"`
				StockName    string `json:"StockName"`
				StockType    string `json:"StockType"`
				StockPercent string `json:"StockPercent"`
				ShouldCapi   string `json:"ShouldCapi"`
				ShoudDate    string `json:"ShoudDate"`
				StakeDate    string `json:"StakeDate"`
				CreditCode   string `json:"CreditCode"`
				Area         string `json:"Area"`
			} `json:"RegStockList"`
		} `json:"RegStockInfo"`
		PubStockInfo struct {
			Paging struct {
				PageSize     int `json:"PageSize"`
				PageIndex    int `json:"PageIndex"`
				TotalRecords int `json:"TotalRecords"`
			} `json:"Paging"`
			PubStockInList []struct {
				KeyNo        string `json:"KeyNo"`
				StockName    string `json:"StockName"`
				StockPercent string `json:"StockPercent"`
				HoldType     string `json:"HoldType"`
				Amount       string `json:"Amount"`
				CreditCode   string `json:"CreditCode"`
				Area         string `json:"Area"`
			} `json:"PubStockInList"`
		} `json:"PubStockInfo"`
	} `json:"Data"`
}

// ECIRecentPartnerGetList 股东信息(最新公示&工商登记) https://openapi.qcc.com/dataApi/1026
func (a *Api) ECIRecentPartnerGetList(ctx context.Context, req *ECIRecentPartnerGetListReq) (*ECIRecentPartnerGetListResp, error) {
	var resp ECIRecentPartnerGetListResp
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
	if req.SearchType > 0 {
		c.SetQueryParam("searchType", fmt.Sprintf("%d", req.SearchType))
	}
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}

	reply, err := c.SetResult(&resp).Get("https://api.qichacha.com/ECIRecentPartner/GetList")
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
