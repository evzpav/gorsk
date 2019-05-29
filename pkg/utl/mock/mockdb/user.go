package mockdb

import (
	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/jinzhu/gorm"
)

// User database mock
type User struct {
	CreateFn         func(*gorm.DB, gorsk.User) (*gorsk.User, error)
	ViewFn           func(*gorm.DB, int) (*gorsk.User, error)
	FindByUsernameFn func(*gorm.DB, string) (*gorsk.User, error)
	FindByTokenFn    func(*gorm.DB, string) (*gorsk.User, error)
	ListFn           func(*gorm.DB, *gorsk.ListQuery, *gorsk.Pagination) ([]gorsk.User, error)
	DeleteFn         func(*gorm.DB, *gorsk.User) error
	UpdateFn         func(*gorm.DB, *gorsk.User) error
}

// Create mock
func (u *User) Create(db *gorm.DB, usr gorsk.User) (*gorsk.User, error) {
	return u.CreateFn(db, usr)
}

// View mock
func (u *User) View(db *gorm.DB, id int) (*gorsk.User, error) {
	return u.ViewFn(db, id)
}

// FindByUsername mock
func (u *User) FindByUsername(db *gorm.DB, uname string) (*gorsk.User, error) {
	return u.FindByUsernameFn(db, uname)
}

// FindByToken mock
func (u *User) FindByToken(db *gorm.DB, token string) (*gorsk.User, error) {
	return u.FindByTokenFn(db, token)
}

// List mock
func (u *User) List(db *gorm.DB, lq *gorsk.ListQuery, p *gorsk.Pagination) ([]gorsk.User, error) {
	return u.ListFn(db, lq, p)
}

// Delete mock
func (u *User) Delete(db *gorm.DB, usr *gorsk.User) error {
	return u.DeleteFn(db, usr)
}

// Update mock
func (u *User) Update(db *gorm.DB, usr *gorsk.User) error {
	return u.UpdateFn(db, usr)
}
