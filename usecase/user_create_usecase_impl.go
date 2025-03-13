package usecase

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
	"github.com/snpiyasooriya/construction_design_api/models"
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

func (uc *userCreateUseCaseImpl) Execute(input dto.UserCreateDTO) (*dto.UserGetDTO, error) {
	// Validate input
	if input.Email == "" || input.Password == "" {
		return nil, errors.New("email and password are required")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	// Create user model
	user := models.User{
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
	createdUser, err := uc.userRepository.Create(user)
	if err != nil {
		log.Error("Failed to create user:", err)
		return nil, err
	}

	// Convert to DTO
	return &dto.UserGetDTO{
		ID:         createdUser.ID,
		FirstName:  createdUser.FirstName,
		LastName:   createdUser.LastName,
		Email:      createdUser.Email,
		Phone:      createdUser.Phone,
		DOB:        createdUser.DOB,
		NIC:        createdUser.NIC,
		EmployeeID: createdUser.EmployeeID,
		Role:       createdUser.Role,
		CreatedAt:  createdUser.CreatedAt,
		UpdatedAt:  createdUser.UpdatedAt,
	}, nil
}
