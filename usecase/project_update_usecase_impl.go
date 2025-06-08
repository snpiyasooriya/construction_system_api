package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/models"
	"gorm.io/gorm"
)

type projectUpdateUseCaseImpl struct {
	projectRepository repository.ProjectRepository
}

func NewProjectUpdateUseCase(projectRepository repository.ProjectRepository) ProjectUpdateUseCase {
	return &projectUpdateUseCaseImpl{
		projectRepository: projectRepository,
	}
}

func (uc *projectUpdateUseCaseImpl) Execute(input dto.ProjectUpdateDTO) error {
	// Create users slice if user IDs are provided
	var users []models.User
	if len(input.UserIDs) > 0 {
		users = make([]models.User, len(input.UserIDs))
		for i, id := range input.UserIDs {
			users[i] = models.User{Model: gorm.Model{ID: id}}
		}
	}

	// Create project model
	project := models.Project{
		Model:         gorm.Model{ID: input.ID},
		Name:          input.Name,
		ProjectID:     input.ProjectID,
		ProjectTypeID: input.TypeID,
		LeaderID:      input.LeaderID,
		Address:       input.Address,
		StartDate:     input.StartDate,
		EndDate:       input.EndDate,
		Note:          input.Note,
		Status:        input.Status,
		Users:         users,
	}

	_, err := uc.projectRepository.UpdateByID(project)
	return err
}
