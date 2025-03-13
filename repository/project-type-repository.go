package repository

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/models"
	"gorm.io/gorm"
)

type ProjectTypeRepository struct {
	DB *gorm.DB
}

func NewProjectTypeRepository(db *gorm.DB) *ProjectTypeRepository {
	return &ProjectTypeRepository{
		DB: db,
	}
}

func (g *ProjectTypeRepository) Create(input *dto.ProjectTypeCreateDTO) error {
	projectTypeModel := models.ProjectType{
		Type: input.Type,
	}
	if err := g.DB.Create(&projectTypeModel).Error; err != nil {
		return err
	}
	input.CreatedAt = projectTypeModel.CreatedAt
	input.ID = projectTypeModel.ID
	input.Type = projectTypeModel.Type
	return nil
}

func (g *ProjectTypeRepository) FindById(input uint) (dto.ProjectTypeGetDto, error) {
	projectType := models.ProjectType{}
	if err := g.DB.Where("id = ?", input).First(&projectType).Error; err != nil {
		return dto.ProjectTypeGetDto{}, err
	}

	projectTypeDto := dto.ProjectTypeGetDto{
		ID:        projectType.ID,
		Type:      projectType.Type,
		CreatedAt: projectType.CreatedAt,
		UpdatedAt: projectType.UpdatedAt,
	}

	return projectTypeDto, nil
}

func (g *ProjectTypeRepository) Find() ([]dto.ProjectTypeGetDto, error) {
	projectTypes := []models.ProjectType{}
	if err := g.DB.Find(&projectTypes).Error; err != nil {
		return nil, err
	}
	projectTypeDtos := make([]dto.ProjectTypeGetDto, len(projectTypes))
	for i, projectType := range projectTypes {
		projectTypeDtos[i] = dto.ProjectTypeGetDto{
			ID:        projectType.ID,
			Type:      projectType.Type,
			CreatedAt: projectType.CreatedAt,
			UpdatedAt: projectType.UpdatedAt,
		}
	}
	return projectTypeDtos, nil
}

func (g *ProjectTypeRepository) DeleteById(id uint) error {
	if err := g.DB.Where("id = ?", id).Delete(&models.ProjectType{}).Error; err != nil {
		return err
	}
	return nil
}

func (g *ProjectTypeRepository) Update(input *dto.ProjectTypeUpdateDTO) error {
	projectType := models.ProjectType{}
	if err := g.DB.First(&projectType, input.ID).Error; err != nil {
		return err
	}

	projectType.Type = input.Type
	if err := g.DB.Save(&projectType).Error; err != nil {
		return err
	}

	input.UpdatedAt = projectType.UpdatedAt
	return nil
}
