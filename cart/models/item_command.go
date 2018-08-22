package models

import uuid "github.com/satori/go.uuid"

type ItemCommandModel struct {
	ID *uuid.UUID `json:"id"`
}
