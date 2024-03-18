package services

import (
	"api-parking-system/models"
)

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	UpdateUser(*models.User) error
}
