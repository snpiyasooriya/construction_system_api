package repository

import "github.com/snpiyasooriya/construction_design_api/entities"

type ShapeRepositoryInterFace interface {
	Create(shape entities.Shape) (*entities.Shape, error)
	GetByID(id uint) (*entities.Shape, error)
	Get() ([]entities.Shape, error)
	DeleteByID(id uint) error
}
