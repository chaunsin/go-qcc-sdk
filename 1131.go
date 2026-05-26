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

type HKDocCreateOrderReq struct {
	// 搜索关键词（企业中文名称、企业英文名称、企业商业登记号码）
	SearchKey string
	// 公告编号（仅支持可供查阅的公告）
	DocNumber string
}

type HKDocCreateOrderResp struct {
	Response[HKDocCreateOrderRespResult]
}

type HKDocCreateOrderRespResult struct {
	OrderNo string `json:"OrderNo"`
}

// HKDocCreateOrder 公告下单 https://openapi.qcc.com/dataApi/1131
func (a *Api) HKDocCreateOrder(ctx context.Context, req *HKDocCreateOrderReq) (*HKDocCreateOrderResp, error) {
	var resp HKDocCreateOrderResp
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
	c.SetQueryParam("docNumber", req.DocNumber)

	reply, err := c.SetResult(&resp).Get("/HKDoc/CreateOrder")
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

type HKDocGetDataReq struct {
	// 订单号（仅支持 15 天内的订单数
	OrderNo string
}

type HKDocGetDataResp struct {
	Response[HKDocGetDataRespResult]
}

type HKDocGetDataRespResult struct {
	DataStatus   string `json:"DataStatus"`
	OriginalFile string `json:"OriginalFile"`
}

// HKDocGetData 数据获取 https://openapi.qcc.com/dataApi/1131
func (a *Api) HKDocGetData(ctx context.Context, req *HKDocGetDataReq) (*HKDocGetDataResp, error) {
	var resp HKDocGetDataResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("orderNo", req.OrderNo)

	reply, err := c.SetResult(&resp).Get("/HKDoc/GetData")
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
