package mock

import (
	"database/sql"
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/evzpav/gorsk/pkg/utl/postgres"

	"github.com/fortytw2/dockertest"
)

// NewPGContainer instantiates new PostgreSQL docker container
func NewPGContainer(t *testing.T) *dockertest.Container {
	container, err := dockertest.RunContainer("postgres:alpine", "5432", func(addr string) error {
		db, err := sql.Open("postgres", "postgres://postgres:postgres@"+addr+"?sslmode=disable")
		fatalErr(t, err)

		return db.Ping()
	})
	fatalErr(t, err)

	return container
}

// NewDB instantiates new postgresql database connection via docker container
func NewDB(t *testing.T, con *dockertest.Container, models ...interface{}) *gorm.DB {
	db, err := postgres.New("postgres://postgres:postgres@"+con.Addr+"/postgres?sslmode=disable", 10, false)
	fatalErr(t, err)

	for _, v := range models {
		fatalErr(t, db.CreateTable(v).Error)
	}

	return db
}

// InsertMultiple inserts multiple values into database
func InsertMultiple(db *gorm.DB, models ...interface{}) error {
	for _, v := range models {
		if err := db.Create(v).Error; err != nil {
			return err
		}
	}
	return nil
}

func fatalErr(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}
