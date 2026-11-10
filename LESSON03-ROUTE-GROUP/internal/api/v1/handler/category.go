package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"minhbao.com/hoc-golang/utils"
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

	if err := utils.ValidationInList("Category", category, validCategory); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message":  "Category found",
		"category": category,
	})
}
