package usecase

import (
	"errors"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
)

type LoginUseCaseImpl struct {
	userRepo repository.UserRepository
}

func NewLoginUseCaseImpl(userRepo repository.UserRepository) LoginUseCase {
	return &LoginUseCaseImpl{
		userRepo: userRepo,
	}
}

func (l *LoginUseCaseImpl) Execute(loginInput dto.LoginInputDTO) (*dto.LoginOutputDTO, error) {
	user, err := l.userRepo.GetUserByEmail(loginInput.Email)
	if err != nil {
		return nil, errors.New("user email not found")
	}
	if !utils.CheckPasswordHash(loginInput.Password, user.Password) {
		return nil, errors.New("invalid password")
	}
	token, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	return &dto.LoginOutputDTO{
		UserId:    user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Role:      user.Role,
		Token:     token,
	}, nil

}
