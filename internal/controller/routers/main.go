package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlify/internal/controller"
	"urlify/internal/model/infrastructure/container"
)

func InitializeRouters(engine *gin.Engine) {
	healthRouters(engine)

	referenceRouters(engine)
}

func healthRouters(engine *gin.Engine) gin.IRoutes {
	return engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}

func referenceRouters(engine *gin.Engine) {
	apiHandler := controller.NewReferenceAPIController(container.Get().GetEncoder())
	handler := controller.NewReferenceController(container.Get().GetEncoder())

	api := engine.Group("api")
	{
		api.POST("url", apiHandler.Create)
		api.GET("url/:hash", apiHandler.View)
	}

	engine.GET(":hash", handler.View)
}
