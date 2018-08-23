package controllers

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/projector/common"
	"github.com/sysco-middleware/commander"
)

// OnCartCreated creates the card
func OnCartCreated(event *commander.Event) {
	req := models.EventCreatedModel{}
	err := json.Unmarshal(event.Data, &req)
	if err != nil {
		return
	}

	cart := models.CartModelView{
		ID:    req.ID,
		User:  req.User,
		Items: req.Items,
	}

	query := common.Database.Save(&cart)
	if query.Error != nil {
		return
	}
}
