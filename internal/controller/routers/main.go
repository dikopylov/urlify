package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlify/internal/controller"
	"urlify/internal/model/infrastructure/container"
)

func InitializeRouters(engine *gin.Engine, container *container.Container) {
	healthRouters(engine)

	referenceRouters(engine, controller.NewReferenceController(container.MakeReferenceService()))
}

func healthRouters(engine *gin.Engine) gin.IRoutes {
	return engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

func referenceRouters(engine *gin.Engine, controller controller.ReferenceController) {
	engine.POST("api/url", controller.Create)
	engine.GET("api/:hash", controller.View)
}
