package controllers

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/logic/common"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/sysco-middleware/commander"
)

// OnCartCreated creates a new empty cart
func OnCartCreated(command *commander.Command) *commander.Event {
	cart := models.CartModel{}
	req := models.CommandCreateModel{}
	err := json.Unmarshal(command.Data, &req)

	if err != nil {
		return command.NewErrorEvent("DataParseError", nil)
	}

	cart.User = req.User

	query := common.Database.Save(&cart)
	if query.Error != nil {
		return command.NewErrorEvent("CartInvalid", nil)
	}

	event := models.EventCreatedModel{
		ID:    cart.ID,
		User:  cart.User,
		Items: cart.Items,
	}

	res, _ := json.Marshal(event)
	return command.NewEvent("CartCreated", 1, *cart.ID, res)
}
