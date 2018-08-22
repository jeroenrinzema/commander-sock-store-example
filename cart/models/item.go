package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// ItemModel holds the information about a item
type ItemModel struct {
	ID          *uuid.UUID `gorm:"type:uuid; primary_key"`
	Name        string     `gorm:"not null;unique"`
	Description string
	WarehouseID string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName returns the table name of UserModal
func (modal *ItemModel) TableName() string {
	return "ServiceItemsView"
}
