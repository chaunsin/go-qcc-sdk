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

type ShopSumSaleGetInfoReq struct {
	// 平台ID
	PlatformID string
	// 店铺ID
	ShopID string
	// 月份（如“2021-01”）
	DataMonth string
}

type ShopSumSaleGetInfoResp struct {
	Response[ShopSumSaleGetInfoRespResult]
}

type ShopSumSaleGetInfoRespResult struct {
	VerifyResult int64                            `json:"VerifyResult"`
	Data         ShopSumSaleGetInfoRespResultData `json:"Data"`
}

type ShopSumSaleGetInfoRespResultData struct {
	DataMonth                string `json:"DataMonth"`
	PlatformID               string `json:"PlatformId"`
	Platform                 string `json:"Platform"`
	ShopID                   string `json:"ShopId"`
	ShopName                 string `json:"ShopName"`
	Creditcode               string `json:"Creditcode"`
	CompanyName              string `json:"CompanyName"`
	ShopSalesVolumeSumSix    string `json:"ShopSalesVolumeSumSix"`
	ShopSalesAmountSumSix    string `json:"ShopSalesAmountSumSix"`
	ShopSalesVolumeSumTwelve string `json:"ShopSalesVolumeSumTwelve"`
	ShopSalesAmountSumTwelve string `json:"ShopSalesAmountSumTwelve"`
	ShopSalesVolumeSumYoy    string `json:"ShopSalesVolumeSumYoy"`
	ShopSalesAmountSumYoy    string `json:"ShopSalesAmountSumYoy"`
	ShopSalesVolumeSum       string `json:"ShopSalesVolumeSum"`
	ShopSalesAmountSum       string `json:"ShopSalesAmountSum"`
}

// ShopSumSaleGetInfo 店铺累计销售 https://openapi.qcc.com/dataApi/1108
func (a *Api) ShopSumSaleGetInfo(ctx context.Context, req *ShopSumSaleGetInfoReq) (*ShopSumSaleGetInfoResp, error) {
	var resp ShopSumSaleGetInfoResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("platformId", req.PlatformID)
	c.SetQueryParam("shopId", req.ShopID)
	c.SetQueryParam("dataMonth", req.DataMonth)

	reply, err := c.SetResult(&resp).Get("/ShopSumSale/GetInfo")
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
