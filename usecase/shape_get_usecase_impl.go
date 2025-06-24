package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/interfaces/usecase"
)

// ShapeGetUseCaseImpl implements the ShapeGetUseCase interface
type ShapeGetUseCaseImpl struct {
	shapeRepo repository.ShapeRepositoryInterFace
}

// NewShapeGetUseCase creates a new instance of ShapeGetUseCaseImpl
func NewShapeGetUseCase(shapeRepo repository.ShapeRepositoryInterFace) usecase.ShapeGetUseCase {
	return &ShapeGetUseCaseImpl{
		shapeRepo: shapeRepo,
	}
}

// Execute retrieves all shapes
func (uc *ShapeGetUseCaseImpl) Execute() ([]dto.ShapeDTO, error) {
	// Get shapes from repository
	shapes, err := uc.shapeRepo.Get()
	if err != nil {
		return nil, err
	}

	// Convert entities to DTOs
	shapeDTOs := make([]dto.ShapeDTO, len(shapes))
	for i, shape := range shapes {
		shapeDTOs[i] = dto.ShapeDTO{
			ID:         shape.ID,
			Name:       shape.Name,
			Path:       shape.Path,
			Dimensions: shape.Dimensions,
			Labels:     shape.Labels,
			// Note: CreatedAt and UpdatedAt would be set here if they were available in the entity
		}
	}

	return shapeDTOs, nil
}
