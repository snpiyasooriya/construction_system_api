package repository

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type ProjectTypeRepositoryInterface interface {
	Create(input *dto.ProjectTypeCreateDTO) error
	FindById(id uint) (dto.ProjectTypeGetDto, error)
	Find() ([]dto.ProjectTypeGetDto, error)
	DeleteById(id uint) error
	Update(input *dto.ProjectTypeUpdateDTO) error
}
