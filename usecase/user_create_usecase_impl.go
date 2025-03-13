package usecase

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/entities"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
)

type userCreateUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserCreateUseCase(userRepo repository.UserRepository) UserCreateUseCase {
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
		ID:        0,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Email:     input.Email,
		Phone:     input.Phone,
		DOB:       input.DOB.ToTime(),
		NIC:       input.NIC,
		Password:  hashedPassword,
		Role:      input.Role,
	}

	// Save the user using the repository
	userEntity, err := uc.userRepository.CreateUser(user)
	if err != nil {
		log.Error("Failed to create user:", err)
		return nil, err
	}

	return userEntity, nil
}
