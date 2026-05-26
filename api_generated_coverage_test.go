package api

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedQCCInterfacesUseExpectedRequests(t *testing.T) {
	tests := []struct {
		name   string
		path   string
		method string
		query  map[string]string
		body   map[string]any
		result string
		call   func(*Api) error
	}{
		{
			name:   "231/TmSearchByApplicant",
			path:   "/tm/SearchByApplicant",
			method: "GET",
			query:  map[string]string{"keyword": "test-keyword", "pageSize": "7", "pageIndex": "7"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.TmSearchByApplicant(ctx, &TmSearchByApplicantReq{Keyword: "test-keyword", PageSize: 7, PageIndex: 7})
				return err
			},
		},
		{
			name:   "233/CopyRightSearchCopyRight",
			path:   "/CopyRight/SearchCopyRight",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageSize": "7", "pageIndex": "7"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.CopyRightSearchCopyRight(ctx, &CopyRightSearchCopyRightReq{SearchKey: "test-searchKey", PageSize: 7, PageIndex: 7})
				return err
			},
		},
		{
			name:   "255/ECICertificationSearchCertification",
			path:   "/ECICertification/SearchCertification",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageSize": "7", "pageIndex": "7", "isValid": "test-isValid"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.ECICertificationSearchCertification(ctx, &ECICertificationSearchCertificationReq{SearchKey: "test-searchKey", PageSize: 7, PageIndex: 7, IsValid: "test-isValid"})
				return err
			},
		},
		{
			name:   "260/ReportCreateReport",
			path:   "/Report/CreateReport",
			method: "POST",
			query:  map[string]string{"keyNo": "test-keyNo", "reportFormat": "test-reportFormat"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ReportCreateReport(ctx, &ReportCreateReportReq{KeyNo: "test-keyNo", ReportFormat: "test-reportFormat"})
				return err
			},
		},
		{
			name:   "514/PatentV4Search",
			path:   "/PatentV4/Search",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "ipc": "test-ipc", "pubDateBegin": "test-pubDateBegin", "pubDateEnd": "test-pubDateEnd", "pageSize": "test-pageSize", "pageIndex": "test-pageIndex"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.PatentV4Search(ctx, &PatentV4SearchReq{SearchKey: "test-searchKey", Ipc: "test-ipc", PubDateBegin: "test-pubDateBegin", PubDateEnd: "test-pubDateEnd", PageSize: "test-pageSize", PageIndex: "test-pageIndex"})
				return err
			},
		},
		{
			name:   "515/WebSiteV4GetCompanyWebSite",
			path:   "/WebSiteV4/GetCompanyWebSite",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.WebSiteV4GetCompanyWebSite(ctx, &WebSiteV4GetCompanyWebSiteReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "521/CompanyProductV4SearchCompanyCompanyProducts",
			path:   "/CompanyProductV4/SearchCompanyCompanyProducts",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageSize": "test-pageSize", "pageIndex": "test-pageIndex"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.CompanyProductV4SearchCompanyCompanyProducts(ctx, &CompanyProductV4SearchCompanyCompanyProductsReq{SearchKey: "test-searchKey", PageSize: "test-pageSize", PageIndex: "test-pageIndex"})
				return err
			},
		},
		{
			name:   "691/TaxCreditGetList",
			path:   "/TaxCredit/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.TaxCreditGetList(ctx, &TaxCreditGetListReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "701/CompanyNewsSearchNews",
			path:   "/CompanyNews/SearchNews",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "emotionType": "test-emotionType", "startDate": "test-startDate", "endDate": "test-endDate", "pageSize": "7", "pageIndex": "7"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.CompanyNewsSearchNews(ctx, &CompanyNewsSearchNewsReq{SearchKey: "test-searchKey", EmotionType: "test-emotionType", StartDate: "test-startDate", EndDate: "test-endDate", PageSize: 7, PageIndex: 7})
				return err
			},
		},
		{
			name:   "718/RecruitmentGetList",
			path:   "/Recruitment/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.RecruitmentGetList(ctx, &RecruitmentGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "723/CustomerGetList",
			path:   "/Customer/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "year": "test-year", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CustomerGetList(ctx, &CustomerGetListReq{SearchKey: "test-searchKey", Year: "test-year", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "724/SupplierGetList",
			path:   "/Supplier/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "year": "test-year", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.SupplierGetList(ctx, &SupplierGetListReq{SearchKey: "test-searchKey", Year: "test-year", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "726/TelecomLicenseGetList",
			path:   "/TelecomLicense/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.TelecomLicenseGetList(ctx, &TelecomLicenseGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "727/CreditRatingGetList",
			path:   "/CreditRating/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.CreditRatingGetList(ctx, &CreditRatingGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "739/ExceptionCheckGetList",
			path:   "/ExceptionCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ExceptionCheckGetList(ctx, &ExceptionCheckGetListReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "745/LandMortgageCheckGetList",
			path:   "/LandMortgageCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.LandMortgageCheckGetList(ctx, &LandMortgageCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "746/EnvPunishmentCheckGetList",
			path:   "/EnvPunishmentCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.EnvPunishmentCheckGetList(ctx, &EnvPunishmentCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "747/ChattelMortgageCheckGetList",
			path:   "/ChattelMortgageCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ChattelMortgageCheckGetList(ctx, &ChattelMortgageCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "748/SeriousIllegalCheckGetList",
			path:   "/SeriousIllegalCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.SeriousIllegalCheckGetList(ctx, &SeriousIllegalCheckGetListReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "749/SimpleCancelCheckGetInfo",
			path:   "/SimpleCancelCheck/GetInfo",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.SimpleCancelCheckGetInfo(ctx, &SimpleCancelCheckGetInfoReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "750/PublishNoticeCheckGetList",
			path:   "/PublishNoticeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PublishNoticeCheckGetList(ctx, &PublishNoticeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "751/EquityPledgedCheckGetList",
			path:   "/EquityPledgedCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.EquityPledgedCheckGetList(ctx, &EquityPledgedCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "753/StockRightPledgeCheckGetList",
			path:   "/StockRightPledgeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.StockRightPledgeCheckGetList(ctx, &StockRightPledgeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "756/TaxIllegalCheckGetList",
			path:   "/TaxIllegalCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.TaxIllegalCheckGetList(ctx, &TaxIllegalCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "757/TaxOweNoticeCheckGetList",
			path:   "/TaxOweNoticeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.TaxOweNoticeCheckGetList(ctx, &TaxOweNoticeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "758/EndExecuteCaseCheckGetList",
			path:   "/EndExecuteCaseCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.EndExecuteCaseCheckGetList(ctx, &EndExecuteCaseCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "759/DeliveryNoticeCheckGetList",
			path:   "/DeliveryNoticeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.DeliveryNoticeCheckGetList(ctx, &DeliveryNoticeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "760/InquiryAssessCheckGetList",
			path:   "/InquiryAssessCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.InquiryAssessCheckGetList(ctx, &InquiryAssessCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "762/OffFilingCheckGetList",
			path:   "/OffFilingCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.OffFilingCheckGetList(ctx, &OffFilingCheckGetListReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "764/PersonSXCheckGetList",
			path:   "/PersonSXCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonSXCheckGetList(ctx, &PersonSXCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "765/PersonZXCheckGetList",
			path:   "/PersonZXCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonZXCheckGetList(ctx, &PersonZXCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "766/PersonSumptuaryCheckGetList",
			path:   "/PersonSumptuaryCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonSumptuaryCheckGetList(ctx, &PersonSumptuaryCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "767/PersonSEFreezeCheckGetList",
			path:   "/PersonSEFreezeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonSEFreezeCheckGetList(ctx, &PersonSEFreezeCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "768/PersonSEPledgedCheckGetList",
			path:   "/PersonSEPledgedCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonSEPledgedCheckGetList(ctx, &PersonSEPledgedCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "769/PersonJudgeDocCheckGetList",
			path:   "/PersonJudgeDocCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonJudgeDocCheckGetList(ctx, &PersonJudgeDocCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "770/PersonSRPledgeCheckGetList",
			path:   "/PersonSRPledgeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonSRPledgeCheckGetList(ctx, &PersonSRPledgeCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "771/PersonEndExCaseCheckGetList",
			path:   "/PersonEndExCaseCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonEndExCaseCheckGetList(ctx, &PersonEndExCaseCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "772/PersonIllegalCheckGetList",
			path:   "/PersonIllegalCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonIllegalCheckGetList(ctx, &PersonIllegalCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "773/PersonCaseFilingCheckGetList",
			path:   "/PersonCaseFilingCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonCaseFilingCheckGetList(ctx, &PersonCaseFilingCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "774/PersonCourtNoticeCheckGetList",
			path:   "/PersonCourtNoticeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonCourtNoticeCheckGetList(ctx, &PersonCourtNoticeCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "775/PersonCSACheckGetList",
			path:   "/PersonCSACheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonCSACheckGetList(ctx, &PersonCSACheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "776/PersonDeliverNoticeCheckGetList",
			path:   "/PersonDeliverNoticeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonDeliverNoticeCheckGetList(ctx, &PersonDeliverNoticeCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "777/PersonInquiryAssessCheckGetList",
			path:   "/PersonInquiryAssessCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonInquiryAssessCheckGetList(ctx, &PersonInquiryAssessCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "806/PersonHisSXCheckGetList",
			path:   "/PersonHisSXCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonHisSXCheckGetList(ctx, &PersonHisSXCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "807/PersonHisZXCheckGetList",
			path:   "/PersonHisZXCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonHisZXCheckGetList(ctx, &PersonHisZXCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "808/PersonHisSTCheckGetList",
			path:   "/PersonHisSTCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonHisSTCheckGetList(ctx, &PersonHisSTCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "880/BelongGroupGetInfo",
			path:   "/BelongGroup/GetInfo",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.BelongGroupGetInfo(ctx, &BelongGroupGetInfoReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "883/GroupMemberGetList",
			path:   "/GroupMember/GetList",
			method: "GET",
			query:  map[string]string{"groupId": "test-groupId", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.GroupMemberGetList(ctx, &GroupMemberGetListReq{GroupID: "test-groupId", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "891/OperationCheckGetList",
			path:   "/OperationCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.OperationCheckGetList(ctx, &OperationCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "892/AdminLicenseCheckGetList",
			path:   "/AdminLicenseCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.AdminLicenseCheckGetList(ctx, &AdminLicenseCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "896/CreditorRightsCheckGetList",
			path:   "/CreditorRightsCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CreditorRightsCheckGetList(ctx, &CreditorRightsCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "921/HistoryEciCheckGetInfo",
			path:   "/HistoryEciCheck/GetInfo",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryEciCheckGetInfo(ctx, &HistoryEciCheckGetInfoReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "922/HistoryOperCheckGetList",
			path:   "/HistoryOperCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryOperCheckGetList(ctx, &HistoryOperCheckGetListReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "923/HistoryEmployeeCheckGetList",
			path:   "/HistoryEmployeeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryEmployeeCheckGetList(ctx, &HistoryEmployeeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "924/HistoryInvestmentCheckGetList",
			path:   "/HistoryInvestmentCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryInvestmentCheckGetList(ctx, &HistoryInvestmentCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "925/HistoryPartnerCheckGetList",
			path:   "/HistoryPartnerCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryPartnerCheckGetList(ctx, &HistoryPartnerCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "926/HistoryShiXinCheckGetList",
			path:   "/HistoryShiXinCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryShiXinCheckGetList(ctx, &HistoryShiXinCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "927/HistoryZhiXingCheckGetList",
			path:   "/HistoryZhiXingCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryZhiXingCheckGetList(ctx, &HistoryZhiXingCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "928/HistoryEndExCaseCheckGetList",
			path:   "/HistoryEndExCaseCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryEndExCaseCheckGetList(ctx, &HistoryEndExCaseCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "929/HistorySumptuaryCheckGetList",
			path:   "/HistorySumptuaryCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistorySumptuaryCheckGetList(ctx, &HistorySumptuaryCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "930/HistoryCNoticeCheckGetList",
			path:   "/HistoryCNoticeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryCNoticeCheckGetList(ctx, &HistoryCNoticeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "931/HistoryDNoticeCheckGetList",
			path:   "/HistoryDNoticeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryDNoticeCheckGetList(ctx, &HistoryDNoticeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "932/HistoryJudgeDocCheckGetList",
			path:   "/HistoryJudgeDocCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryJudgeDocCheckGetList(ctx, &HistoryJudgeDocCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "933/HistoryAdminPenaltyCheckGetList",
			path:   "/HistoryAdminPenaltyCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryAdminPenaltyCheckGetList(ctx, &HistoryAdminPenaltyCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "934/HistoryChattelMortgageCheckGetList",
			path:   "/HistoryChattelMortgageCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryChattelMortgageCheckGetList(ctx, &HistoryChattelMortgageCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "935/HistoryCourtAnnoCheckGetList",
			path:   "/HistoryCourtAnnoCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryCourtAnnoCheckGetList(ctx, &HistoryCourtAnnoCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "936/HistoryEquityPledgedCheckGetList",
			path:   "/HistoryEquityPledgedCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryEquityPledgedCheckGetList(ctx, &HistoryEquityPledgedCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "937/HistoryAdminLicenseCheckGetList",
			path:   "/HistoryAdminLicenseCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryAdminLicenseCheckGetList(ctx, &HistoryAdminLicenseCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "939/HistoryEquityFreezeCheckGetList",
			path:   "/HistoryEquityFreezeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryEquityFreezeCheckGetList(ctx, &HistoryEquityFreezeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "940/HistoryBankruptcyCheckGetList",
			path:   "/HistoryBankruptcyCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryBankruptcyCheckGetList(ctx, &HistoryBankruptcyCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "941/HistoryLandMortgageCheckGetList",
			path:   "/HistoryLandMortgageCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryLandMortgageCheckGetList(ctx, &HistoryLandMortgageCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "944/PersonHisOperCompanyCheckGetList",
			path:   "/PersonHisOperCompanyCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonHisOperCompanyCheckGetList(ctx, &PersonHisOperCompanyCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "945/PersonHisInvestCompanyCheckGetList",
			path:   "/PersonHisInvestCompanyCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonHisInvestCompanyCheckGetList(ctx, &PersonHisInvestCompanyCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "946/PersonHisJobCompanyCheckGetList",
			path:   "/PersonHisJobCompanyCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonHisJobCompanyCheckGetList(ctx, &PersonHisJobCompanyCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "948/ShellScanGetList",
			path:   "/ShellScan/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ShellScanGetList(ctx, &ShellScanGetListReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "949/AcctScanGetInfo",
			path:   "/AcctScan/GetInfo",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.AcctScanGetInfo(ctx, &AcctScanGetInfoReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "950/CompanyFinancingSearchGetList",
			path:   "/CompanyFinancingSearch/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CompanyFinancingSearchGetList(ctx, &CompanyFinancingSearchGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "951/HistoryExceptionCheckGetList",
			path:   "/HistoryExceptionCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryExceptionCheckGetList(ctx, &HistoryExceptionCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "953/DoubleRandomCheckGetList",
			path:   "/DoubleRandomCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.DoubleRandomCheckGetList(ctx, &DoubleRandomCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "958/TenderCheckGetList",
			path:   "/TenderCheck/GetList",
			method: "GET",
			query:  map[string]string{"keyword": "test-keyword", "areaCode": "test-areaCode", "msgType": "test-msgType", "pubDateStart": "test-pubDateStart", "pubDateEnd": "test-pubDateEnd", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.TenderCheckGetList(ctx, &TenderCheckGetListReq{Keyword: "test-keyword", AreaCode: "test-areaCode", MsgType: "test-msgType", PubDateStart: "test-pubDateStart", PubDateEnd: "test-pubDateEnd", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "959/BangDanCheckGetList",
			path:   "/BangDanCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.BangDanCheckGetList(ctx, &BangDanCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "961/BusinessRealTraceCheckGetInfo",
			path:   "/BusinessRealTraceCheck/GetInfo",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.BusinessRealTraceCheckGetInfo(ctx, &BusinessRealTraceCheckGetInfoReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "979/SpotCheckGetList",
			path:   "/SpotCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.SpotCheckGetList(ctx, &SpotCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "980/LiquidationCheckGetDetail",
			path:   "/LiquidationCheck/GetDetail",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.LiquidationCheckGetDetail(ctx, &LiquidationCheckGetDetailReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "987/PersonSelfLimitExitCheckGetList",
			path:   "/PersonSelfLimitExitCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "personName": "test-personName", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PersonSelfLimitExitCheckGetList(ctx, &PersonSelfLimitExitCheckGetListReq{SearchKey: "test-searchKey", PersonName: "test-personName", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "988/LimitExitCheckGetList",
			path:   "/LimitExitCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.LimitExitCheckGetList(ctx, &LimitExitCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "989/LandMergeCheckGetList",
			path:   "/LandMergeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.LandMergeCheckGetList(ctx, &LandMergeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "990/HistoryEnvPunishmentCheckGetList",
			path:   "/HistoryEnvPunishmentCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryEnvPunishmentCheckGetList(ctx, &HistoryEnvPunishmentCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "991/HKDataCreateOrder",
			path:   "/HKData/CreateOrder",
			method: "GET",
			query:  map[string]string{"hkEntityName": "test-hkEntityName"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HKDataCreateOrder(ctx, &HKDataCreateOrderReq{HkEntityName: "test-hkEntityName"})
				return err
			},
		},
		{
			name:   "997/InternationalPatentCheckGetList",
			path:   "/InternationalPatentCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.InternationalPatentCheckGetList(ctx, &InternationalPatentCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "998/HKNRTDataCreateOrder",
			path:   "/HKNRTData/CreateOrder",
			method: "GET",
			query:  map[string]string{"hkEntityName": "test-hkEntityName"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HKNRTDataCreateOrder(ctx, &HKNRTDataCreateOrderReq{HkEntityName: "test-hkEntityName"})
				return err
			},
		},
		{
			name:   "1001/HonorQualificationGetList",
			path:   "/HonorQualification/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageSize": "test-pageSize", "pageIndex": "test-pageIndex"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HonorQualificationGetList(ctx, &HonorQualificationGetListReq{SearchKey: "test-searchKey", PageSize: "test-pageSize", PageIndex: "test-pageIndex"})
				return err
			},
		},
		{
			name:   "1022/HistoryTaxOweNoticeCheckGetList",
			path:   "/HistoryTaxOweNoticeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryTaxOweNoticeCheckGetList(ctx, &HistoryTaxOweNoticeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1023/HistorySeriousIllegalCheckGetList",
			path:   "/HistorySeriousIllegalCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistorySeriousIllegalCheckGetList(ctx, &HistorySeriousIllegalCheckGetListReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "1024/GuarantorCheckGetList",
			path:   "/GuarantorCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.GuarantorCheckGetList(ctx, &GuarantorCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1025/CreditorBreachCheckGetList",
			path:   "/CreditorBreachCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "searchType": "test-searchType", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CreditorBreachCheckGetList(ctx, &CreditorBreachCheckGetListReq{SearchKey: "test-searchKey", SearchType: "test-searchType", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1029/IPRPledgeCheckGetList",
			path:   "/IPRPledgeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.IPRPledgeCheckGetList(ctx, &IPRPledgeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1034/StandardCheckGetList",
			path:   "/StandardCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.StandardCheckGetList(ctx, &StandardCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1051/ZhixingDetailGetDetail",
			path:   "/ZhixingDetail/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ZhixingDetailGetDetail(ctx, &ZhixingDetailGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "1052/ShixinDetailGetDetail",
			path:   "/ShixinDetail/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ShixinDetailGetDetail(ctx, &ShixinDetailGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "1053/SumptuaryDetailGetDetail",
			path:   "/SumptuaryDetail/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.SumptuaryDetailGetDetail(ctx, &SumptuaryDetailGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "1055/JudgmentDocDetailGetDetail",
			path:   "/JudgmentDocDetail/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.JudgmentDocDetailGetDetail(ctx, &JudgmentDocDetailGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "1056/CourtNoticeDetailGetDetail",
			path:   "/CourtNoticeDetail/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CourtNoticeDetailGetDetail(ctx, &CourtNoticeDetailGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "1057/CourtAnnoDetailGetDetail",
			path:   "/CourtAnnoDetail/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CourtAnnoDetailGetDetail(ctx, &CourtAnnoDetailGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "1061/CaseFilingDetailGetDetail",
			path:   "/CaseFilingDetail/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CaseFilingDetailGetDetail(ctx, &CaseFilingDetailGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "1103/CompanyMonthSaleGetInfo",
			path:   "/CompanyMonthSale/GetInfo",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "dataMonth": "test-dataMonth"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CompanyMonthSaleGetInfo(ctx, &CompanyMonthSaleGetInfoReq{SearchKey: "test-searchKey", DataMonth: "test-dataMonth"})
				return err
			},
		},
		{
			name:   "1104/CompanySumSaleGetInfo",
			path:   "/CompanySumSale/GetInfo",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "dataMonth": "test-dataMonth"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CompanySumSaleGetInfo(ctx, &CompanySumSaleGetInfoReq{SearchKey: "test-searchKey", DataMonth: "test-dataMonth"})
				return err
			},
		},
		{
			name:   "1105/CompanyHistorySaleGetInfo",
			path:   "/CompanyHistorySale/GetInfo",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "dataMonth": "test-dataMonth"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CompanyHistorySaleGetInfo(ctx, &CompanyHistorySaleGetInfoReq{SearchKey: "test-searchKey", DataMonth: "test-dataMonth"})
				return err
			},
		},
		{
			name:   "1106/ShopInfoGetInfo",
			path:   "/ShopInfo/GetInfo",
			method: "GET",
			query:  map[string]string{"platformId": "test-platformId", "shopId": "test-shopId", "dataMonth": "test-dataMonth"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ShopInfoGetInfo(ctx, &ShopInfoGetInfoReq{PlatformID: "test-platformId", ShopID: "test-shopId", DataMonth: "test-dataMonth"})
				return err
			},
		},
		{
			name:   "1107/ShopMonthSaleGetInfo",
			path:   "/ShopMonthSale/GetInfo",
			method: "GET",
			query:  map[string]string{"platformId": "test-platformId", "shopId": "test-shopId", "dataMonth": "test-dataMonth"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ShopMonthSaleGetInfo(ctx, &ShopMonthSaleGetInfoReq{PlatformID: "test-platformId", ShopID: "test-shopId", DataMonth: "test-dataMonth"})
				return err
			},
		},
		{
			name:   "1108/ShopSumSaleGetInfo",
			path:   "/ShopSumSale/GetInfo",
			method: "GET",
			query:  map[string]string{"platformId": "test-platformId", "shopId": "test-shopId", "dataMonth": "test-dataMonth"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ShopSumSaleGetInfo(ctx, &ShopSumSaleGetInfoReq{PlatformID: "test-platformId", ShopID: "test-shopId", DataMonth: "test-dataMonth"})
				return err
			},
		},
		{
			name:   "1109/ShopHistorySaleGetInfo",
			path:   "/ShopHistorySale/GetInfo",
			method: "GET",
			query:  map[string]string{"platformId": "test-platformId", "shopId": "test-shopId", "dataMonth": "test-dataMonth"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ShopHistorySaleGetInfo(ctx, &ShopHistorySaleGetInfoReq{PlatformID: "test-platformId", ShopID: "test-shopId", DataMonth: "test-dataMonth"})
				return err
			},
		},
		{
			name:   "1115/TechScoreGetInfo",
			path:   "/TechScore/GetInfo",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.TechScoreGetInfo(ctx, &TechScoreGetInfoReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "1116/CreditScoreGetInfo",
			path:   "/CreditScore/GetInfo",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CreditScoreGetInfo(ctx, &CreditScoreGetInfoReq{SearchKey: "test-searchKey"})
				return err
			},
		},
		{
			name:   "1120/PrelitigationMediationCheckGetList",
			path:   "/PrelitigationMediationCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PrelitigationMediationCheckGetList(ctx, &PrelitigationMediationCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1121/HistoryPrelitigationMediationCheckGetList",
			path:   "/HistoryPrelitigationMediationCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryPrelitigationMediationCheckGetList(ctx, &HistoryPrelitigationMediationCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1124/TaxDataCreateOrder",
			path:   "/TaxData/CreateOrder",
			method: "POST",
			query:  map[string]string{},
			body:   map[string]any{"searchKey": "test-searchKey", "userName": "test-userName", "password": "test-password"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.TaxDataCreateOrder(ctx, &TaxDataCreateOrderReq{SearchKey: "test-searchKey", UserName: "test-userName", Password: "test-password"})
				return err
			},
		},
		{
			name:   "1127/DisciplinaryCheckGetList",
			path:   "/DisciplinaryCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.DisciplinaryCheckGetList(ctx, &DisciplinaryCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1128/EnvCreditEvaluationGetList",
			path:   "/EnvCreditEvaluation/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.EnvCreditEvaluationGetList(ctx, &EnvCreditEvaluationGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1130/HKAnnouncementListGetList",
			path:   "/HKAnnouncementList/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageSize": "test-pageSize", "pageIndex": "test-pageIndex"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.HKAnnouncementListGetList(ctx, &HKAnnouncementListGetListReq{SearchKey: "test-searchKey", PageSize: "test-pageSize", PageIndex: "test-pageIndex"})
				return err
			},
		},
		{
			name:   "1131/HKDocCreateOrder",
			path:   "/HKDoc/CreateOrder",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "docNumber": "test-docNumber"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HKDocCreateOrder(ctx, &HKDocCreateOrderReq{SearchKey: "test-searchKey", DocNumber: "test-docNumber"})
				return err
			},
		},
		{
			name:   "1132/TaxAbnormalCheckGetList",
			path:   "/TaxAbnormalCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.TaxAbnormalCheckGetList(ctx, &TaxAbnormalCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1133/LaborCourtAnnoCheckGetList",
			path:   "/LaborCourtAnnoCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.LaborCourtAnnoCheckGetList(ctx, &LaborCourtAnnoCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1134/LaborDeliveryNoticeCheckGetList",
			path:   "/LaborDeliveryNoticeCheck/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.LaborDeliveryNoticeCheckGetList(ctx, &LaborDeliveryNoticeCheckGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "1137/CompanyAnnouncementGetList",
			path:   "/CompanyAnnouncement/GetList",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CompanyAnnouncementGetList(ctx, &CompanyAnnouncementGetListReq{SearchKey: "test-searchKey", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "231/TmGetDetails",
			path:   "/tm/GetDetails",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call:   func(api *Api) error { _, err := api.TmGetDetails(ctx, &TmGetDetailsReq{ID: "test-id"}); return err },
		},
		{
			name:   "233/CopyRightGetCopyRight",
			path:   "/CopyRight/GetCopyRight",
			method: "GET",
			query:  map[string]string{"personName": "test-personName", "productName": "test-productName", "registeNo": "test-registeNo", "pageSize": "7", "pageIndex": "7"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.CopyRightGetCopyRight(ctx, &CopyRightGetCopyRightReq{PersonName: "test-personName", ProductName: "test-productName", RegisteNo: "test-registeNo", PageSize: 7, PageIndex: 7})
				return err
			},
		},
		{
			name:   "233/CopyRightSearchSoftwareCr",
			path:   "/CopyRight/SearchSoftwareCr",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageSize": "7", "pageIndex": "7"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.CopyRightSearchSoftwareCr(ctx, &CopyRightSearchSoftwareCrReq{SearchKey: "test-searchKey", PageSize: 7, PageIndex: 7})
				return err
			},
		},
		{
			name:   "233/CopyRightGetSoftwareCr",
			path:   "/CopyRight/GetSoftwareCr",
			method: "GET",
			query:  map[string]string{"personName": "test-personName", "fullName": "test-fullName", "shortName": "test-shortName", "registeNo": "test-registeNo", "pageSize": "7", "pageIndex": "7"},
			result: `[]`,
			call: func(api *Api) error {
				_, err := api.CopyRightGetSoftwareCr(ctx, &CopyRightGetSoftwareCrReq{PersonName: "test-personName", FullName: "test-fullName", ShortName: "test-shortName", RegisteNo: "test-registeNo", PageSize: 7, PageIndex: 7})
				return err
			},
		},
		{
			name:   "255/ECICertificationGetCertificationDetailByID",
			path:   "/ECICertification/GetCertificationDetailById",
			method: "GET",
			query:  map[string]string{"certId": "test-certId"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ECICertificationGetCertificationDetailByID(ctx, &ECICertificationGetCertificationDetailByIDReq{CertID: "test-certId"})
				return err
			},
		},
		{
			name:   "260/ReportGetReportInfo",
			path:   "/Report/GetReportInfo",
			method: "GET",
			query:  map[string]string{"orderNo": "test-orderNo"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.ReportGetReportInfo(ctx, &ReportGetReportInfoReq{OrderNo: "test-orderNo"})
				return err
			},
		},
		{
			name:   "514/PatentV4GetDetails",
			path:   "/PatentV4/GetDetails",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PatentV4GetDetails(ctx, &PatentV4GetDetailsReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "514/PatentV4SearchMultiPatents",
			path:   "/PatentV4/SearchMultiPatents",
			method: "GET",
			query:  map[string]string{"searchKey": "test-searchKey", "pageSize": "7", "pageIndex": "7", "kindcode": "test-kindcode", "ipc": "test-ipc", "pubDateBegin": "test-pubDateBegin", "pubDateEnd": "test-pubDateEnd", "appDateBegin": "test-appDateBegin", "appDateEnd": "test-appDateEnd"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.PatentV4SearchMultiPatents(ctx, &PatentV4SearchMultiPatentsReq{SearchKey: "test-searchKey", PageSize: 7, PageIndex: 7, Kindcode: "test-kindcode", Ipc: "test-ipc", PubDateBegin: "test-pubDateBegin", PubDateEnd: "test-pubDateEnd", AppDateBegin: "test-appDateBegin", AppDateEnd: "test-appDateEnd"})
				return err
			},
		},
		{
			name:   "701/CompanyNewsGetNewsDetail",
			path:   "/CompanyNews/GetNewsDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.CompanyNewsGetNewsDetail(ctx, &CompanyNewsGetNewsDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "718/RecruitmentGetDetail",
			path:   "/Recruitment/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.RecruitmentGetDetail(ctx, &RecruitmentGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "726/TelecomLicenseGetDetail",
			path:   "/TelecomLicense/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.TelecomLicenseGetDetail(ctx, &TelecomLicenseGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "892/AdminLicenseCheckGetDetail",
			path:   "/AdminLicenseCheck/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.AdminLicenseCheckGetDetail(ctx, &AdminLicenseCheckGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "934/HistoryChattelMortgageCheckGetDetail",
			path:   "/HistoryChattelMortgageCheck/GetDetail",
			method: "GET",
			query:  map[string]string{"keyNo": "test-keyNo", "id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HistoryChattelMortgageCheckGetDetail(ctx, &HistoryChattelMortgageCheckGetDetailReq{KeyNo: "test-keyNo", ID: "test-id"})
				return err
			},
		},
		{
			name:   "958/TenderCheckGetDetail",
			path:   "/TenderCheck/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.TenderCheckGetDetail(ctx, &TenderCheckGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "959/BangDanCheckGetDetail",
			path:   "/BangDanCheck/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id", "pageIndex": "test-pageIndex", "pageSize": "test-pageSize"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.BangDanCheckGetDetail(ctx, &BangDanCheckGetDetailReq{ID: "test-id", PageIndex: "test-pageIndex", PageSize: "test-pageSize"})
				return err
			},
		},
		{
			name:   "989/LandMergeCheckGetPurchaseDetail",
			path:   "/LandMergeCheck/GetPurchaseDetail",
			method: "GET",
			query:  map[string]string{"landPurId": "test-landPurId"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.LandMergeCheckGetPurchaseDetail(ctx, &LandMergeCheckGetPurchaseDetailReq{LandPurID: "test-landPurId"})
				return err
			},
		},
		{
			name:   "989/LandMergeCheckGetPublishDetail",
			path:   "/LandMergeCheck/GetPublishDetail",
			method: "GET",
			query:  map[string]string{"landPubId": "test-landPubId"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.LandMergeCheckGetPublishDetail(ctx, &LandMergeCheckGetPublishDetailReq{LandPubID: "test-landPubId"})
				return err
			},
		},
		{
			name:   "991/HKDataGetData",
			path:   "/HKData/GetData",
			method: "GET",
			query:  map[string]string{"orderNo": "test-orderNo"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HKDataGetData(ctx, &HKDataGetDataReq{OrderNo: "test-orderNo"})
				return err
			},
		},
		{
			name:   "997/InternationalPatentCheckGetDetail",
			path:   "/InternationalPatentCheck/GetDetail",
			method: "GET",
			query:  map[string]string{"id": "test-id"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.InternationalPatentCheckGetDetail(ctx, &InternationalPatentCheckGetDetailReq{ID: "test-id"})
				return err
			},
		},
		{
			name:   "998/HKNRTDataGetData",
			path:   "/HKNRTData/GetData",
			method: "GET",
			query:  map[string]string{"orderNo": "test-orderNo"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HKNRTDataGetData(ctx, &HKNRTDataGetDataReq{OrderNo: "test-orderNo"})
				return err
			},
		},
		{
			name:   "1124/TaxDataSendCode",
			path:   "/TaxData/SendCode",
			method: "GET",
			query:  map[string]string{"orderNo": "test-orderNo", "verifyCode": "test-verifyCode"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.TaxDataSendCode(ctx, &TaxDataSendCodeReq{OrderNo: "test-orderNo", VerifyCode: "test-verifyCode"})
				return err
			},
		},
		{
			name:   "1124/TaxDataGetData",
			path:   "/TaxData/GetData",
			method: "GET",
			query:  map[string]string{"orderNo": "test-orderNo"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.TaxDataGetData(ctx, &TaxDataGetDataReq{OrderNo: "test-orderNo"})
				return err
			},
		},
		{
			name:   "1131/HKDocGetData",
			path:   "/HKDoc/GetData",
			method: "GET",
			query:  map[string]string{"orderNo": "test-orderNo"},
			result: `{}`,
			call: func(api *Api) error {
				_, err := api.HKDocGetData(ctx, &HKDocGetDataReq{OrderNo: "test-orderNo"})
				return err
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api, closeServer := newTestAPI(t, func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, tt.path, r.URL.Path)
				assert.Equal(t, tt.method, r.Method)
				assert.NotEmpty(t, r.Header.Get("Token"))
				assert.NotEmpty(t, r.Header.Get("Timespan"))
				query := r.URL.Query()
				assert.Equal(t, []string{"test-key"}, query["key"])
				actualQuery := make(map[string]string, len(query))
				for key, values := range query {
					if key == "key" {
						continue
					}
					if assert.Len(t, values, 1, key) {
						actualQuery[key] = values[0]
					}
				}
				assert.Equal(t, tt.query, actualQuery)
				if tt.body != nil {
					raw, err := io.ReadAll(r.Body)
					assert.NoError(t, err)
					var body map[string]any
					assert.NoError(t, json.Unmarshal(raw, &body))
					assert.Equal(t, tt.body, body)
				}
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"Status":"200","Message":"OK","Result":` + tt.result + `}`))
			})
			defer closeServer()
			assert.NoError(t, tt.call(api))
		})
	}
}
