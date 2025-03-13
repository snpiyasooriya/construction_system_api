package repository

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/models"
)

type ProjectRepository interface {
	Create(input dto.ProjectCreateInputDTO) (*dto.ProjectCreateOutputDTO, error)
	UpdateByID(input models.Project) (*models.Project, error)
	DeleteByID(id uint) error
	GetByID(id uint) (*models.Project, error)
	Get() (*dto.ProjectsGetDTO, error)
}
