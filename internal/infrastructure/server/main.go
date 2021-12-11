package server

import (
	"github.com/gin-gonic/gin"
	"urlify/internal/infrastructure/container"
)

type routerFunction func(engine *gin.Engine, container *container.Container)

func New(container *container.Container, initializeRouters routerFunction) *gin.Engine {
	engine := gin.Default()

	initializeRouters(engine, container)

	return engine
}
