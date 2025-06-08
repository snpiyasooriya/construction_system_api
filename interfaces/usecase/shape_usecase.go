package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

// ShapeCreateUseCase defines the interface for creating a shape
type ShapeCreateUseCase interface {
	Execute(input *dto.ShapeDTO) error
}

// ShapeGetUseCase defines the interface for retrieving shapes
type ShapeGetUseCase interface {
	Execute() ([]dto.ShapeDTO, error)
}

// ShapeGetByIDUseCase defines the interface for retrieving a shape by ID
type ShapeGetByIDUseCase interface {
	Execute(id uint) (*dto.ShapeDTO, error)
}

// ShapeDeleteUseCase defines the interface for deleting a shape
type ShapeDeleteUseCase interface {
	Execute(id uint) error
}
