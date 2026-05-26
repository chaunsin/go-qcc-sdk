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

type ECICertificationSearchCertificationReq struct {
	// 搜索关键字（公司名称）
	SearchKey string
	// 每页数据条数，默认为10，最大20
	PageSize int64
	// 页码，默认第1页
	PageIndex int64
	// 是否有效（0-无效，1-有效，2-未披露，默认为空）
	IsValid string
}

type ECICertificationSearchCertificationResp struct {
	Response[[]ECICertificationSearchCertificationRespResult]
}

type ECICertificationSearchCertificationRespResult struct {
	ID              string   `json:"Id"`
	Name            string   `json:"Name"`
	Type            string   `json:"Type"`
	StartDate       string   `json:"StartDate"`
	EndDate         string   `json:"EndDate"`
	No              string   `json:"No"`
	TypeDesc        string   `json:"TypeDesc"`
	InstitutionList []string `json:"InstitutionList"`
	Status          string   `json:"Status"`
}

// ECICertificationSearchCertification 资质证书 https://openapi.qcc.com/dataApi/255
func (a *Api) ECICertificationSearchCertification(ctx context.Context, req *ECICertificationSearchCertificationReq) (*ECICertificationSearchCertificationResp, error) {
	var resp ECICertificationSearchCertificationResp
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
	if req.PageSize > 0 {
		c.SetQueryParam("pageSize", fmt.Sprintf("%d", req.PageSize))
	}
	if req.PageIndex > 0 {
		c.SetQueryParam("pageIndex", fmt.Sprintf("%d", req.PageIndex))
	}
	if req.IsValid != "" {
		c.SetQueryParam("isValid", req.IsValid)
	}

	reply, err := c.SetResult(&resp).Get("/ECICertification/SearchCertification")
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

type ECICertificationGetCertificationDetailByIDReq struct {
	// 查询证书的 ID
	CertID string
}

type ECICertificationGetCertificationDetailByIDResp struct {
	Response[ECICertificationGetCertificationDetailByIDRespResult]
}

type ECICertificationGetCertificationDetailByIDRespResult struct {
	ID     string                                                      `json:"Id"`
	Data   ECICertificationGetCertificationDetailByIDRespResultData    `json:"Data"`
	Schema *ECICertificationGetCertificationDetailByIDRespResultSchema `json:"Schema"`
}

type ECICertificationGetCertificationDetailByIDRespResultData struct {
	ProductNameAndUnitMain                    string `json:"产品名称及单元（主）"`
	IssuingAuthorityApprovalNo                string `json:"发证机构批准号"`
	ProductCategory                           string `json:"产品类别"`
	CertifiedOrgName                          string `json:"获证组织基本信息-企业名称"`
	ClientOrgAddress                          string `json:"认证委托人-组织地址"`
	IsMultiSiteCovered                        string `json:"是否覆盖多场所"`
	CertifiedOrgPostalCode                    string `json:"获证组织基本信息-邮政编码"`
	ClientCountryOrRegion                     string `json:"认证委托人-所在国别地区"`
	CertifiedOrgCreditCode                    string `json:"获证组织基本信息-统一社会信用代码/组织机构代码"`
	ProducerOrgName                           string `json:"生产企业-组织名称"`
	ReportDate                                string `json:"信息上报日期"`
	CertifiedOrgCountryOrRegion               string `json:"获证组织基本信息-所在国别地区"`
	ManufacturerCreditCode                    string `json:"生产者（制造商）-统一社会信用代码/组织机构代码"`
	FirstCertificationDate                    string `json:"初次获证日期"`
	CertifiedOrgCoveredPersonCount            string `json:"获证组织基本信息-本证书体系覆盖人数"`
	CertificationStandardTechnicalRequirement string `json:"认证依据的标准和技术要求"`
	IssuingAuthorityAddress                   string `json:"发证机构地址"`
	IssueDate                                 string `json:"颁证日期"`
	ChangeDate                                string `json:"变更日期"`
	ManufacturerOrgAddress                    string `json:"生产者（制造商）-组织地址"`
	CertificateNo                             string `json:"证书编号"`
	CertificationScope                        string `json:"认证范围/认证覆盖的业务范围"`
	CertificationCoveredSite                  string `json:"认证覆盖的场所名称及地址"`
	ProducerOrgAddress                        string `json:"生产企业-组织地址"`
	ManufacturerCountryOrRegion               string `json:"生产者（制造商）-所在国别地区"`
	CertificationBasis                        string `json:"认证依据"`
	RecertificationCount                      string `json:"再认证次数"`
	IssuingAuthorityPhone                     string `json:"发证机构电话"`
	IssuingAuthorityName                      string `json:"发证机构名称"`
	IssuingAuthorityValidUntil                string `json:"发证机构有效期"`
	CompanyName                               string `json:"企业名称"`
	CertificateExpireDate                     string `json:"证书到期日期"`
	IssuingAuthorityStatus                    string `json:"发证机构状态"`
	ProductNameAndUnitSecondary               string `json:"产品名称及单元（次）"`
	ManufacturerOrgName                       string `json:"生产者（制造商）-组织名称"`
	CertificateStatus                         string `json:"证书状态"`
	MainCertificateNo                         string `json:"主认证证书号"`
	IssuingAuthorityWebsite                   string `json:"发证机构网址"`
	ClientCreditCode                          string `json:"认证委托人-统一社会信用代码/组织机构代码"`
	IssuingAuthorityBusinessScope             string `json:"发证机构业务范围"`
	IsSubCertificate                          string `json:"是否是子证书"`
	SupervisionCount                          string `json:"监督次数"`
	CertificationRelatedEmployeeCount         string `json:"认证相关员工数量"`
	CertifiedOrgAddress                       string `json:"获证组织基本信息-组织地址"`
	SpecificationModel                        string `json:"规格型号"`
	ClientOrgName                             string `json:"认证委托人-组织名称"`
	ServiceCertificationField                 string `json:"服务认证所属领域"`
	CertificateAccreditationMark              string `json:"证书使用的认可标识"`
}

type ECICertificationGetCertificationDetailByIDRespResultSchema struct {
}

// ECICertificationGetCertificationDetailByID 资质证书详情 https://openapi.qcc.com/dataApi/255
func (a *Api) ECICertificationGetCertificationDetailByID(ctx context.Context, req *ECICertificationGetCertificationDetailByIDReq) (*ECICertificationGetCertificationDetailByIDResp, error) {
	var resp ECICertificationGetCertificationDetailByIDResp
	token, unix, err := a.auth()
	if err != nil {
		return nil, fmt.Errorf("auth: %w", err)
	}
	c := a.cli.R().
		SetContext(ctx).
		SetHeader("Token", token).
		SetHeader("Timespan", unix).
		SetQueryParam("key", a.cfg.Key).
		SetQueryParam("certId", req.CertID)

	reply, err := c.SetResult(&resp).Get("/ECICertification/GetCertificationDetailById")
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
