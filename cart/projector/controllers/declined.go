package controllers

import (
	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/projector/common"
	"github.com/sysco-middleware/commander"
)

// OnCartDeclined mark the given cart as declined
func OnCartDeclined(event *commander.Event) {
	cart := models.CartModelView{
		ID: &event.Key,
	}

	common.Database.Delete(&cart)
}
