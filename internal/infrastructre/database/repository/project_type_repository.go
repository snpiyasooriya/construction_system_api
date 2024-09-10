package repository

import (
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database/models"
	"gorm.io/gorm"
)

type GormProjectTypeRepository struct {
	DB *gorm.DB
}

func NewProjectTypeGORMRepository(db *gorm.DB) *GormProjectTypeRepository {
	return &GormProjectTypeRepository{
		DB: db,
	}
}

func (g *GormProjectTypeRepository) Create(input models.ProjectType) (*models.ProjectType, error) {
	if err := g.DB.Create(&input).Error; err != nil {
		return nil, err
	}
	return &input, nil
}

func (g *GormProjectTypeRepository) FindById(input uint) (*models.ProjectType, error) {
	//TODO implement me
	panic("implement me")
}

func (g *GormProjectTypeRepository) Find() (*[]models.ProjectType, error) {
	//TODO implement me
	panic("implement me")
}
