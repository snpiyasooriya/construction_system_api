package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type userDeleteUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserDeleteUseCase(userRepo repository.UserRepository) UserDeleteUseCase {
	return &userDeleteUseCaseImpl{
		userRepository: userRepo,
	}
}

func (uc *userDeleteUseCaseImpl) Execute(id uint) error {
	return uc.userRepository.Delete(id)
}
