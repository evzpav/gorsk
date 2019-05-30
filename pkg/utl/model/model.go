package gorsk

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Base contains common fields for all tables
type Base struct {
	ID        int       `json:"id" gorm:"auto_increment;PRIMARY_KEY"`
	CreatedAt time.Time `json:"created_at" gorm:"default:now()"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:now()"`
}

// ListQuery holds company/location data used for list db queries
type ListQuery struct {
	Query string
	ID    int
}

// BeforeInsert hooks into insert operations, setting createdAt and updatedAt to current time
func (b *Base) BeforeInsert(_ *gorm.DB) error {
	now := time.Now()
	b.CreatedAt = now
	b.UpdatedAt = now
	return nil
}

// BeforeUpdate hooks into update operations, setting updatedAt to current time
func (b *Base) BeforeUpdate(_ *gorm.DB) error {
	b.UpdatedAt = time.Now()
	return nil
}
