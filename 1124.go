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

type TaxDataCreateOrderReq struct {
	// 搜索关键词（统一社会信用代码、企业名称）
	SearchKey string
	// 税务局用户名（SM4加密，密钥为用户key）
	UserName string
	// 税务局密码（SM4加密，密钥为用户key）
	Password string
}

type TaxDataCreateOrderResp struct {
	Response[TaxDataCreateOrderRespResult]
}

type TaxDataCreateOrderRespResult struct {
	DataStatus  string `json:"DataStatus"`
	OrderNo     string `json:"OrderNo"`
	OrderResult string `json:"OrderResult"`
}

// TaxDataCreateOrder 数据下单 https://openapi.qcc.com/dataApi/1124
func (a *Api) TaxDataCreateOrder(ctx context.Context, req *TaxDataCreateOrderReq) (*TaxDataCreateOrderResp, error) {
	var resp TaxDataCreateOrderResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	body := map[string]any{}
	body["searchKey"] = req.SearchKey
	body["userName"] = req.UserName
	body["password"] = req.Password
	c.SetBody(body)

	reply, err := c.SetResult(&resp).Post("/TaxData/CreateOrder")
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

type TaxDataSendCodeReq struct {
	// 订单号
	OrderNo string
	// 验证码
	VerifyCode string
}

type TaxDataSendCodeResp struct {
	Response[TaxDataSendCodeRespResult]
}

type TaxDataSendCodeRespResult struct {
	DataStatus  string `json:"DataStatus"`
	OrderNo     string `json:"OrderNo"`
	OrderResult string `json:"OrderResult"`
}

// TaxDataSendCode 验证码发送 https://openapi.qcc.com/dataApi/1124
func (a *Api) TaxDataSendCode(ctx context.Context, req *TaxDataSendCodeReq) (*TaxDataSendCodeResp, error) {
	var resp TaxDataSendCodeResp
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
	c.SetQueryParam("verifyCode", req.VerifyCode)

	reply, err := c.SetResult(&resp).Get("/TaxData/SendCode")
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

type TaxDataGetDataReq struct {
	// 订单号（有效期 7 天）
	OrderNo string
}

type TaxDataGetDataResp struct {
	Response[TaxDataGetDataRespResult]
}

type TaxDataGetDataRespResult struct {
	DataStatus      string `json:"DataStatus"`
	Data            any    `json:"Data"`
	TaxBurdenRateLi []any  `json:"TaxBurdenRateLi"`
	SaleList        []any  `json:"SaleList"`
	TaxData         any    `json:"TaxData"`
	DeclarationDeta any    `json:"DeclarationDeta"`
	CollectionDetai any    `json:"CollectionDetai"`
	FinancialIndexL []any  `json:"FinancialIndexL"`
	FinancialList   []any  `json:"FinancialList"`
	BreakLawDetailL []any  `json:"BreakLawDetailL"`
	BreakLawSummary []any  `json:"BreakLawSummary"`
	SupplierCustome []any  `json:"SupplierCustome"`
	TopCustomerList []any  `json:"TopCustomerList"`
	TopSupplierList []any  `json:"TopSupplierList"`
	ExpenseDetail   any    `json:"ExpenseDetail"`
	CashFlowList    []any  `json:"CashFlowList"`
}

// TaxDataGetData 数据获取 https://openapi.qcc.com/dataApi/1124
func (a *Api) TaxDataGetData(ctx context.Context, req *TaxDataGetDataReq) (*TaxDataGetDataResp, error) {
	var resp TaxDataGetDataResp
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

	reply, err := c.SetResult(&resp).Get("/TaxData/GetData")
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
