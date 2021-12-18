package main

import (
	"urlify/internal/controller/routers"
	"urlify/internal/model/infrastructure/config"
	"urlify/internal/model/infrastructure/container"
	"urlify/internal/model/infrastructure/database"
	"urlify/internal/model/infrastructure/server"
)

func main() {
	configuration := config.LoadConfig()

	db := database.Connect(&configuration.Database)

	appContainer := container.New(db)

	engine := server.New(appContainer, routers.InitializeRouters)

	engine.Run()
}
