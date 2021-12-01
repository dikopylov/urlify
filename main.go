package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"urlify/app/config"
	"urlify/app/database"
	"urlify/app/server"
)

func main() {
	sureLogToFile()

	configuration := config.NewConfig(*inputFile)

	database.Connect(configuration)

	server.Serve()
}

func sureLogToFile() {
	if *stdout != "" {
		f, _ := os.Create(*stdout)
		gin.DefaultWriter = io.MultiWriter(f)
	}
}
