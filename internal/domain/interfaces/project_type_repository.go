package interfaces

import "github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database/models"

type ProjectTypeRepository interface {
	Create(input models.ProjectType) (*models.ProjectType, error)
	FindById(input uint) (*models.ProjectType, error)
	Find() (*[]models.ProjectType, error)
}
