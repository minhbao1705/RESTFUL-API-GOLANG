package repository

import "hoc-golang/internal/models"

type UserRepository interface {
	Create(user *models.User)
	FindById(id int)
}