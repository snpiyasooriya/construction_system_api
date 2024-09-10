package usecases

//
//import (
//	"fmt"
//	log "github.com/sirupsen/logrus"
//	"github.com/snpiyasooriya/construction_design_api/internal/domain/entities"
//	"github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
//	"github.com/snpiyasooriya/construction_design_api/internal/dto"
//)
//
//type ProjectTypeCreateUseCaseImpl struct {
//	projectTypeRepository interfaces.ProjectTypeRepository
//}
//
//func NewProjectTypeCreateUseCase(userRepo interfaces.ProjectTypeRepository) ProjectTypeCreateUseCase {
//	return &ProjectTypeCreateUseCaseImpl{
//		projectTypeRepository: userRepo,
//	}
//}
//
//func (p *ProjectTypeCreateUseCaseImpl) Execute(input dto.ProjectTypeCreateInputDTO) (*dto.ProjectTypeCreateOutputDTO, error) {
//
//	// Convert DTO to entity
//	projectTypeEntity := entities.ProjectType{
//		Type: input.Type,
//	}
//
//	// Save the user using the repository
//	projectTypeEntityCreated, err := p.projectTypeRepository.Create(projectTypeEntity)
//	fmt.Println(projectTypeEntityCreated)
//	if err != nil {
//		log.Error("Failed to create project type:", err)
//		return nil, err
//	}
//
//	projectTypeCreateOutputDTO := dto.ProjectTypeCreateOutputDTO{
//		ID:        projectTypeEntityCreated.ID,
//		Type:      projectTypeEntityCreated.Type,
//		CreatedAt: projectTypeEntityCreated.CreatedAt,
//	}
//
//	return &projectTypeCreateOutputDTO, nil
//}
