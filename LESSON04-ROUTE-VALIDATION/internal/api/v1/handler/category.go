package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{"message": "Category found"})
}
