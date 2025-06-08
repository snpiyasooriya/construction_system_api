package repository

import (
	"github.com/snpiyasooriya/construction_design_api/models"
)

type UserRepository interface {
	Create(user models.User) (*models.User, error)
	GetByID(id uint) (*models.User, error)
	Update(user models.User) (*models.User, error)
	Delete(id uint) error
	Get() ([]models.User, error)
	GetByEmail(email string) (*models.User, error)
}
