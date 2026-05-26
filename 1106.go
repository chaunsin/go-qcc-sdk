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

type ShopInfoGetInfoReq struct {
	// 平台ID
	PlatformID string
	// 店铺ID
	ShopID string
	// 月份（如“2021-01”）
	DataMonth string
}

type ShopInfoGetInfoResp struct {
	Response[ShopInfoGetInfoRespResult]
}

type ShopInfoGetInfoRespResult struct {
	VerifyResult int64                         `json:"VerifyResult"`
	Data         ShopInfoGetInfoRespResultData `json:"Data"`
}

type ShopInfoGetInfoRespResultData struct {
	DataMonth            string `json:"DataMonth"`
	Shopid               string `json:"Shopid"`
	Shopname             string `json:"Shopname"`
	Platformid           string `json:"Platformid"`
	Platform             string `json:"Platform"`
	Creditcode           string `json:"Creditcode"`
	Companyname          string `json:"Companyname"`
	Shopitemnum          string `json:"Shopitemnum"`
	Isselfrun            string `json:"Isselfrun"`
	Shopurl              string `json:"Shopurl"`
	Shopopendate         string `json:"Shopopendate"`
	Shopopenyears        string `json:"Shopopenyears"`
	Shopsatatus          string `json:"Shopsatatus"`
	Shopclosedate        string `json:"Shopclosedate"`
	Mainbusiness         string `json:"Mainbusiness"`
	Statmainbusiness     string `json:"Statmainbusiness"`
	Businessscope        string `json:"Businessscope"`
	Companyprovince      string `json:"Companyprovince"`
	Companycity          string `json:"Companycity"`
	Companycounty        string `json:"Companycounty"`
	Companyareacode      string `json:"Companyareacode"`
	Shopprovince         string `json:"Shopprovince"`
	Shopcity             string `json:"Shopcity"`
	Shopcounty           string `json:"Shopcounty"`
	Shopareacode         string `json:"Shopareacode"`
	Shopdeliveryprovince string `json:"Shopdeliveryprovince"`
	Shopdeliverycity     string `json:"Shopdeliverycity"`
	Shopdeliverycounty   string `json:"Shopdeliverycounty"`
	Shopdeliverycode     string `json:"Shopdeliverycode"`
}

// ShopInfoGetInfo 店铺基本信息 https://openapi.qcc.com/dataApi/1106
func (a *Api) ShopInfoGetInfo(ctx context.Context, req *ShopInfoGetInfoReq) (*ShopInfoGetInfoResp, error) {
	var resp ShopInfoGetInfoResp
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
	if req.DataMonth != "" {
		c.SetQueryParam("dataMonth", req.DataMonth)
	}

	reply, err := c.SetResult(&resp).Get("/ShopInfo/GetInfo")
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
