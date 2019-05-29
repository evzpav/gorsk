package pgsql

import (
	gorsk "github.com/evzpav/gorsk/pkg/utl/model"
	"github.com/jinzhu/gorm"
)

// NewTrade returns a new trade database instance
func NewTrade() *Trade {
	return &Trade{}
}

// Trade represents the client for trade table
type Trade struct{}

// Create creates a new trade on database
func (t *Trade) Create(db *gorm.DB, trade gorsk.Trade) (*gorsk.Trade, error) {
	if err := db.Create(&trade).Error; err != nil {
		return nil, err
	}
	return &trade, nil
}

// View returns single trade by ID
func (t *Trade) View(db *gorm.DB, id int) (*gorsk.Trade, error) {
	var trade = new(gorsk.Trade)
	if err := db.First(trade, id).Error; err != nil {
		return nil, err
	}

	return trade, nil
}

// Update updates trade info
func (t *Trade) Update(db *gorm.DB, trade *gorsk.Trade) error {
	return db.Save(trade).Error
}

// List returns list of all trades
func (t *Trade) List(db *gorm.DB, qp *gorsk.ListQuery, p *gorsk.Pagination) ([]gorsk.Trade, error) {
	var trades []gorsk.Trade
	q := db.Select("*").Find(&trades).Limit(p.Limit).Offset(p.Offset).Order("trade.id desc")
	if qp != nil {
		q.Where(qp.Query, qp.ID)
	}
	return trades, nil
}

// Delete sets deleted_at for a trade
func (t *Trade) Delete(db *gorm.DB, trade *gorsk.Trade) error {
	return db.Delete(trade).Error
}
