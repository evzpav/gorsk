package postgres

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// New creates new database connection to a postgres database
func New(psn string, timeout int, enableLog bool) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", psn)
	if err != nil {
		return nil, err
	}
	db.LogMode(enableLog)
	return db, nil
}
