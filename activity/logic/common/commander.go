package controllers

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sysco-middleware/commander"
)

// OnItemSelected creates a new empty cart
func OnItemSelected(command *commander.Command) *commander.Event {
	id, _ := uuid.NewV4()
	return command.NewEvent("ItemSelected", 1, id, nil)
}
