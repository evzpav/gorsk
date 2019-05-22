package pgsql

import (
	"log"

	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/go-pg/pg/orm"
)

// NewTrade returns a new trade database instance
func NewTrade() *Trade {
	return &Trade{}
}

// Trade represents the client for trade table
type Trade struct{}

// Create creates a new trade on database
func (t *Trade) Create(db orm.DB, trade gorsk.Trade) (*gorsk.Trade, error) {
	if err := db.Insert(&trade); err != nil {
		return nil, err
	}
	return &trade, nil
}

// View returns single trade by ID
func (t *Trade) View(db orm.DB, id int) (*gorsk.Trade, error) {
	var trade = new(gorsk.Trade)
	err := db.Model(trade).Where("id = (?)", id).Select()
	if err != nil {
		return nil, err
	}

	return trade, nil
}

// Update updates trade info
func (t *Trade) Update(db orm.DB, trade *gorsk.Trade) error {
	log.Println("update trade:", trade)
	return db.Update(trade)
}

// List returns list of all trades
func (t *Trade) List(db orm.DB, qp *gorsk.ListQuery, p *gorsk.Pagination) ([]gorsk.Trade, error) {
	var trades []gorsk.Trade
	q := db.Model(&trades).Column("*").Limit(p.Limit).Offset(p.Offset).Order("trade.id desc")
	if qp != nil {
		q.Where(qp.Query, qp.ID)
	}
	if err := q.Select(); err != nil {
		return nil, err
	}
	return trades, nil
}

// Delete sets deleted_at for a trade
func (t *Trade) Delete(db orm.DB, trade *gorsk.Trade) error {
	return db.Delete(trade)
}
