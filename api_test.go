package api

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	ctx = context.Background()
	srv = New(&Config{
		Key:       "",
		SecretKey: "",
	})
)

func TestQCCService_GetCompanyGraph(t *testing.T) {
	got, err := srv.GetCompanyGraph(ctx, "企查查科技股份有限公司")
	assert.NoError(t, err)
	t.Logf("GetCompanyGraph: %+v", got)
}

func TestQCCService_GetList(t *testing.T) {
	got, err := srv.GetList(ctx, &FuzzySearchGetListReq{SearchKey: "企查查科技股份有限公司"})
	assert.NoError(t, err)
	t.Logf("GetList: %+v", got)
}

func TestQCCService_AdminPenaltyCheckGetList(t *testing.T) {
	var req = &AdminPenaltyCheckGetListReq{
		SearchKey: "企查查科技股份有限公司",
		PageIndex: 1,
		PageSize:  10,
	}
	got, err := srv.AdminPenaltyCheckGetList(ctx, req)
	assert.NoError(t, err)
	t.Logf("AdminPenaltyCheckGetList: %+v", got)
}
