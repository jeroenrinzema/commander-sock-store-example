package models

import (
	"time"

	"github.com/lib/pq"
)

// CartModel holds the information about a cart
type CartModel struct {
	ID          uint `gorm:"primary_key"`
	User        uint
	Items       pq.Int64Array `gorm:"type:integer[]"`
	Purchased   bool
	PurchasedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName returns the table name of CartModel
func (modal *CartModel) TableName() string {
	return "ServiceCartView"
}

// CartModelView holds the information about a cart
type CartModelView struct {
	ID          uint `gorm:"primary_key"`
	User        uint
	Items       pq.Int64Array `gorm:"type:integer[]"`
	Purchased   bool
	PurchasedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName returns the table name of CartModelView
func (modal *CartModelView) TableName() string {
	return "ProjectorCartView"
}
