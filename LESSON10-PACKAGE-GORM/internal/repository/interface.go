package repository

import "hoc-golang/internal/models"

type UserRepository interface {
	Create(user *models.User) error
	FindById(user *models.User, id int) error
}