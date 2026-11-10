package v1handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a number"})
		return
	}
	if id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be positive"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Get user by ID (V1)",
		"user_id": id,
	})
}

func (u *UserHandler) GetUsersByUuIdV1(ctx *gin.Context) {
	uuidStr := ctx.Param("uuid")

	_, err := uuid.Parse(uuidStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID must be a valid UUID"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Get user by ID (V1)",
		"user_uuid": uuidStr,
	})
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
