package pgsql

import (
	"net/http"
	"strings"

	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// NewUser returns a new user database instance
func NewUser() *User {
	return &User{}
}

// User represents the client for user table
type User struct{}

// Custom errors
var (
	ErrAlreadyExists = echo.NewHTTPError(http.StatusInternalServerError, "Username or email already exists.")
)

// Create creates a new user on database
func (u *User) Create(db *gorm.DB, usr gorsk.User) (*gorsk.User, error) {
	var user = new(gorsk.User)
	err := db.Model(user).Where("lower(username) = ? or lower(email) = ?",
		strings.ToLower(usr.Username), strings.ToLower(usr.Email)).Error

	if err != nil {
		return nil, ErrAlreadyExists

	}

	if err := db.Create(&usr).Error; err != nil {
		return nil, err
	}
	return &usr, nil
}

// View returns single user by ID
func (u *User) View(db *gorm.DB, id int) (*gorsk.User, error) {
	var user = new(gorsk.User)
	sql := `SELECT "user".*, "role"."id" AS "role__id", "role"."access_level" AS "role__access_level", "role"."name" AS "role__name" 
	FROM "users" AS "user" LEFT JOIN "roles" AS "role" ON "role"."id" = "user"."role_id" 
	WHERE ("user"."id" = ? and deleted_at is null)`
	if err := db.Raw(sql, id).Scan(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates user's contact info
func (u *User) Update(db *gorm.DB, user *gorsk.User) error {
	return db.Save(user).Error
}

// List returns list of all users retrievable for the current user, depending on role
func (u *User) List(db *gorm.DB, qp *gorsk.ListQuery, p *gorsk.Pagination) ([]gorsk.User, error) {
	var users []gorsk.User
	q := db.Select("user.*", "Role").Find(&users).Limit(p.Limit).Offset(p.Offset).Order("user.id desc")
	if qp != nil {
		q.Where(qp.Query, qp.ID)
	}
	// if err := q.Select(); err != nil {
	// 	return nil, err
	// }
	return users, nil
}

// Delete sets deleted_at for a user
func (u *User) Delete(db *gorm.DB, user *gorsk.User) error {
	return db.Delete(user).Error
}
