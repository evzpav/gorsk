package mockdb

import (
	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/jinzhu/gorm"
)

// Trade database mock
type Trade struct {
	CreateFn          func(*gorm.DB, gorsk.Trade) (*gorsk.Trade, error)
	ViewFn            func(*gorm.DB, int) (*gorsk.Trade, error)
	FindByTradenameFn func(*gorm.DB, string) (*gorsk.Trade, error)
	FindByTokenFn     func(*gorm.DB, string) (*gorsk.Trade, error)
	ListFn            func(*gorm.DB, *gorsk.ListQuery, *gorsk.Pagination) ([]gorsk.Trade, error)
	DeleteFn          func(*gorm.DB, *gorsk.Trade) error
	UpdateFn          func(*gorm.DB, *gorsk.Trade) error
}

// Create mock
func (u *Trade) Create(db *gorm.DB, usr gorsk.Trade) (*gorsk.Trade, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *Trade) View(db *gorm.DB, id int) (*gorsk.Trade, error) {
	return u.ViewFn(db, id)
}

// FindByToken mock
func (u *Trade) FindByToken(db *gorm.DB, token string) (*gorsk.Trade, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *Trade) List(db *gorm.DB, lq *gorsk.ListQuery, p *gorsk.Pagination) ([]gorsk.Trade, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *Trade) Delete(db *gorm.DB, usr *gorsk.Trade) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *Trade) Update(db *gorm.DB, usr *gorsk.Trade) error {
	return u.UpdateFn(db, usr)
}
