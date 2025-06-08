package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type UserCreateUseCase interface {
	Execute(input dto.UserCreateDTO) (*dto.UserGetDTO, error)
}
