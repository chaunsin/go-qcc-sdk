package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratedResponseSchemasParseReviewedFields(t *testing.T) {
	t.Run("TaxDataGetData keeps payload fields under Data", func(t *testing.T) {
		var resp TaxDataGetDataResp
		err := json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{
				"DataStatus":"S",
				"Data":{
					"FinancialIndexList":[{"IndexName":"毛利率","ValueList":[{"Date":"2024","Value":"14.00%"}]}],
					"DeclarationDetail":{"CorporateInTaxDeclareList":[{"ThisYearSaleRevenue":"49.84"}]},
					"CollectionDetail":{"CorporateInTaxCollectionList":[{"ActualAmount":"0.03"}]},
					"SaleList":[],
					"TaxData":{"TotalTaxList":[{"Year":"2024","DataList":[{"Month":"1","Amount":"0.18"}]}]},
					"TaxBurdenRateList":[],
					"FinancialList":[],
					"SupplierCustomerList":[],
					"TopCustomerList":[],
					"TopSupplierList":[],
					"BreakLawDetailList":[],
					"BreakLawSummaryList":[],
					"ExpenseDetail":{"ElectricityExpenseList":[{"Year":"2024","DataList":[{"Month":"1","Amount":"0.35"}]}]},
					"CashFlowList":[]
				}
			}
		}`), &resp)

		assert.NoError(t, err)
		assert.Equal(t, "S", resp.Result.DataStatus)
		assert.Len(t, resp.Result.Data.FinancialIndexList, 1)
		assert.Equal(t, "毛利率", resp.Result.Data.FinancialIndexList[0].IndexName)
		assert.Equal(t, "49.84", resp.Result.Data.DeclarationDetail.CorporateInTaxDeclareList[0].ThisYearSaleRevenue)
		assert.Equal(t, "0.03", resp.Result.Data.CollectionDetail.CorporateInTaxCollectionList[0].ActualAmount)
		assert.Equal(t, "0.18", resp.Result.Data.TaxData.TotalTaxList[0].DataList[0].Amount)
		assert.Equal(t, "0.35", resp.Result.Data.ExpenseDetail.ElectricityExpenseList[0].DataList[0].Amount)
	})

	t.Run("HK and certification payloads use concrete structs", func(t *testing.T) {
		var hkResp HKDataGetDataResp
		err := json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{
				"DataStatus":"S",
				"Data":{
					"basic":{"Details":{"CompanyNameEng":"ACME LIMITED","OriginalNameList":[{"Name":"旧名"}]}},
					"capitalstructure_live_simplified":{"Details":[{"Currency":"HKD","TotalAmount":"100"}]},
					"directors_live":{"Details":[{"FullNameEng":"CHAN","Type":"自然人"}]},
					"shareholders":{"Date":"2024-01-01","Details":[{"FullName":"张三","NumberofShares":"1"}],"verified":"1"},
					"company_secretaries_live":{"Details":[{"FullNameChn":"秘书","AppointedDate":"2024-01-01"}]}
				}
			}
		}`), &hkResp)
		assert.NoError(t, err)
		assert.Equal(t, "ACME LIMITED", hkResp.Result.Data.Basic.Details.CompanyNameEng)
		assert.Equal(t, "旧名", hkResp.Result.Data.Basic.Details.OriginalNameList[0].Name)
		assert.Equal(t, "HKD", hkResp.Result.Data.CapitalStructureLiveSimplified.Details[0].Currency)
		assert.Equal(t, "1", hkResp.Result.Data.Shareholders.Verified)

		var hknrtResp HKNRTDataGetDataResp
		err = json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{
				"DataStatus":"S",
				"Data":{
					"basic":{"Details":{"CompanyNameEng":"ACME LIMITED"}},
					"capitalstructure_historical":{"Date":"2024-01-01","Details":[{"ClassofShares":"Ordinary","TotalNumber":"100"}]},
					"director_historical":{"Details":[{"FullNameEng":"CHAN","PassportNumber":"P1"}]},
					"shareholders":{"Details":[{"FullName":"张三","PercentofClass":"100.00%"}]},
					"company_secretaries_historical":{"Details":[{"FullNameEng":"SECRETARY"}]},
					"verified":"1",
					"OriginalFile":"https://example.invalid/file.pdf"
				}
			}
		}`), &hknrtResp)
		assert.NoError(t, err)
		assert.Equal(t, "Ordinary", hknrtResp.Result.Data.CapitalStructureHistorical.Details[0].ClassOfShares)
		assert.Equal(t, "P1", hknrtResp.Result.Data.DirectorHistorical.Details[0].PassportNumber)
		assert.Equal(t, "https://example.invalid/file.pdf", hknrtResp.Result.Data.OriginalFile)

		var certResp ECICertificationGetCertificationDetailByIDResp
		err = json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{
				"Id":"cert-1",
				"Data":{"企业名称":"企查查科技有限公司","证书编号":"17418Q20358R1M"},
				"Schema":null
			}
		}`), &certResp)
		assert.NoError(t, err)
		assert.Equal(t, "企查查科技有限公司", certResp.Result.Data.CompanyName)
		assert.Equal(t, "17418Q20358R1M", certResp.Result.Data.CertificateNo)
		assert.Nil(t, certResp.Result.Schema)
	})

	t.Run("truncated JSON tags parse official names", func(t *testing.T) {
		var crResp CopyRightSearchSoftwareCrResp
		err := json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":[{
				"RegisterAperDate":"2023-01-10",
				"FinishDevelopDate":"2022-04-23"
			}]
		}`), &crResp)
		assert.NoError(t, err)
		if assert.Len(t, crResp.Result, 1) {
			assert.Equal(t, "2023-01-10", crResp.Result[0].RegisterAperDate)
			assert.Equal(t, "2022-04-23", crResp.Result[0].FinishDevelopDate)
		}

		var patentResp PatentV4GetDetailsResp
		err = json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{
				"ApplicationNumber":"CN202111250145.3",
				"PublicationNumber":"CN114154014A",
				"InventorStringList":["张三"],
				"AssigneestringList":["企查查"],
				"PatentLegalHistory":[{"Desc":"公开","LegalStatus":"","LegalStatusDate":"2022-03-08"}]
			}
		}`), &patentResp)
		assert.NoError(t, err)
		assert.Equal(t, "CN202111250145.3", patentResp.Result.ApplicationNumber)
		assert.Equal(t, "CN114154014A", patentResp.Result.PublicationNumber)
		assert.Equal(t, []string{"张三"}, patentResp.Result.InventorStringList)
		assert.Equal(t, []string{"企查查"}, patentResp.Result.AssigneestringList)
		if assert.Len(t, patentResp.Result.PatentLegalHistory, 1) {
			assert.Equal(t, "公开", patentResp.Result.PatentLegalHistory[0].Desc)
		}

		var intlResp InternationalPatentCheckGetDetailResp
		err = json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{
				"BasicInfo":{"ApplicationNumber":"PCT/CN2019/116526","PatenteeList":[{"KeyNo":"k1","Name":"n1"}]},
				"InstructionImgList":["https://example.invalid/patent.jpg"]
			}
		}`), &intlResp)
		assert.NoError(t, err)
		assert.Equal(t, "PCT/CN2019/116526", intlResp.Result.BasicInfo.ApplicationNumber)
		assert.Equal(t, []string{"https://example.invalid/patent.jpg"}, intlResp.Result.InstructionImgList)

		var tmResp TmGetDetailsResp
		err = json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{
				"AnnouncementDate":"2024-01-01",
				"HouQiZhiDingDate":"2024-12-31",
				"IntCls":38,
				"Status":4,
				"FlowItems":[{"FlowId":"f1","FlowItem":"申请","FlowDate":"2024-01-01"}]
			}
		}`), &tmResp)
		assert.NoError(t, err)
		assert.Equal(t, "2024-01-01", tmResp.Result.AnnouncementDate)
		assert.Equal(t, "2024-12-31", tmResp.Result.HouQiZhiDingDate)
		assert.Equal(t, int64(38), tmResp.Result.IntCls)
		assert.Equal(t, int64(4), tmResp.Result.Status)
		if assert.Len(t, tmResp.Result.FlowItems, 1) {
			assert.Equal(t, "申请", tmResp.Result.FlowItems[0].FlowItem)
		}
	})

	t.Run("nested paramList fields use concrete structs", func(t *testing.T) {
		var chattelResp ChattelMortgageCheckGetListResp
		err := json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{
				"VerifyResult":1,
				"Data":[{
					"Detail":{
						"GuaranteeList":[{"Name":"设备","Ownership":"企业","KeyNoList":[{"KeyNo":"k1","Name":"n1"}]}],
						"CancelInfo":{"CancelDate":"2024-01-01","CancelReason":"注销"},
						"ChangeList":[{"ChangeDate":"2024-01-02","ChangeContent":"变更"}]
					}
				}]
			}
		}`), &chattelResp)
		assert.NoError(t, err)
		if assert.Len(t, chattelResp.Result.Data, 1) {
			assert.Equal(t, "设备", chattelResp.Result.Data[0].Detail.GuaranteeList[0].Name)
			assert.Equal(t, "注销", chattelResp.Result.Data[0].Detail.CancelInfo.CancelReason)
			assert.Equal(t, "变更", chattelResp.Result.Data[0].Detail.ChangeList[0].ChangeContent)
		}

		var historyChattelResp HistoryChattelMortgageCheckGetDetailResp
		err = json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{
				"Pledge":{"RegistNo":"R1"},
				"PledgeeList":[{"Name":"抵押权人","IdentityType":"工商营业执照","IdentityNo":"N1","KeyNo":"k1"}],
				"SecuredClaim":{"Kind":"其他合同","Amount":"200万元"},
				"GuaranteeList":[{"Name":"设备","OwnershipList":[{"KeyNo":"k2","Name":"所有权人"}]}],
				"Pledgor":{"KeyNo":"k3","Name":"抵押人"}
			}
		}`), &historyChattelResp)
		assert.NoError(t, err)
		assert.Equal(t, "R1", historyChattelResp.Result.Pledge.RegistNo)
		assert.Equal(t, "抵押权人", historyChattelResp.Result.PledgeeList[0].Name)
		assert.Equal(t, "所有权人", historyChattelResp.Result.GuaranteeList[0].OwnershipList[0].Name)

		var tenderResp TenderCheckGetListResp
		err = json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{"VerifyResult":1,"Data":[{"WinBidUnitList":[{"WinBidAmt":"59.6万元","KeyNo":"k1"}],"AgentUnitList":[{"Name":"代理"}]}]}
		}`), &tenderResp)
		assert.NoError(t, err)
		if assert.Len(t, tenderResp.Result.Data, 1) {
			assert.Equal(t, "59.6万元", tenderResp.Result.Data[0].WinBidUnitList[0].WinBidAmt)
			assert.Equal(t, "代理", tenderResp.Result.Data[0].AgentUnitList[0].Name)
		}

		var tenderDetailResp TenderCheckGetDetailResp
		err = json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{"ContentUrl":"https://example.invalid","PublishDate":"2024-01-01","Title":"招标","Data":{"Content":"详情正文"}}
		}`), &tenderDetailResp)
		assert.NoError(t, err)
		assert.Equal(t, "详情正文", tenderDetailResp.Result.Data.Content)

		var bangDanResp BangDanCheckGetDetailResp
		err = json.Unmarshal([]byte(`{
			"Status":"200",
			"Result":{"VerifyResult":1,"Data":[{"Name":"华为","RelatedInfoList":[{"KeyNo":"k1","Name":"华为技术有限公司"}]}]}
		}`), &bangDanResp)
		assert.NoError(t, err)
		assert.Equal(t, int64(1), bangDanResp.Result.VerifyResult)
		if assert.Len(t, bangDanResp.Result.Data, 1) {
			assert.Equal(t, "k1", bangDanResp.Result.Data[0].RelatedInfoList[0].KeyNo)
		}
	})
}
