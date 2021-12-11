package main

import (
	"urlify/internal/application/routers"
	"urlify/internal/infrastructure/config"
	"urlify/internal/infrastructure/container"
	"urlify/internal/infrastructure/database"
	"urlify/internal/infrastructure/server"
)

func main() {
	configuration := config.LoadConfig()

	db := database.Connect(&configuration.Database)

	appContainer := container.NewContainer(db)

	engine := server.New(&appContainer, routers.InitializeRouters)

	engine.Run()
}
