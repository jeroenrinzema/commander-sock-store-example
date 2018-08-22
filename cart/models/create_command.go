package models

import uuid "github.com/satori/go.uuid"

// CommandCreateModel holds the information about a cart
type CommandCreateModel struct {
	User uuid.UUID `json:"user"`
}
