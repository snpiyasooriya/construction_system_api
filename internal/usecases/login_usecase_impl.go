package usecases

import (
	"errors"
	"github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
)

type LoginUseCaseImpl struct {
	userRepo interfaces.UserRepository
}

func NewLoginUseCaseImpl(userRepo interfaces.UserRepository) LoginUseCase {
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
