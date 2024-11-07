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

type EquityThroughGetEquityThroughReq struct {
	Keyword string
	Level   string
}

type EquityThroughGetEquityThroughResp struct {
	Response[EquityThroughGetEquityThroughRespResult]
}

type EquityThroughGetEquityThroughRespResult struct {
	KeyNo    string `json:"KeyNo"`
	Name     string `json:"Name"`
	Count    string `json:"Count"`
	Children []struct {
		KeyNo               string `json:"KeyNo"`
		Name                string `json:"Name"`
		Category            string `json:"Category"`
		FundedRatio         string `json:"FundedRatio"`
		InParentActualRadio string `json:"InParentActualRadio"`
		Count               string `json:"Count"`
		Grade               string `json:"Grade"`
		ShouldCapi          string `json:"ShouldCapi"`
		StockRightNum       string `json:"StockRightNum"`
		Children            []struct {
			KeyNo               string      `json:"KeyNo"`
			Name                string      `json:"Name"`
			Category            string      `json:"Category"`
			FundedRatio         string      `json:"FundedRatio"`
			InParentActualRadio string      `json:"InParentActualRadio"`
			Count               string      `json:"Count"`
			Grade               string      `json:"Grade"`
			ShouldCapi          string      `json:"ShouldCapi"`
			StockRightNum       string      `json:"StockRightNum"`
			Children            interface{} `json:"Children"`
			ShortStatus         string      `json:"ShortStatus"`
		} `json:"Children"`
		ShortStatus string `json:"ShortStatus"`
	} `json:"Children"`
}

// EquityThroughGetEquityThrough 股权穿透(四层) https://openapi.qcc.com/dataApi/642
func (a *Api) EquityThroughGetEquityThrough(ctx context.Context, req *EquityThroughGetEquityThroughReq) (*EquityThroughGetEquityThroughResp, error) {
	var resp EquityThroughGetEquityThroughResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}

	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("keyWord", req.Keyword)
	if req.Level != "" {
		c.SetQueryParam("level", req.Level)
	}

	reply, err := c.SetResult(&resp).Get("https://api.qichacha.com/EquityThrough/GetEquityThrough")
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
