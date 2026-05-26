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

type CompanyShopInfoGetInfoReq struct {
	SearchKey string
	DataMonth string
}

type CompanyShopInfoGetInfoResp struct {
	Response[CompanyShopInfoGetInfoRespResult]
}

type CompanyShopInfoGetInfoRespResult struct {
	VerifyResult int64                          `json:"VerifyResult"` // 验证结果，1:成功 0:失败
	Data         CompanyShopInfoGetInfoRespData `json:"Data"`
}

type CompanyShopInfoGetInfoRespData struct {
	CompanyShopNum string                              `json:"CompanyShopNum"` // 店铺数(家)
	ShopList       []CompanyShopInfoGetInfoRespShopRow `json:"ShopList"`
}

type CompanyShopInfoGetInfoRespShopRow struct {
	DataMonth   string `json:"DataMonth"`   // 月份，格式：yyyy-MM
	PlatformId  string `json:"PlatformId"`  // 平台ID
	Platform    string `json:"Platform"`    // 平台名称
	ShopId      string `json:"ShopId"`      // 店铺ID
	ShopName    string `json:"ShopName"`    // 店铺名称
	CreditCode  string `json:"CreditCode"`  // 统一社会信用代码
	CompanyName string `json:"CompanyName"` // 企业名称
}

// CompanyShopInfoGetInfo 企业店铺信息 https://openapi.qcc.com/dataApi/1102
func (a *Api) CompanyShopInfoGetInfo(ctx context.Context, req *CompanyShopInfoGetInfoReq) (*CompanyShopInfoGetInfoResp, error) {
	var resp CompanyShopInfoGetInfoResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("searchKey", req.SearchKey).
		SetQueryParam("dataMonth", req.DataMonth).
		SetResult(&resp).
		Get("/CompanyShopInfo/GetInfo")
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
