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
					"FinancialIndexList":[{"IndexName":"毛利率"}],
					"DeclarationDetail":{"CorporateInTaxDeclareList":[]},
					"CollectionDetail":{"CorporateInTaxCollectionList":[]},
					"SaleList":[],
					"TaxData":{"TotalTaxList":[]},
					"TaxBurdenRateList":[],
					"FinancialList":[],
					"SupplierCustomerList":[],
					"TopCustomerList":[],
					"TopSupplierList":[],
					"BreakLawDetailList":[],
					"BreakLawSummaryList":[],
					"ExpenseDetail":{"ElectricityExpenseList":[]},
					"CashFlowList":[]
				}
			}
		}`), &resp)

		assert.NoError(t, err)
		assert.Equal(t, "S", resp.Result.DataStatus)
		assert.Len(t, resp.Result.Data.FinancialIndexList, 1)
		assert.NotEmpty(t, resp.Result.Data.DeclarationDetail)
		assert.NotEmpty(t, resp.Result.Data.CollectionDetail)
		assert.NotEmpty(t, resp.Result.Data.TaxData)
		assert.NotEmpty(t, resp.Result.Data.ExpenseDetail)
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
				"FlowItems":[{"Name":"申请"}]
			}
		}`), &tmResp)
		assert.NoError(t, err)
		assert.Equal(t, "2024-01-01", tmResp.Result.AnnouncementDate)
		assert.Equal(t, "2024-12-31", tmResp.Result.HouQiZhiDingDate)
		assert.Equal(t, int64(38), tmResp.Result.IntCls)
		assert.Equal(t, int64(4), tmResp.Result.Status)
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
