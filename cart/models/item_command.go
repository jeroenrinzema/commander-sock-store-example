package models

import uuid "github.com/satori/go.uuid"

// ItemCommandModel holds the info about a item
type ItemCommandModel struct {
	ID   *uuid.UUID `json:"id"`
	Item int64      `json:"item"`
}
