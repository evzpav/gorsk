package pgsql

import (
	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/jinzhu/gorm"
)

// NewUser returns a new user database instance
func NewUser() *User {
	return &User{}
}

// User represents the client for user table
type User struct{}

// View returns single user by ID
func (u *User) View(db *gorm.DB, id int) (*gorsk.User, error) {
	user := &gorsk.User{Base: gorsk.Base{ID: id}}
	if err := db.First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates user's info
func (u *User) Update(db *gorm.DB, user *gorsk.User) error {
	return db.Save(user).Error
}
