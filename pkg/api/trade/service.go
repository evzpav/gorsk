package trade

import (
	"github.com/evzpav/gorsk/pkg/api/trade/platform/pgsql"
	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/labstack/echo"
)

// Service represents trade application interface
type Service interface {
	Create(echo.Context, gorsk.Trade) (*gorsk.Trade, error)
	List(echo.Context, *gorsk.Pagination) ([]gorsk.Trade, error)
	View(echo.Context, int) (*gorsk.Trade, error)
	Delete(echo.Context, int) error
	Update(echo.Context, *Update) (*gorsk.Trade, error)
}

// New creates new Trade application service
func New(db *pg.DB, tdb TradeDB, rbac RBAC, sec Securer) *Trade {
	return &Trade{db: db, tdb: tdb, rbac: rbac, sec: sec}
}

// Initialize initalizes Trade application service with defaultsUser
func Initialize(db *pg.DB, rbac RBAC, sec Securer) *Trade {
	return New(db, pgsql.NewTrade(), rbac, sec)
}

// Trade represents Trade application service
type Trade struct {
	db   *pg.DB
	tdb  TradeDB
	rbac RBAC
	sec  Securer
}

// Securer represents security interface
type Securer interface {
	Hash(string) string
}

// TradeDB represents trade repository interface
type TradeDB interface {
	Create(orm.DB, gorsk.Trade) (*gorsk.Trade, error)
	View(orm.DB, int) (*gorsk.Trade, error)
	List(orm.DB, *gorsk.ListQuery, *gorsk.Pagination) ([]gorsk.Trade, error)
	Update(orm.DB, *gorsk.Trade) error
	Delete(orm.DB, *gorsk.Trade) error
}

// RBAC represents role-based-access-control interface
type RBAC interface {
	User(echo.Context) *gorsk.AuthUser
	EnforceUser(echo.Context, int) error
	AccountCreate(echo.Context, gorsk.AccessRole, int, int) error
	IsLowerRole(echo.Context, gorsk.AccessRole) error
}
