package repository

import (
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database/models"
	"gorm.io/gorm"
)

type GORMProjectRepository struct {
	db *gorm.DB
}

func NewGORMProjectRepository(db *gorm.DB) *GORMProjectRepository {
	return &GORMProjectRepository{
		db: db,
	}
}

func (G *GORMProjectRepository) Create(input dto.ProjectCreateInputDTO) (*dto.ProjectCreateOutputDTO, error) {
	createdProject := &models.Project{
		Model:         gorm.Model{},
		Name:          input.Name,
		ProjectTypeID: input.TypeID,
	}
	if err := G.db.Create(&createdProject).Error; err != nil {
		return nil, err
	}
	projectCreateOutput := dto.ProjectCreateOutputDTO{
		ID:        createdProject.ID,
		Name:      createdProject.Name,
		TypeID:    createdProject.ProjectTypeID,
		CreatedAt: createdProject.CreatedAt,
	}
	return &projectCreateOutput, nil
}

func (G *GORMProjectRepository) UpdateByID(input models.Project) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (G *GORMProjectRepository) DeleteByID(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (G *GORMProjectRepository) GetByID(id uint) (*models.Project, error) {
	//TODO implement me
	panic("implement me")
}

func (G *GORMProjectRepository) Get() (*dto.ProjectsGetDTO, error) {
	var projects []models.Project
	var projectsDTO dto.ProjectsGetDTO

	if err := G.db.Preload("ProjectType").Find(&projects).Error; err != nil {
		return nil, err
	}
	for _, project := range projects {
		projectDTO := dto.ProjectGetDTO{
			ID:          project.ID,
			Name:        project.Name,
			ProjectType: project.ProjectType.Type,
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
		}
		projectsDTO.Projects = append(projectsDTO.Projects, projectDTO)
	}
	return &projectsDTO, nil

}
