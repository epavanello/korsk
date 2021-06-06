package query_test

import (
	"github.com/epavanello/gorsk/pkg/models"
	"testing"

	"github.com/labstack/echo"

	"github.com/stretchr/testify/assert"

	"github.com/epavanello/gorsk/pkg/utl/query"
)

func TestList(t *testing.T) {
	type args struct {
		user models.AuthUser
	}
	cases := []struct {
		name     string
		args     args
		wantData *models.ListQuery
		wantErr  error
	}{
		{
			name: "Super admin user",
			args: args{user: models.AuthUser{
				Role: models.SuperAdminRole,
			}},
		},
		{
			name: "Company admin user",
			args: args{user: models.AuthUser{
				Role:      models.CompanyAdminRole,
				CompanyID: 1,
			}},
			wantData: &models.ListQuery{
				Query: "company_id = ?",
				ID:    1},
		},
		{
			name: "Location admin user",
			args: args{user: models.AuthUser{
				Role:       models.LocationAdminRole,
				CompanyID:  1,
				LocationID: 2,
			}},
			wantData: &models.ListQuery{
				Query: "location_id = ?",
				ID:    2},
		},
		{
			name: "Normal user",
			args: args{user: models.AuthUser{
				Role: models.UserRole,
			}},
			wantErr: echo.ErrForbidden,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			q, err := query.List(tt.args.user)
			assert.Equal(t, tt.wantData, q)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
