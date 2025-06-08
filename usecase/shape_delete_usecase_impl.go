package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/interfaces/usecase"
)

// ShapeDeleteUseCaseImpl implements the ShapeDeleteUseCase interface
type ShapeDeleteUseCaseImpl struct {
	shapeRepo repository.ShapeRepositoryInterFace
}

// NewShapeDeleteUseCase creates a new instance of ShapeDeleteUseCaseImpl
func NewShapeDeleteUseCase(shapeRepo repository.ShapeRepositoryInterFace) usecase.ShapeDeleteUseCase {
	return &ShapeDeleteUseCaseImpl{
		shapeRepo: shapeRepo,
	}
}

// Execute deletes a shape by ID
func (uc *ShapeDeleteUseCaseImpl) Execute(id uint) error {
	return uc.shapeRepo.DeleteByID(id)
}
