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

	saveQuery := common.Database.Save(&cart)
	if saveQuery.Error != nil {
		return command.NewErrorEvent("CartInvalid", nil)
	}

	res, _ := json.Marshal(cart)
	return command.NewEvent("CartPurchased", 1, cart.ID, res)
}
