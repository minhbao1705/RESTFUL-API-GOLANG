package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u *ProductHandler) GetProductsV1(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")

	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (V1)",
		"limit":   limit,
	})
}

func (u *ProductHandler) GetProductsBySlugV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get product by Slug (V1)"})
}

func (u *ProductHandler) PostProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create product (V1)"})
}

func (u *ProductHandler) PutProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update product (V1)"})
}

func (u *ProductHandler) DeleteProductsV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Delete product (V1)"})
}
