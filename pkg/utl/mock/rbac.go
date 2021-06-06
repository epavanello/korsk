package mock

import (
	"github.com/epavanello/gorsk/pkg/models"
	"github.com/labstack/echo"
)

// RBAC Mock
type RBAC struct {
	UserFn            func(echo.Context) models.AuthUser
	EnforceRoleFn     func(echo.Context, models.AccessRole) error
	EnforceUserFn     func(echo.Context, int) error
	EnforceCompanyFn  func(echo.Context, int) error
	EnforceLocationFn func(echo.Context, int) error
	AccountCreateFn   func(echo.Context, models.AccessRole, int, int) error
	IsLowerRoleFn     func(echo.Context, models.AccessRole) error
}

// User mock
func (a RBAC) User(c echo.Context) models.AuthUser {
	return a.UserFn(c)
}

// EnforceRole mock
func (a RBAC) EnforceRole(c echo.Context, role models.AccessRole) error {
	return a.EnforceRoleFn(c, role)
}

// EnforceUser mock
func (a RBAC) EnforceUser(c echo.Context, id int) error {
	return a.EnforceUserFn(c, id)
}

// EnforceCompany mock
func (a RBAC) EnforceCompany(c echo.Context, id int) error {
	return a.EnforceCompanyFn(c, id)
}

// EnforceLocation mock
func (a RBAC) EnforceLocation(c echo.Context, id int) error {
	return a.EnforceLocationFn(c, id)
}

// AccountCreate mock
func (a RBAC) AccountCreate(c echo.Context, roleID models.AccessRole, companyID, locationID int) error {
	return a.AccountCreateFn(c, roleID, companyID, locationID)
}

// IsLowerRole mock
func (a RBAC) IsLowerRole(c echo.Context, role models.AccessRole) error {
	return a.IsLowerRoleFn(c, role)
}
