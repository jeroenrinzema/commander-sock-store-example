package controllers

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/logic/common"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/sysco-middleware/commander"
)

// OnCartDeclined mark a cart as declined
func OnCartDeclined(command *commander.Command) *commander.Event {
	cart := models.CartModel{}
	req := models.CommandDeclinedModel{}
	err := json.Unmarshal(command.Data, &req)

	if err != nil {
		return command.NewErrorEvent("DataParseError", nil)
	}

	// Find the cart based on the given cart ID
	findQuery := common.Database.First(&cart, req.ID.String())
	if findQuery.Error != nil {
		return command.NewErrorEvent("CartNotFound", nil)
	}

	common.Database.Delete(&cart)
	return command.NewEvent("CartDeclined", 1, cart.ID, nil)
}
