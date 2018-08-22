package models

import uuid "github.com/satori/go.uuid"

// CommandPurcaseModel holds the information about a cart
type CommandPurcaseModel struct {
	ID uuid.UUID `json:"id"`
}
