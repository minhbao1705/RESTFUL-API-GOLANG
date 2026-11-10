package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
}

var validCategory = map[string]bool{
	"php":    true,
	"golang": true,
	"python": true,
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (c *CategoryHandler) GetCategoryByCategoryV1(ctx *gin.Context) {
	category := ctx.Param("category")

	if !validCategory[category] {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Category must be one of: php, python, golang"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": category,
	})
}
