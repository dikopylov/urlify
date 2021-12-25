package main

import (
	"urlify/internal/controller/routers"
	"urlify/internal/model/infrastructure/config"
	"urlify/internal/model/infrastructure/container"
	"urlify/internal/model/infrastructure/database"
	"urlify/internal/model/infrastructure/server"
)

var AppContainer *container.Container

func main() {
	configuration := config.LoadConfig()

	db := database.Connect(&configuration.Database)

	AppContainer = container.New(db)

	engine := server.New(routers.InitializeRouters)

	engine.Run()
}
