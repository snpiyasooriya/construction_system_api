package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
)

type LoginUseCase interface {
	Execute(dto dto.LoginInputDTO) (*dto.LoginOutputDTO, error)
}
