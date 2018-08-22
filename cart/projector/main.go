package main

import (
	"github.com/jeroenrinzema/commander-sock-store-example/cart/models"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/projector/common"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/projector/controllers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	commander := common.OpenCommander()
	database := common.OpenDatabase()

	database.AutoMigrate(&models.CartModelView{})

	commander.NewEventHandle("CartCreated", []int{1}, controllers.OnCartCreated)
	commander.NewEventHandle("CartPurchased", []int{1}, controllers.OnCartPurchased)
	commander.NewEventHandle("CartDeclined", []int{1}, controllers.OnCartDeclined)

	commander.NewEventHandle("ItemAdded", []int{1}, controllers.OnItemAdded)
	commander.NewEventHandle("ItemRemoved", []int{1}, controllers.OnItemRemoved)

	go commander.Consume()
	commander.CloseOnSIGTERM()
}
