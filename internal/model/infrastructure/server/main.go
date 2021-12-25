package server

import (
	"github.com/gin-gonic/gin"
)

type routerFunction func(engine *gin.Engine)

func New(initializeRouters routerFunction) *gin.Engine {
	engine := gin.Default()

	initializeRouters(engine)

	return engine
}
