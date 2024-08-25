package usecases

import "github.com/snpiyasooriya/construction_design_api/internal/dto"

type LoginUseCase interface {
	Execute(dto dto.LoginInputDTO) (*dto.LoginOutputDTO, error)
}
