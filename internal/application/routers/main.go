package routers

import (
	"github.com/gin-gonic/gin"
	"urlify/internal/application/controllers"
	"urlify/internal/infrastructure/container"
)

func InitializeRouters(engine *gin.Engine, container *container.Container) {
	healthRouters(engine)

	referenceRouters(engine, controllers.NewReferenceController(container.MakeReferenceService()))
}

func healthRouters(engine *gin.Engine) gin.IRoutes {
	return engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}

func referenceRouters(engine *gin.Engine, controller controllers.ReferenceController) {
	engine.POST("api/url", controller.Create)
	engine.GET("api/url", controller.View)
}
