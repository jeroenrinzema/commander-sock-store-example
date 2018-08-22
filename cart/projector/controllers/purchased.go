package controllers

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/projector/common"
	"github.com/sysco-middleware/commander"
)

// OnCartPurchased mark the given cart as purchased
func OnCartPurchased(event *commander.Event) {
	cart := models.CartModelView{}
	req := models.CartModel{}
	err := json.Unmarshal(event.Data, &cart)
	if err != nil {
		return
	}

	cart.ID = req.ID

	query := common.Database.First(&cart)
	if query.Error != nil {
		fmt.Println("query error", query.Error)
		return
	}

	cart.Purchased = true
	cart.PurchasedAt = time.Now()

	common.Database.Save(&cart)
}
