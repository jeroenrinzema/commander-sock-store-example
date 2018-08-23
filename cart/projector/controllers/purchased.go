package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/projector/common"
	"github.com/sysco-middleware/commander"
)

// OnCartPurchased mark the given cart as purchased
func OnCartPurchased(event *commander.Event) {
	req := models.EventPurchaseModel{}
	err := json.Unmarshal(event.Data, &req)

	if err != nil {
		return
	}

	cart := models.CartModelView{
		ID: &event.Key,
	}

	query := common.Database.First(&cart)
	if query.Error != nil {
		fmt.Println("query error", query.Error)
		return
	}

	cart.Purchased = true
	cart.PurchasedAt = req.Time

	common.Database.Save(&cart)
}
