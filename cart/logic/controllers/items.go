package controllers

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/logic/common"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/sysco-middleware/commander"
)

// OnItemAdded adds the given item to the cart
func OnItemAdded(command *commander.Command) *commander.Event {
	cart := models.CartModel{}
	req := models.ItemCommandModel{}
	err := json.Unmarshal(command.Data, &req)

	if err != nil {
		return command.NewErrorEvent("DataParseError", nil)
	}

	// Find the cart based on the given cart ID
	findQuery := common.Database.First(&cart, req.ID)
	if findQuery.Error != nil {
		return command.NewErrorEvent("CartNotFound", nil)
	}

	cart.Items = append(cart.Items, req.ID)

	saveQuery := common.Database.Save(&cart)
	if saveQuery.Error != nil {
		return command.NewErrorEvent("CartInvalid", nil)
	}

	res, _ := json.Marshal(req)
	return command.NewEvent("ItemAdded", 1, cart.ID, res)
}

// OnItemRemoved removes the given item from the cart
func OnItemRemoved(command *commander.Command) *commander.Event {
	cart := models.CartModel{}
	req := models.ItemCommandModel{}
	err := json.Unmarshal(command.Data, &req)

	if err != nil {
		return command.NewErrorEvent("DataParseError", nil)
	}

	// Find the cart based on the given cart ID
	findQuery := common.Database.First(&cart, req.ID)
	if findQuery.Error != nil {
		return command.NewErrorEvent("CartNotFound", nil)
	}

	// Remove item from cart
	for index, item := range cart.Items {
		if item == req.ID {
			cart.Items = append(cart.Items[:index], cart.Items[index+1:]...)
		}
	}

	saveQuery := common.Database.Save(&cart)
	if saveQuery.Error != nil {
		return command.NewErrorEvent("CartInvalid", nil)
	}

	res, _ := json.Marshal(req)
	return command.NewEvent("ItemRemoved", 1, cart.ID, res)
}
