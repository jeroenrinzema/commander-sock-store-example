package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/logic/common"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/sysco-middleware/commander"
)

// OnAddItem adds the given item to the cart
func OnAddItem(command *commander.Command) *commander.Event {
	cart := models.CartModel{}
	req := models.ItemCommandModel{}
	err := json.Unmarshal(command.Data, &req)

	if err != nil {
		fmt.Println(err)
		return command.NewErrorEvent("DataParseError", nil)
	}

	cart.ID = req.ID

	// Find the cart based on the given cart ID
	find := common.Database.First(&cart)
	if find.Error != nil {
		return command.NewErrorEvent("CartNotFound", nil)
	}

	cart.Items = append(cart.Items, req.Item)

	save := common.Database.Save(&cart)
	if save.Error != nil {
		return command.NewErrorEvent("CartInvalid", nil)
	}

	event := models.EventItemModel{
		ID:   cart.ID,
		Item: req.Item,
	}

	res, _ := json.Marshal(event)
	return command.NewEvent("ItemAdded", 1, *cart.ID, res)
}

// OnRemoveItem removes the given item from the cart
func OnRemoveItem(command *commander.Command) *commander.Event {
	cart := models.CartModel{}
	req := models.ItemCommandModel{}
	err := json.Unmarshal(command.Data, &req)

	if err != nil {
		return command.NewErrorEvent("DataParseError", nil)
	}

	cart.ID = req.ID

	// Find the cart based on the given cart ID
	findQuery := common.Database.First(&cart)
	if findQuery.Error != nil {
		return command.NewErrorEvent("CartNotFound", nil)
	}

	// Remove item from cart
	for index, item := range cart.Items {
		if item == req.Item {
			cart.Items = append(cart.Items[:index], cart.Items[index+1:]...)
		}
	}

	saveQuery := common.Database.Save(&cart)
	if saveQuery.Error != nil {
		return command.NewErrorEvent("CartInvalid", nil)
	}

	event := models.EventItemModel{
		ID:   cart.ID,
		Item: req.Item,
	}

	res, _ := json.Marshal(event)
	return command.NewEvent("ItemRemoved", 1, *cart.ID, res)
}
