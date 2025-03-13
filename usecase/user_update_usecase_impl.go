package usecase

import (
	"errors"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type userUpdateUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserUpdateUseCase(userRepo repository.UserRepository) UserUpdateUseCase {
	return &userUpdateUseCaseImpl{
		userRepository: userRepo,
	}
}

func (uc *userUpdateUseCaseImpl) Execute(id uint, input dto.UserUpdateDTO) (*dto.UserGetDTO, error) {
	// Get existing user
	existingUser, err := uc.userRepository.GetByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Update user fields
	existingUser.FirstName = input.FirstName
	existingUser.LastName = input.LastName
	existingUser.Email = input.Email
	existingUser.Phone = input.Phone
	existingUser.DOB = input.DOB.ToTime()
	existingUser.NIC = input.NIC
	existingUser.Role = input.Role

	// Update existing user fields
	existingUser.FirstName = input.FirstName
	existingUser.LastName = input.LastName
	existingUser.Email = input.Email
	existingUser.Phone = input.Phone
	existingUser.DOB = input.DOB.ToTime()
	existingUser.NIC = input.NIC
	existingUser.Role = input.Role

	// Save updated user
	updatedUser, err := uc.userRepository.Update(*existingUser)
	if err != nil {
		return nil, err
	}

	return &dto.UserGetDTO{
		ID:         updatedUser.ID,
		FirstName:  updatedUser.FirstName,
		LastName:   updatedUser.LastName,
		Email:      updatedUser.Email,
		Phone:      updatedUser.Phone,
		DOB:        updatedUser.DOB,
		NIC:        updatedUser.NIC,
		EmployeeID: updatedUser.EmployeeID,
		Role:       updatedUser.Role,
		CreatedAt:  updatedUser.CreatedAt,
		UpdatedAt:  updatedUser.UpdatedAt,
	}, nil
}
