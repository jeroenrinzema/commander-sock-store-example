package controllers

import (
	"encoding/json"
	"time"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/logic/common"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/sysco-middleware/commander"
)

// OnCartPurchased mark the given cart as purchased
func OnCartPurchased(command *commander.Command) *commander.Event {
	cart := models.CartModel{}
	req := models.CommandPurcaseModel{}
	err := json.Unmarshal(command.Data, &req)

	if err != nil {
		return command.NewErrorEvent("DataParseError", nil)
	}

	// Find the cart based on the given cart ID
	findQuery := common.Database.First(&cart, req.ID.String())
	if findQuery.Error != nil {
		return command.NewErrorEvent("CartNotFound", nil)
	}

	cart.Purchased = true
	cart.PurchasedAt = time.Now()

	saveQuery := common.Database.Save(&cart)
	if saveQuery.Error != nil {
		return command.NewErrorEvent("CartInvalid", nil)
	}

	res, _ := json.Marshal(cart)
	return command.NewEvent("CartPurchased", 1, cart.ID, res)
}
