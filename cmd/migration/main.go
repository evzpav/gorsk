package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/evzpav/gorsk/pkg/utl/secure"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	dbInsert := `INSERT INTO public.companies VALUES (1, now(), now(), 'admin_company', true);
	INSERT INTO public.locations VALUES (1, now(), now(), 'headquarters', true, 'admin_address', 1);
	INSERT INTO public.roles VALUES (100, 100, 'SUPER_ADMIN');
	INSERT INTO public.roles VALUES (110, 110, 'ADMIN');
	INSERT INTO public.roles VALUES (120, 120, 'COMPANY_ADMIN');
	INSERT INTO public.roles VALUES (130, 130, 'LOCATION_ADMIN');
	INSERT INTO public.roles VALUES (200, 200, 'USER');`
	var psn = os.Getenv("POSTGRES_URL")
	queries := strings.Split(dbInsert, ";")
	db, err := gorm.Open("postgres", psn)
	checkErr(err)

	err = db.Exec("SELECT 1").Error
	checkErr(err)

	createSchema(db, &gorsk.Company{}, &gorsk.Location{}, &gorsk.Role{}, &gorsk.User{}, &gorsk.Trade{})

	for _, v := range queries[0 : len(queries)-1] {
		err := db.Exec(v).Error
		checkErr(err)
	}

	sec := secure.New(1, nil)

	userInsert := `INSERT INTO public.users (created_at, updated_at, first_name, last_name, username, password, email, active, role_id, company_id, location_id) VALUES (now(),now(),'Admin', 'Admin', 'admin', '%s', 'johndoe@mail.com', true, 100, 1, 1);`
	err = db.Exec(fmt.Sprintf(userInsert, sec.Hash("admin"))).Error
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func createSchema(db *gorm.DB, models ...interface{}) {
	for _, model := range models {
		err := db.CreateTable(model).Error
		checkErr(err)
	}
}
