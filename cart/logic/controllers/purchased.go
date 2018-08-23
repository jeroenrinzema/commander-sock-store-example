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
	req := models.CommandPurcaseModel{}
	err := json.Unmarshal(command.Data, &req)

	if err != nil {
		return command.NewErrorEvent("DataParseError", nil)
	}

	cart := models.CartModel{
		ID: req.ID,
	}

	// Find the cart based on the given cart ID
	find := common.Database.First(&cart)
	if find.Error != nil {
		return command.NewErrorEvent("CartNotFound", nil)
	}

	cart.Purchased = true
	cart.PurchasedAt = time.Now()

	save := common.Database.Save(&cart)
	if save.Error != nil {
		return command.NewErrorEvent("CartInvalid", nil)
	}

	event := models.EventPurchaseModel{
		ID:   cart.ID,
		Time: cart.PurchasedAt,
	}

	res, _ := json.Marshal(event)
	return command.NewEvent("CartPurchased", 1, *cart.ID, res)
}
