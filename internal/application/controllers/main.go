package controllers

import (
	"github.com/gin-gonic/gin"
	encoding "urlify/internal/domain/reference/services"
)

type ReferenceController struct {
	service encoding.ReferenceService
}

func NewReferenceController(service encoding.ReferenceService) ReferenceController {
	return ReferenceController{service: service}
}

func (controller *ReferenceController) Create(ctx *gin.Context) {

}

func (controller *ReferenceController) View(ctx *gin.Context) {

}
