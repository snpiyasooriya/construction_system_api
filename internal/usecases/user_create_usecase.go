package usecases

import (
	"github.com/snpiyasooriya/construction_design_api/internal/domain/entities"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
)

type UserCreateUseCase interface {
	Execute(input dto.UserCreateDTO) (*entities.User, error)
}
