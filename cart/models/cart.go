package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// CartModel holds the information about a cart
type CartModel struct {
	ID          *uuid.UUID `gorm:"type:uuid; primary_key"`
	User        *uuid.UUID
	Items       []*uuid.UUID
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
	ID          *uuid.UUID `gorm:"type:uuid; primary_key"`
	User        *uuid.UUID
	Items       []*uuid.UUID
	Purchased   bool
	PurchasedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName returns the table name of CartModelView
func (modal *CartModelView) TableName() string {
	return "ProjectorCartView"
}
