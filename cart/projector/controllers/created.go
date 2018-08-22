package controllers

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/projector/common"
	"github.com/sysco-middleware/commander"
)

// OnCartCreated creates the card
func OnCartCreated(event *commander.Event) {
	cart := models.CartModelView{}
	err := json.Unmarshal(event.Data, &cart)
	if err != nil {
		return
	}

	query := common.Database.Save(&cart)
	if query.Error != nil {
		return
	}
}
