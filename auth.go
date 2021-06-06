package gorsk

import (
	"github.com/epavanello/gorsk/pkg/models"
	"github.com/labstack/echo"
)

// AuthToken holds authentication token details with refresh token
type AuthToken struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshToken holds authentication token details
type RefreshToken struct {
	Token string `json:"token"`
}

// RBACService represents role-based access control service interface
type RBACService interface {
	User(echo.Context) models.AuthUser
	EnforceRole(echo.Context, models.AccessRole) error
	EnforceUser(echo.Context, int) error
	EnforceCompany(echo.Context, int) error
	EnforceLocation(echo.Context, int) error
	AccountCreate(echo.Context, models.AccessRole, int, int) error
	IsLowerRole(echo.Context, models.AccessRole) error
}
