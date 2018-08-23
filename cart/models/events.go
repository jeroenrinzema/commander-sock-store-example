package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// EventCreatedModel holds the information about a cart
type EventCreatedModel struct {
	ID    *uuid.UUID `json:"id"`
	User  uint       `json:"user"`
	Items []int64    `json:"items"`
}

// EventPurcaseModel holds the information about a purchase
type EventPurcaseModel struct {
	ID   *uuid.UUID `json:"id"`
	Time time.Time  `json:"time"`
}

// EventItemModel holds the info about a item
type EventItemModel struct {
	ID   *uuid.UUID `json:"id"`
	Item int64      `json:"item"`
}
