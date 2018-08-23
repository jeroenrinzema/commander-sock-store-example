package models

import uuid "github.com/satori/go.uuid"

// CommandCreateModel holds the information about a cart
type CommandCreateModel struct {
	User uint `json:"user"`
}

// CommandDeclinedModel holds the information about a cart
type CommandDeclinedModel struct {
	ID *uuid.UUID `json:"id"`
}

// CommandPurcaseModel holds the information about a cart
type CommandPurcaseModel struct {
	ID *uuid.UUID `json:"id"`
}

// ItemCommandModel holds the info about a item
type ItemCommandModel struct {
	ID   *uuid.UUID `json:"id"`
	Item int64      `json:"item"`
}
