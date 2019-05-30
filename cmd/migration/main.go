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

	var psn = os.Getenv("POSTGRES_URL")
	db, err := gorm.Open("postgres", psn)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	createSchema(db, &gorsk.User{}, &gorsk.Company{}, &gorsk.Location{}, &gorsk.Role{})
	dbInsert := `INSERT INTO public.companies VALUES (1, now(), now(), 'admin_company', true);
	INSERT INTO public.locations VALUES (1, now(), now(), 'admin_location', true, 'admin_address', 1);
	INSERT INTO public.roles VALUES (100, 100, 'SUPER_ADMIN');
	INSERT INTO public.roles VALUES (110, 110, 'ADMIN');
	INSERT INTO public.roles VALUES (120, 120, 'COMPANY_ADMIN');
	INSERT INTO public.roles VALUES (130, 130, 'LOCATION_ADMIN');
	INSERT INTO public.roles VALUES (200, 200, 'USER');`

	queries := strings.Split(dbInsert, ";")

	for _, v := range queries[0 : len(queries)-1] {
		err := db.Exec(v).Error
		if err != nil {
			log.Fatal(err)
		}
	}

	userInsert := `INSERT INTO public.users (id, first_name, last_name, username, password, email, active, role_id, company_id, location_id) VALUES (1,'Admin', 'Admin', 'admin', '%s', 'johndoe@mail.com', true, 100, 1, 1);`
	err = db.Exec(fmt.Sprintf(userInsert, secure.New(1, nil).Hash("admin"))).Error
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Initial migration completed")

}

func createSchema(db *gorm.DB, models ...interface{}) {
	for _, model := range models {
		err := db.CreateTable(model).Error
		if err != nil {
			log.Printf("Error creating table: %+v", err)
		}
	}
}
