package models

import uuid "github.com/satori/go.uuid"

// EventCreatedModel holds the information about a cart
type EventCreatedModel struct {
	ID    *uuid.UUID `json:"id"`
	User  uint       `json:"user"`
	Items []int64    `json:"items"`
}

// EventDeclinedModel holds the information about a cart
type EventDeclinedModel struct {
	ID *uuid.UUID `json:"id"`
}

// EventPurcaseModel holds the information about a cart
type EventPurcaseModel struct {
	ID *uuid.UUID `json:"id"`
}

// EventItemModel holds the info about a item
type EventItemModel struct {
	ID   *uuid.UUID `json:"id"`
	Item int64      `json:"item"`
}
