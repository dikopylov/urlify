package main

import (
	"urlify/internal/infrastructure/config"
	"urlify/internal/infrastructure/database"
	"urlify/internal/infrastructure/server"
)

func main() {
	configuration := config.LoadConfig()

	database.Connect(&configuration.Database)

	server.Serve()
}
