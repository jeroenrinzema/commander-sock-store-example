package main

import (
	"github.com/jeroenrinzema/commander-sock-store-example/activity/logic/common"
	"github.com/jeroenrinzema/commander-sock-store-example/activity/logic/controllers"
	"github.com/jeroenrinzema/commander-sock-store-example/cart/logic/common"
)

func main() {
	commander := common.OpenCommander()

	commander.NewCommandHandle("ItemSelected", controllers.OnItemSelected)

	go commander.Consume()
	commander.CloseOnSIGTERM()
}
