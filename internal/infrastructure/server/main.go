package server

import (
	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	engine := gin.Default()

	initializeRouters(engine)

	return engine
}

func initializeRouters(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
