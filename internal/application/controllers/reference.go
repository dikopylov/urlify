package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	encoding "urlify/internal/domain/reference/services"
)

type ReferenceController struct {
	service encoding.ReferenceService
}

func NewReferenceController(service encoding.ReferenceService) ReferenceController {
	return ReferenceController{service: service}
}

func (controller *ReferenceController) Create(ctx *gin.Context) {
	url := ctx.PostForm("url")

	reference := controller.service.CreateReference(url)

	ctx.JSON(http.StatusCreated, reference)
}

func (controller *ReferenceController) View(ctx *gin.Context) {
	hash := ctx.Param("hash")

	reference := controller.service.GetByHash(hash)

	if reference == nil {
		ctx.JSON(http.StatusNotFound, "Not Found")
	}

	ctx.JSON(http.StatusOK, reference)
}
