// Package trade contains trade application services
package trade

import (
	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/evzpav/gorsk/pkg/utl/query"
	"github.com/evzpav/gorsk/pkg/utl/structs"
	"github.com/labstack/echo"
)

// Create creates a new trade
func (t *Trade) Create(c echo.Context, req gorsk.Trade) (*gorsk.Trade, error) {
	return t.tdb.Create(t.db, req)
}

// List returns list of trades
func (t *Trade) List(c echo.Context, p *gorsk.Pagination) ([]gorsk.Trade, error) {
	au := t.rbac.User(c)
	q, err := query.List(au)
	if err != nil {
		return nil, err
	}
	return t.tdb.List(t.db, q, p)
}

// View returns single user
func (t *Trade) View(c echo.Context, id int) (*gorsk.Trade, error) {
	return t.tdb.View(t.db, id)
}

// Delete deletes a user
func (t *Trade) Delete(c echo.Context, id int) error {
	trade, err := t.tdb.View(t.db, id)
	if err != nil {
		return err
	}
	return t.tdb.Delete(t.db, trade)
}

// Update updates user's contact information
func (t *Trade) Update(c echo.Context, req *gorsk.Trade) (*gorsk.Trade, error) {
	if err := t.rbac.EnforceUser(c, req.ID); err != nil {
		return nil, err
	}

	trade, err := t.tdb.View(t.db, req.ID)
	if err != nil {
		return nil, err
	}
	req.ChangeHistory = trade.ChangeHistory
	req.Atr = trade.Atr

	structs.Merge(trade, req)

	if err := t.tdb.Update(t.db, trade); err != nil {
		return nil, err
	}

	return trade, nil
}
