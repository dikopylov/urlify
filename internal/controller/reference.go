package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"urlify/internal/controller/requests"
	"urlify/internal/model/application"
)

type ReferenceController struct {
	encoder application.Encoder
}

func NewReferenceController(encoder application.Encoder) ReferenceController {
	return ReferenceController{encoder: encoder}
}

func (controller *ReferenceController) Create(ctx *gin.Context) {
	var creationRequest requests.CreationRequest

	if err := ctx.ShouldBindJSON(&creationRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

		return
	}

	reference, err := controller.encoder.Encode(creationRequest.Url)

	switch reference {
	case nil:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
	default:
		ctx.JSON(http.StatusCreated, reference)
	}
}

func (controller *ReferenceController) View(ctx *gin.Context) {
	hash := ctx.Param("hash")

	reference, err := controller.encoder.Decode(hash)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	if reference == nil {
		ctx.JSON(http.StatusNotFound, "Not Found")
	}

	ctx.JSON(http.StatusOK, reference)
}
