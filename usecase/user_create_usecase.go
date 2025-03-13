package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/entities"
)

type UserCreateUseCase interface {
	Execute(input dto.UserCreateDTO) (*entities.User, error)
}
