package handler

import "user-management-api/internal/service"

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{}
}

func (ush * UserHandler) GetAllUser() {

}