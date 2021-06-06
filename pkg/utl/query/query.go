package query

import (
	"github.com/epavanello/gorsk/pkg/models"
	"github.com/labstack/echo"
)

// List prepares data for list queries
func List(u models.AuthUser) (*models.ListQuery, error) {
	switch true {
	case u.Role <= models.AdminRole: // user is SuperAdmin or Admin
		return nil, nil
	case u.Role == models.CompanyAdminRole:
		return &models.ListQuery{Query: "company_id = ?", ID: u.CompanyID}, nil
	case u.Role == models.LocationAdminRole:
		return &models.ListQuery{Query: "location_id = ?", ID: u.LocationID}, nil
	default:
		return nil, echo.ErrForbidden
	}
}
