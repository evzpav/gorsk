package mockdb

import (
	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/go-pg/pg/orm"
)

// Trade database mock
type Trade struct {
	CreateFn          func(orm.DB, gorsk.Trade) (*gorsk.Trade, error)
	ViewFn            func(orm.DB, int) (*gorsk.Trade, error)
	FindByTradenameFn func(orm.DB, string) (*gorsk.Trade, error)
	FindByTokenFn     func(orm.DB, string) (*gorsk.Trade, error)
	ListFn            func(orm.DB, *gorsk.ListQuery, *gorsk.Pagination) ([]gorsk.Trade, error)
	DeleteFn          func(orm.DB, *gorsk.Trade) error
	UpdateFn          func(orm.DB, *gorsk.Trade) error
}

// Create mock
func (u *Trade) Create(db orm.DB, usr gorsk.Trade) (*gorsk.Trade, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *Trade) View(db orm.DB, id int) (*gorsk.Trade, error) {
	return u.ViewFn(db, id)
}

// FindByToken mock
func (u *Trade) FindByToken(db orm.DB, token string) (*gorsk.Trade, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *Trade) List(db orm.DB, lq *gorsk.ListQuery, p *gorsk.Pagination) ([]gorsk.Trade, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *Trade) Delete(db orm.DB, usr *gorsk.Trade) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *Trade) Update(db orm.DB, usr *gorsk.Trade) error {
	return u.UpdateFn(db, usr)
}
