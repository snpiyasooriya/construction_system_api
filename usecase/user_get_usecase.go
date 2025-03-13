package usecase

import "github.com/snpiyasooriya/construction_design_api/dto"

type UserGetUseCase interface {
	Execute() (*dto.UsersGetDTO, error)
}

type UserGetByIDUseCase interface {
	Execute(id uint) (*dto.UserGetDTO, error)
}
