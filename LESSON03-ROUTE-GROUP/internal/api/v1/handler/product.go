package v1handler

import (
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (u *ProductHandler) GetProductsV1(ctx *gin.Context) {
	search := ctx.Query("search")

	if search == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Search is required"})
		return
	}

	if len(search) < 3 || len(search) > 50 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Search must be between 3 and 50 characters"})
		return
	}

	if !searchRegex.MatchString(search) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Search much contain only letters, numbers and space"})
		return
	}

	limitsStr := ctx.DefaultQuery("limit", "10")
	limit, err := strconv.Atoi(limitsStr)
	if err != nil || limit <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Limit must be a positive number"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "List all products (V1)",
		"search":  search,
		"limit":   limit,
	})
}

func (u *ProductHandler) GetProductsBySlugV1(ctx *gin.Context) {
	slug := ctx.Param("slug")

	if !slugRegex.MatchString(slug) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Slug must contain only lowercase letter, numbers, hyphens and dots"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get product by Slug (V1)",
		"slug":    slug,
	})
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
