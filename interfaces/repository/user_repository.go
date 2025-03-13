package repository

import (
	"github.com/snpiyasooriya/construction_design_api/entities"
)

type UserRepository interface {
	CreateUser(user entities.User) (*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	DeleteUserByID(id int) error
	GetAllUsers() ([]entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
}
