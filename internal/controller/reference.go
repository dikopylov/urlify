package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlify/internal/controller/requests"
	"urlify/internal/model/domain/reference/services"
)

type ReferenceController struct {
	service encoding.ReferenceService
}

func NewReferenceController(service encoding.ReferenceService) ReferenceController {
	return ReferenceController{service: service}
}

func (controller *ReferenceController) Create(ctx *gin.Context) {
	var creationRequest requests.CreationRequest

	if err := ctx.ShouldBindJSON(&creationRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

		return
	}

	reference := controller.service.CreateReference(creationRequest.Url)

	switch reference {
	case nil:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Error"})
	default:
		ctx.JSON(http.StatusCreated, reference)
	}
}

func (controller *ReferenceController) View(ctx *gin.Context) {
	hash := ctx.Param("hash")

	reference := controller.service.GetByHash(hash)

	if reference == nil {
		ctx.JSON(http.StatusNotFound, "Not Found")
	}

	ctx.JSON(http.StatusOK, reference)
}
