package password

import (
	"github.com/evzpav/gorsk/pkg/api/password/platform/pgsql"
	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Service represents password application interface
type Service interface {
	Change(echo.Context, int, string, string) error
}

// New creates new password application service
func New(db *gorm.DB, udb UserDB, rbac RBAC, sec Securer) *Password {
	return &Password{
		db:   db,
		udb:  udb,
		rbac: rbac,
		sec:  sec,
	}
}

// Initialize initalizes password application service with defaults
func Initialize(db *gorm.DB, rbac RBAC, sec Securer) *Password {
	return New(db, pgsql.NewUser(), rbac, sec)
}

// Password represents password application service
type Password struct {
	db   *gorm.DB
	udb  UserDB
	rbac RBAC
	sec  Securer
}

// UserDB represents user repository interface
type UserDB interface {
	View(*gorm.DB, int) (*gorsk.User, error)
	Update(*gorm.DB, *gorsk.User) error
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
	HashMatchesPassword(string, string) bool
	Password(string, ...string) bool
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	EnforceUser(echo.Context, int) error
}
