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

type ReportCreateReportReq struct {
	// 企业名称、统一社会信用代码
	KeyNo string
	// 报告格式（1-pdf，2-word，默认pdf）
	ReportFormat string
}

type ReportCreateReportResp struct {
	Response[ReportCreateReportRespResult]
}

type ReportCreateReportRespResult struct {
	OrderNo string `json:"OrderNo"`
}

// ReportCreateReport 报告下单 https://openapi.qcc.com/dataApi/260
func (a *Api) ReportCreateReport(ctx context.Context, req *ReportCreateReportReq) (*ReportCreateReportResp, error) {
	var resp ReportCreateReportResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key)
	c.SetQueryParam("keyNo", req.KeyNo)
	if req.ReportFormat != "" {
		c.SetQueryParam("reportFormat", req.ReportFormat)
	}

	reply, err := c.SetResult(&resp).Post("/Report/CreateReport")
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

type ReportGetReportInfoReq struct {
	// 报表订单号
	OrderNo string
}

type ReportGetReportInfoResp struct {
	Response[ReportGetReportInfoRespResult]
}

type ReportGetReportInfoRespResult struct {
	ReportStatus string `json:"ReportStatus"`
	ReportURL    string `json:"ReportUrl"`
}

// ReportGetReportInfo 获取报告信息 https://openapi.qcc.com/dataApi/260
func (a *Api) ReportGetReportInfo(ctx context.Context, req *ReportGetReportInfoReq) (*ReportGetReportInfoResp, error) {
	var resp ReportGetReportInfoResp
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

	reply, err := c.SetResult(&resp).Get("/Report/GetReportInfo")
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
