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

type ActualControlSuspectedActualControlReq struct {
	Keyword string
}

type ActualControlSuspectedActualControlResp struct {
	Response[ActualControlSuspectedActualControlRespResult]
}

type ActualControlSuspectedActualControlRespResult struct {
	KeyNo          string `json:"KeyNo"`
	CompanyName    string `json:"CompanyName"`
	UpdateTime     string `json:"UpdateTime"`
	ControllerData struct {
		KeyNo        string `json:"KeyNo"`
		Name         string `json:"Name"`
		Percent      string `json:"Percent"`
		PercentTotal string `json:"PercentTotal"`
		Level        string `json:"Level"`
		PathCount    string `json:"PathCount"`
		Paths        [][]struct {
			KeyNo        string `json:"KeyNo"`
			Name         string `json:"Name"`
			Percent      string `json:"Percent"`
			PercentTotal string `json:"PercentTotal"`
			Level        string `json:"Level"`
		} `json:"Paths"`
		ControlPercent string `json:"ControlPercent"`
	} `json:"ControllerData"`
	ActualControl struct {
		PersonList []struct {
			Org   int    `json:"Org"`
			Name  string `json:"Name"`
			KeyNo string `json:"KeyNo"`
		} `json:"PersonList"`
		StockPercent string `json:"StockPercent"`
	} `json:"ActualControl"`
	ControllerDataList []struct {
		KeyNo       string `json:"KeyNo"`
		Name        string `json:"Name"`
		IsActual    string `json:"IsActual"`
		JudgeReason string `json:"JudgeReason"`
		ConInfo     *struct {
			VotePercent  string `json:"VotePercent"`
			PathInfoList []struct {
				PercentTotal string `json:"PercentTotal"`
				PathList     []struct {
					Name    string `json:"Name"`
					Level   string `json:"Level"`
					Percent string `json:"Percent"`
				} `json:"PathList"`
			} `json:"PathInfoList"`
		} `json:"ConInfo"`
		StockInfo interface{} `json:"StockInfo"`
	} `json:"ControllerDataList"`
}

// ActualControlSuspectedActualControl 实际控制人 https://openapi.qcc.com/dataApi/643
func (a *Api) ActualControlSuspectedActualControl(ctx context.Context, req *ActualControlSuspectedActualControlReq) (*ActualControlSuspectedActualControlResp, error) {
	var resp ActualControlSuspectedActualControlResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	reply, err := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("keyWord", req.Keyword).
		SetResult(&resp).
		Get("https://api.qichacha.com/ActualControl/SuspectedActualControl")
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
