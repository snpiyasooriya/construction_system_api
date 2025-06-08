package repository

import (
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/models"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository creates a new instance of GormUserRepository.
func NewGormUserRepository(db *gorm.DB) repository.UserRepository {
	return &GormUserRepository{db: db}
}

func (g *GormUserRepository) Create(user models.User) (*models.User, error) {
	if err := g.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *GormUserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	if err := g.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *GormUserRepository) Update(user models.User) (*models.User, error) {
	if err := g.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (g *GormUserRepository) Delete(id uint) error {
	result := g.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (g *GormUserRepository) Get() ([]models.User, error) {
	var users []models.User
	if err := g.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (g *GormUserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User
	if err := g.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
