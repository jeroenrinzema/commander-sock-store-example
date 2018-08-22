package main

import (
	"github.com/jeroenrinzema/commander-sock-store-example/cart/logic/common"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/logic/controllers"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	database := common.OpenDatabase()
	commander := common.OpenCommander()

	database.AutoMigrate(&models.CartModel{})

	commander.NewCommandHandle("CreateCart", controllers.OnCartCreated)
	commander.NewCommandHandle("CartPurchased", controllers.OnCartPurchased)
	commander.NewCommandHandle("CartDeclined", controllers.OnCartDeclined)

	commander.NewCommandHandle("AddItem", controllers.OnAddItem)
	commander.NewCommandHandle("RemoveItem", controllers.OnRemoveItem)

	go commander.Consume()
	commander.CloseOnSIGTERM()
}
