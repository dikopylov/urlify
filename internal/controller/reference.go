package controller

import (
	"database/sql"
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

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	switch reference {
	case nil:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	default:
		ctx.JSON(http.StatusCreated, reference)
	}
}

func (controller *ReferenceController) View(ctx *gin.Context) {
	var viewRequest requests.ViewRequest

	if err := ctx.ShouldBindUri(&viewRequest); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})

		return
	}

	reference, err := controller.encoder.Decode(viewRequest.Hash)

	switch err {
	case sql.ErrNoRows:
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Link Not Found"})
	case nil:
		ctx.JSON(http.StatusOK, reference)
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
