package pgsql_test

import (
	"testing"

	gorsk "github.com/evzpav/gorsk/pkg/utl/model"

	"github.com/evzpav/gorsk/pkg/api/trade/platform/pgsql"
	"github.com/evzpav/gorsk/pkg/utl/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	cases := []struct {
		name     string
		wantErr  bool
		req      gorsk.Trade
		wantData *gorsk.Trade
	}{

		{
			name: "Success",
			req: gorsk.Trade{
				EntryPrice: 10.4,
			},
			wantData: &gorsk.Trade{
				ID:         1,
				EntryPrice: 10.4,
			},
		},
	}

	dbCon := mock.NewPGContainer(t)
	defer dbCon.Shutdown()

	db := mock.NewDB(t, dbCon, &gorsk.Trade{})

	tdb := pgsql.NewTrade()

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := tdb.Create(db, tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantData != nil {
				if resp == nil {
					t.Error("Expected data, but received nil.")
					return
				}
				tt.wantData.CreatedAt = resp.CreatedAt
				tt.wantData.UpdatedAt = resp.UpdatedAt
				assert.Equal(t, tt.wantData, resp)
			}
		})
	}
}

// func TestView(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		id       int
// 		wantData *gorsk.User
// 	}{
// 		{
// 			name:    "User does not exist",
// 			wantErr: true,
// 			id:      1000,
// 		},
// 		{
// 			name: "Success",
// 			id:   2,
// 			wantData: &gorsk.User{
// 				Email:      "tomjones@mail.com",
// 				FirstName:  "Tom",
// 				LastName:   "Jones",
// 				Username:   "tomjones",
// 				RoleID:     1,
// 				CompanyID:  1,
// 				LocationID: 1,
// 				Password:   "newPass",
// 				Base: gorsk.Base{
// 					ID: 2,
// 				},
// 				Role: &gorsk.Role{
// 					ID:          1,
// 					AccessLevel: 1,
// 					Name:        "SUPER_ADMIN",
// 				},
// 			},
// 		},
// 	}

// 	dbCon := mock.NewPGContainer(t)
// 	defer dbCon.Shutdown()

// 	db := mock.NewDB(t, dbCon, &gorsk.Role{}, &gorsk.User{})

// 	if err := mock.InsertMultiple(db, &gorsk.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, cases[1].wantData); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewUser()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			user, err := udb.View(db, tt.id)
// 			assert.Equal(t, tt.wantErr, err != nil)
// 			if tt.wantData != nil {
// 				if user == nil {
// 					t.Errorf("response was nil due to: %v", err)
// 				} else {
// 					tt.wantData.CreatedAt = user.CreatedAt
// 					tt.wantData.UpdatedAt = user.UpdatedAt
// 					assert.Equal(t, tt.wantData, user)
// 				}
// 			}
// 		})
// 	}
// }

// func TestUpdate(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		usr      *gorsk.User
// 		wantData *gorsk.User
// 	}{
// 		{
// 			name: "Success",
// 			usr: &gorsk.User{
// 				Base: gorsk.Base{
// 					ID: 2,
// 				},
// 				FirstName: "Z",
// 				LastName:  "Freak",
// 				Address:   "Address",
// 				Phone:     "123456",
// 				Mobile:    "345678",
// 				Username:  "newUsername",
// 			},
// 			wantData: &gorsk.User{
// 				Email:      "tomjones@mail.com",
// 				FirstName:  "Z",
// 				LastName:   "Freak",
// 				Username:   "tomjones",
// 				RoleID:     1,
// 				CompanyID:  1,
// 				LocationID: 1,
// 				Password:   "newPass",
// 				Address:    "Address",
// 				Phone:      "123456",
// 				Mobile:     "345678",
// 				Base: gorsk.Base{
// 					ID: 2,
// 				},
// 			},
// 		},
// 	}

// 	dbCon := mock.NewPGContainer(t)
// 	defer dbCon.Shutdown()

// 	db := mock.NewDB(t, dbCon, &gorsk.Role{}, &gorsk.User{})

// 	if err := mock.InsertMultiple(db, &gorsk.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, cases[0].usr); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewUser()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {
// 			err := udb.Update(db, tt.wantData)
// 			assert.Equal(t, tt.wantErr, err != nil)
// 			if tt.wantData != nil {
// 				user := &gorsk.User{
// 					Base: gorsk.Base{
// 						ID: tt.usr.ID,
// 					},
// 				}
// 				if err := db.Select(user); err != nil {
// 					t.Error(err)
// 				}
// 				tt.wantData.UpdatedAt = user.UpdatedAt
// 				tt.wantData.CreatedAt = user.CreatedAt
// 				tt.wantData.LastLogin = user.LastLogin
// 				tt.wantData.DeletedAt = user.DeletedAt
// 				assert.Equal(t, tt.wantData, user)
// 			}
// 		})
// 	}
// }

func TestList(t *testing.T) {
	cases := []struct {
		name     string
		wantErr  bool
		qp       *gorsk.ListQuery
		pg       *gorsk.Pagination
		wantData []gorsk.Trade
	}{
		{
			name:    "Invalid pagination values",
			wantErr: true,
			pg: &gorsk.Pagination{
				Limit: -100,
			},
		},
		{
			name: "Success",
			pg: &gorsk.Pagination{
				Limit:  100,
				Offset: 0,
			},
			wantData: []gorsk.Trade{
				{
					EntryPrice: 10.4,
				},
				{
					EntryPrice: 12.4,
				},
			},
		},
	}

	dbCon := mock.NewPGContainer(t)
	defer dbCon.Shutdown()

	db := mock.NewDB(t, dbCon, &gorsk.Role{}, &gorsk.Trade{})

	if err := mock.InsertMultiple(db, &gorsk.Role{
		ID:          1,
		AccessLevel: 1,
		Name:        "SUPER_ADMIN"}, &cases[1].wantData); err != nil {
		t.Error(err)
	}

	tdb := pgsql.NewTrade()

	for _, tt := range cases {

		t.Run(tt.name, func(t *testing.T) {
			trades, err := tdb.List(db, tt.qp, tt.pg)
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantData != nil {
				for i, v := range trades {
					tt.wantData[i].CreatedAt = v.CreatedAt
					tt.wantData[i].UpdatedAt = v.UpdatedAt
				}
				assert.Equal(t, tt.wantData, trades)
			}
		})
	}
}

// func TestDelete(t *testing.T) {
// 	cases := []struct {
// 		name     string
// 		wantErr  bool
// 		usr      *gorsk.User
// 		wantData *gorsk.User
// 	}{
// 		{
// 			name: "Success",
// 			usr: &gorsk.User{
// 				Base: gorsk.Base{
// 					ID:        2,
// 					DeletedAt: mock.TestTime(2018),
// 				},
// 			},
// 			wantData: &gorsk.User{
// 				Email:      "tomjones@mail.com",
// 				FirstName:  "Tom",
// 				LastName:   "Jones",
// 				Username:   "tomjones",
// 				RoleID:     1,
// 				CompanyID:  1,
// 				LocationID: 1,
// 				Password:   "newPass",
// 				Base: gorsk.Base{
// 					ID: 2,
// 				},
// 			},
// 		},
// 	}

// 	dbCon := mock.NewPGContainer(t)
// 	defer dbCon.Shutdown()

// 	db := mock.NewDB(t, dbCon, &gorsk.Role{}, &gorsk.User{})

// 	if err := mock.InsertMultiple(db, &gorsk.Role{
// 		ID:          1,
// 		AccessLevel: 1,
// 		Name:        "SUPER_ADMIN"}, cases[0].wantData); err != nil {
// 		t.Error(err)
// 	}

// 	udb := pgsql.NewUser()

// 	for _, tt := range cases {
// 		t.Run(tt.name, func(t *testing.T) {

// 			err := udb.Delete(db, tt.usr)
// 			assert.Equal(t, tt.wantErr, err != nil)

// 			// Check if the deleted_at was set
// 		})
// 	}
// }
