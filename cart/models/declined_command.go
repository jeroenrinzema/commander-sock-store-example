package models

import uuid "github.com/satori/go.uuid"

// CommandDeclinedModel holds the information about a cart
type CommandDeclinedModel struct {
	ID *uuid.UUID `json:"id"`
}
