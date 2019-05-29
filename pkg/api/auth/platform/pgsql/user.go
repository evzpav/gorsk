package pgsql

import (
	"log"

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
	var user = new(gorsk.User)
	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name" 
	FROM "users" AS "user" LEFT JOIN "roles" AS "role" ON "role"."id" = "user"."role_id" 
	WHERE ("user"."id" = ?)`
	if err := db.Raw(sql, id).Scan(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// FindByUsername queries for single user by username
func (u *User) FindByUsername(db *gorm.DB, uname string) (*gorsk.User, error) {
	var user = new(gorsk.User)
	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name" 
	FROM "users" AS "user" LEFT JOIN "roles" AS "role" ON "role"."id" = "user"."role_id" 
	WHERE ("user"."username" = ?)`
	db.Preload("roles")
	if err := db.Raw(sql, uname).Scan(&user).Error; err != nil {
		return nil, err
	}
	log.Printf("USER %+v", user)
	return user, nil
}

// FindByToken queries for single user by token
func (u *User) FindByToken(db *gorm.DB, token string) (*gorsk.User, error) {
	var user = new(gorsk.User)
	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name" 
	FROM "users" AS "user" LEFT JOIN "roles" AS "role" ON "role"."id" = "user"."role_id" 
	WHERE ("user"."token" = ?)`
	if err := db.Raw(sql, token).Scan(&user).Error; err != nil {
		return nil, err
	}
	log.Printf("%+v", user)
	return user, nil
}

// Update updates user's info
func (u *User) Update(db *gorm.DB, user *gorsk.User) error {
	return db.Save(&user).Error
}
