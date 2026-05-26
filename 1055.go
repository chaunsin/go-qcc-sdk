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

type JudgmentDocDetailGetDetailReq struct {
	// 详情Id
	ID string
}

type JudgmentDocDetailGetDetailResp struct {
	Response[JudgmentDocDetailGetDetailRespResult]
}

type JudgmentDocDetailGetDetailRespResult struct {
	ID             string                                                   `json:"Id"`
	CaseName       string                                                   `json:"CaseName"`
	CaseNo         string                                                   `json:"CaseNo"`
	Court          string                                                   `json:"Court"`
	ContentClear   string                                                   `json:"ContentClear"`
	TrialRound     string                                                   `json:"TrialRound"`
	PublishDate    string                                                   `json:"PublishDate"`
	JudgeDate      string                                                   `json:"JudgeDate"`
	JudgeResult    string                                                   `json:"JudgeResult"`
	Defendantlist  []string                                                 `json:"Defendantlist"`
	Prosecutorlist []string                                                 `json:"Prosecutorlist"`
	RelatedComList []JudgmentDocDetailGetDetailRespResultRelatedComListItem `json:"RelatedComList"`
	RelatedLawList []JudgmentDocDetailGetDetailRespResultRelatedLawListItem `json:"RelatedLawList"`
	JudgeDocInfo   JudgmentDocDetailGetDetailRespResultJudgeDocInfo         `json:"JudgeDocInfo"`
}

type JudgmentDocDetailGetDetailRespResultRelatedComListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

type JudgmentDocDetailGetDetailRespResultRelatedLawListItem struct {
	KeyNo string `json:"KeyNo"`
	Name  string `json:"Name"`
}

type JudgmentDocDetailGetDetailRespResultJudgeDocInfo struct {
	PartyInfo               string `json:"PartyInfo"`
	TrialProcedure          string `json:"TrialProcedure"`
	CourtConsider           string `json:"CourtConsider"`
	PlaintiffRequest        string `json:"PlaintiffRequest"`
	DefendantReply          string `json:"DefendantReply"`
	CourtInspect            string `json:"CourtInspect"`
	PlaintiffRequestOfFirst string `json:"PlaintiffRequestOfFirst"`
	DefendantReplyOfFirst   string `json:"DefendantReplyOfFirst"`
	CourtInspectOfFirst     string `json:"CourtInspectOfFirst"`
	CourtConsiderOfFirst    string `json:"CourtConsiderOfFirst"`
	AppellantRequest        string `json:"AppellantRequest"`
	AppelleeArguing         string `json:"AppelleeArguing"`
	ExecuteProcess          string `json:"ExecuteProcess"`
	CollegiateBench         string `json:"CollegiateBench"`
	JudegeDate              string `json:"Judege_Date"`
	Recorder                string `json:"Recorder"`
}

// JudgmentDocDetailGetDetail 裁判文书详情 https://openapi.qcc.com/dataApi/1055
func (a *Api) JudgmentDocDetailGetDetail(ctx context.Context, req *JudgmentDocDetailGetDetailReq) (*JudgmentDocDetailGetDetailResp, error) {
	var resp JudgmentDocDetailGetDetailResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("id", req.ID)

	reply, err := c.SetResult(&resp).Get("/JudgmentDocDetail/GetDetail")
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
