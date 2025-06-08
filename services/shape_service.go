package services

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/usecase"
)

// ShapeService handles shape-related operations
type ShapeService struct {
	shapeCreateUseCase usecase.ShapeCreateUseCase
	shapeGetUseCase    usecase.ShapeGetUseCase
	shapeGetByIDUseCase usecase.ShapeGetByIDUseCase
	shapeDeleteUseCase usecase.ShapeDeleteUseCase
}

// NewShapeService creates a new instance of ShapeService
func NewShapeService(
	shapeCreateUseCase usecase.ShapeCreateUseCase,
	shapeGetUseCase usecase.ShapeGetUseCase,
	shapeGetByIDUseCase usecase.ShapeGetByIDUseCase,
	shapeDeleteUseCase usecase.ShapeDeleteUseCase,
) *ShapeService {
	return &ShapeService{
		shapeCreateUseCase: shapeCreateUseCase,
		shapeGetUseCase:    shapeGetUseCase,
		shapeGetByIDUseCase: shapeGetByIDUseCase,
		shapeDeleteUseCase: shapeDeleteUseCase,
	}
}

// CreateShape creates a new shape
func (s *ShapeService) CreateShape(input *dto.ShapeDTO) error {
	return s.shapeCreateUseCase.Execute(input)
}

// GetShapes retrieves all shapes
func (s *ShapeService) GetShapes() ([]dto.ShapeDTO, error) {
	return s.shapeGetUseCase.Execute()
}

// GetShapeByID retrieves a shape by ID
func (s *ShapeService) GetShapeByID(id uint) (*dto.ShapeDTO, error) {
	return s.shapeGetByIDUseCase.Execute(id)
}

// DeleteShape deletes a shape by ID
func (s *ShapeService) DeleteShape(id uint) error {
	return s.shapeDeleteUseCase.Execute(id)
}
