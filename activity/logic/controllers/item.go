package controllers

import (
	uuid "github.com/satori/go.uuid"
	"github.com/sysco-middleware/commander"
)

// OnItemWatched marks the given item as watched
func OnItemWatched(command *commander.Command) *commander.Event {
	id := uuid.NewV4()
	return command.NewEvent("ItemWatched", 1, id, nil)
}
