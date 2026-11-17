package routes

import (
	"user-management-api/internal/handler"

	"github.com/gin-gonic/gin"
)

type UserRoutes struct {
	handler *handler.UserHandler
}

func NewUserRoutes(handler *handler.UserHandler) *UserRoutes {
	return &UserRoutes{}
}

func (ur *UserRoutes) Register(r *gin.RouterGroup) {
	
}