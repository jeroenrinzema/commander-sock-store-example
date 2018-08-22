package models

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

// CartModel holds the information about a cart
type CartModel struct {
	ID          *uuid.UUID `gorm:"type:uuid; primary_key"`
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

func (modal *CartModel) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}

// CartModelView holds the information about a cart
type CartModelView struct {
	ID          *uuid.UUID `gorm:"type:uuid; primary_key"`
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

func (modal *CartModelView) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("ID", uuid.NewV4())
	return nil
}
