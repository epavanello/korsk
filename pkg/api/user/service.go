package user

import (
	"github.com/epavanello/gorsk/pkg/models"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"github.com/labstack/echo"

	"github.com/epavanello/gorsk"
	"github.com/epavanello/gorsk/pkg/api/user/platform/pgsql"
)

// Service represents user application interface
type Service interface {
	Create(echo.Context, models.User) (models.User, error)
	List(echo.Context, gorsk.Pagination) ([]models.User, error)
	View(echo.Context, int) (models.User, error)
	Delete(echo.Context, int) error
	Update(echo.Context, Update) (models.User, error)
}

// New creates new user application service
func New(db *pg.DB, udb UDB, rbac RBAC, sec Securer) *User {
	return &User{db: db, udb: udb, rbac: rbac, sec: sec}
}

// Initialize initalizes User application service with defaults
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *User {
	return New(db, pgsql.User{}, rbac, sec)
}

// User represents user application service
type User struct {
	db   *pg.DB
	udb  UDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// UDB represents user repository interface
type UDB interface {
	Create(orm.DB, models.User) (models.User, error)
	View(orm.DB, int) (models.User, error)
	List(orm.DB, *models.ListQuery, gorsk.Pagination) ([]models.User, error)
	Update(orm.DB, models.User) error
	Delete(orm.DB, models.User) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) models.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, models.AccessRole, int, int) error
	IsLowerRole(echo.Context, models.AccessRole) error
}
