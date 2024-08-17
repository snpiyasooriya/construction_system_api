package usecases

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/snpiyasooriya/construction_design_api/internal/domain/entities"
	"github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
)

type userCreateUseCaseImpl struct {
	userRepository interfaces.UserRepository
}

func NewUserCreateUseCase(userRepo interfaces.UserRepository) UserCreateUseCase {
	return &userCreateUseCaseImpl{
		userRepository: userRepo,
	}
}

func (uc *userCreateUseCaseImpl) Execute(input dto.UserCreateDTO) (*entities.User, error) {
	// Validate input (you can also use a validation library for this)
	if input.Email == "" || input.Password == "" {
		return nil, errors.New("email and password are required")
	}

	// Hash the password (pseudo-code)
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	// Convert DTO to entity
	user := entities.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Password:  hashedPassword,
	}

	// Save the user using the repository
	userEntity, err := uc.userRepository.CreateUser(user)
	if err != nil {
		log.Error("Failed to create user:", err)
		return nil, err
	}

	return userEntity, nil
}
