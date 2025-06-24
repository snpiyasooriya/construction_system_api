package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/entities"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/interfaces/usecase"
)

// ShapeCreateUseCaseImpl implements the ShapeCreateUseCase interface
type ShapeCreateUseCaseImpl struct {
	shapeRepo repository.ShapeRepositoryInterFace
}

// NewShapeCreateUseCase creates a new instance of ShapeCreateUseCaseImpl
func NewShapeCreateUseCase(shapeRepo repository.ShapeRepositoryInterFace) usecase.ShapeCreateUseCase {
	return &ShapeCreateUseCaseImpl{
		shapeRepo: shapeRepo,
	}
}

// Execute creates a new shape
func (uc *ShapeCreateUseCaseImpl) Execute(input *dto.ShapeDTO) error {
	// Convert DTO to entity
	shapeEntity := entities.Shape{
		Name:       input.Name,
		Path:       input.Path,
		Dimensions: input.Dimensions,
		Labels:     input.Labels,
	}

	// Create shape
	createdShape, err := uc.shapeRepo.Create(shapeEntity)
	if err != nil {
		return err
	}

	// Update the input DTO with the created shape's data
	input.ID = createdShape.ID
	input.Name = createdShape.Name
	input.Path = createdShape.Path
	input.Dimensions = createdShape.Dimensions
	input.Labels = createdShape.Labels

	// Note: CreatedAt and UpdatedAt would be set here if they were available in the entity

	return nil
}
