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

type AdminPenaltyCheckGetListReq struct {
	SearchKey string
	PageSize  int64
	PageIndex int64
}

type AdminPenaltyCheckGetListResp struct {
	Response[AdminPenaltyCheckGetListRespResult]
}

type AdminPenaltyCheckGetListRespResult struct {
	VerifyResult int64                                    `json:"VerifyResult"`
	Data         []AdminPenaltyCheckGetListRespResultData `json:"Data"`
}

type AdminPenaltyCheckGetListRespResultData struct {
	Id           string `json:"Id"`
	KeyNo        string `json:"KeyNo"`
	PunishReason string `json:"PunishReason"`
	PunishResult string `json:"PunishResult"`
	PunishOffice string `json:"PunishOffice"`
	PunishDate   string `json:"PunishDate"`
	Source       string `json:"Source"`
	SourceCode   string `json:"SourceCode"`
	PunishAmt    string `json:"PunishAmt"`
}

// AdminPenaltyCheckGetList 行政处罚核查 https://openapi.qcc.com/dataApi/865
func (a *Api) AdminPenaltyCheckGetList(ctx context.Context, req *AdminPenaltyCheckGetListReq) (*AdminPenaltyCheckGetListResp, error) {
	var resp AdminPenaltyCheckGetListResp
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
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}

	reply, err := c.SetResult(&resp).Get("https://api.qichacha.com/AdminPenaltyCheck/GetList")
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
