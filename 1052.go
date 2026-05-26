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

type ShixinDetailGetDetailReq struct {
	// 失信详情Id
	ID string
}

type ShixinDetailGetDetailResp struct {
	Response[ShixinDetailGetDetailRespResult]
}

type ShixinDetailGetDetailRespResult struct {
	ID            string                                             `json:"Id"`
	Name          string                                             `json:"Name"`
	Liandate      string                                             `json:"Liandate"`
	Anno          string                                             `json:"Anno"`
	Orgno         string                                             `json:"Orgno"`
	Executegov    string                                             `json:"Executegov"`
	Province      string                                             `json:"Province"`
	Executeunite  string                                             `json:"Executeunite"`
	Yiwu          string                                             `json:"Yiwu"`
	Executestatus string                                             `json:"Executestatus"`
	Actionremark  string                                             `json:"Actionremark"`
	Publicdate    string                                             `json:"Publicdate"`
	Executeno     string                                             `json:"Executeno"`
	NameKeyNoInfo []ShixinDetailGetDetailRespResultNameKeyNoInfoItem `json:"NameKeyNoInfo"`
	Amount        string                                             `json:"Amount"`
}

type ShixinDetailGetDetailRespResultNameKeyNoInfoItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

// ShixinDetailGetDetail 失信详情 https://openapi.qcc.com/dataApi/1052
func (a *Api) ShixinDetailGetDetail(ctx context.Context, req *ShixinDetailGetDetailReq) (*ShixinDetailGetDetailResp, error) {
	var resp ShixinDetailGetDetailResp
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

	reply, err := c.SetResult(&resp).Get("/ShixinDetail/GetDetail")
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
