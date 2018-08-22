package main

import (
	"github.com/jeroenrinzema/commander-sock-store-example/activity/logic/common"
	"github.com/jeroenrinzema/commander-sock-store-example/activity/logic/controllers"
)

func main() {
	commander := common.OpenCommander()

	commander.NewCommandHandle("ItemSelected", controllers.OnItemWatched)

	go commander.Consume()
	commander.CloseOnSIGTERM()
}
