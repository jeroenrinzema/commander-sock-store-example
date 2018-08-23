package controllers

import (
	"encoding/json"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/projector/common"
	"github.com/sysco-middleware/commander"
)

// OnItemAdded adds a item to a cart
func OnItemAdded(event *commander.Event) {
	req := models.EventItemModel{}
	err := json.Unmarshal(event.Data, &req)

	if err != nil {
		return
	}

	cart := models.CartModelView{
		ID: &event.Key,
	}

	query := common.Database.First(&cart)
	if query.Error != nil {
		return
	}

	cart.Items = append(cart.Items, req.Item)
	common.Database.Save(&cart)
}

// OnItemRemoved removes a item from a cart
func OnItemRemoved(event *commander.Event) {
	req := models.EventItemModel{}
	err := json.Unmarshal(event.Data, &req)

	if err != nil {
		return
	}

	cart := models.CartModelView{
		ID: &event.Key,
	}

	query := common.Database.First(&cart)
	if query.Error != nil {
		return
	}

	for index, item := range cart.Items {
		if item == req.Item {
			cart.Items = append(cart.Items[:index], cart.Items[index+1:]...)
		}
	}

	common.Database.Save(&cart)
}
