package repository

import "hoc-golang/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	FindById(id int, user *models.User) error
}