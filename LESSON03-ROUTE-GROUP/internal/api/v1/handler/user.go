package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) GetUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "List all users (V1)"})
}

func (u *UserHandler) GetUsersByIdV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Get user by ID (V1)"})
}

func (u *UserHandler) PostUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{"message": "Create user (V1)"})
}

func (u *UserHandler) PutUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Update user (V1)"})
}

func (u *UserHandler) DeleteUsersV1(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{"message": "Delete user (V1)"})
}
