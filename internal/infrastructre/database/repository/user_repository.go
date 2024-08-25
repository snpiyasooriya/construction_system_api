package repository

import (
	"github.com/snpiyasooriya/construction_design_api/internal/domain/entities"
	"github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database/models"
	"gorm.io/gorm"
)

type GormUserRepository struct {
	db *gorm.DB
}

// NewGormUserRepository creates a new instance of GormUserRepository.
func NewGormUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &GormUserRepository{db: db}
}

func (g *GormUserRepository) CreateUser(user entities.User) (*entities.User, error) {
	gormUser := models.FromUserEntity(user)
	if err := g.db.Create(&gormUser).Error; err != nil {
		return nil, err
	}
	userEntity := gormUser.ToEntity()
	return &userEntity, nil
}

func (g *GormUserRepository) GetUserByID(id int) (*entities.User, error) {
	var gormUser models.User
	if err := g.db.First(&gormUser, id).Error; err != nil {
		return nil, err
	}
	userEntity := gormUser.ToEntity()
	return &userEntity, nil
}

func (g *GormUserRepository) DeleteUserByID(id int) error {
	if err := g.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (g *GormUserRepository) GetAllUsers() ([]entities.User, error) {
	var gormUsers []models.User
	if err := g.db.Find(&gormUsers).Error; err != nil {
		return nil, err
	}
	users := make([]entities.User, len(gormUsers))
	for i, gormUser := range gormUsers {
		users[i] = gormUser.ToEntity()
	}
	return users, nil
}

func (g *GormUserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var gormUser models.User
	if err := g.db.Where("email = ?", email).First(&gormUser).Error; err != nil {
		return nil, err
	}
	userEntity := gormUser.ToEntity()
	return &userEntity, nil
}
