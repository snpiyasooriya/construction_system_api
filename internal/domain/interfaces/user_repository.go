package interfaces

import "github.com/snpiyasooriya/construction_design_api/internal/domain/entities"

type UserRepository interface {
	CreateUser(user entities.User) (*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	DeleteUserByID(id int) error
	GetAllUsers() ([]entities.User, error)
}
