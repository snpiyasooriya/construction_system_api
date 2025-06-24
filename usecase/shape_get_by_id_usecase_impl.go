package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/interfaces/usecase"
)

// ShapeGetByIDUseCaseImpl implements the ShapeGetByIDUseCase interface
type ShapeGetByIDUseCaseImpl struct {
	shapeRepo repository.ShapeRepositoryInterFace
}

// NewShapeGetByIDUseCase creates a new instance of ShapeGetByIDUseCaseImpl
func NewShapeGetByIDUseCase(shapeRepo repository.ShapeRepositoryInterFace) usecase.ShapeGetByIDUseCase {
	return &ShapeGetByIDUseCaseImpl{
		shapeRepo: shapeRepo,
	}
}

// Execute retrieves a shape by ID
func (uc *ShapeGetByIDUseCaseImpl) Execute(id uint) (*dto.ShapeDTO, error) {
	// Get shape from repository
	shape, err := uc.shapeRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Convert entity to DTO
	return &dto.ShapeDTO{
		ID:         shape.ID,
		Name:       shape.Name,
		Path:       shape.Path,
		Dimensions: shape.Dimensions,
		Labels:     shape.Labels,
		// Note: CreatedAt and UpdatedAt would be set here if they were available in the entity
	}, nil
}
