package usecase

import (
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/repository"
)

type userGetUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserGetUseCase(userRepo repository.UserRepository) UserGetUseCase {
	return &userGetUseCaseImpl{
		userRepository: userRepo,
	}
}

func (uc *userGetUseCaseImpl) Execute() (*dto.UsersGetDTO, error) {
	users, err := uc.userRepository.Get()
	if err != nil {
		return nil, err
	}

	userDTOs := make([]dto.UserGetDTO, len(users))
	for i, user := range users {
		userDTOs[i] = dto.UserGetDTO{
			ID:         user.ID,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Email:      user.Email,
			Phone:      user.Phone,
			DOB:        user.DOB,
			NIC:        user.NIC,
			EmployeeID: user.EmployeeID,
			Role:       user.Role,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		}
	}

	return &dto.UsersGetDTO{
		Users: userDTOs,
	}, nil
}

type userGetByIDUseCaseImpl struct {
	userRepository repository.UserRepository
}

func NewUserGetByIDUseCase(userRepo repository.UserRepository) UserGetByIDUseCase {
	return &userGetByIDUseCaseImpl{
		userRepository: userRepo,
	}
}

func (uc *userGetByIDUseCaseImpl) Execute(id uint) (*dto.UserGetDTO, error) {
	user, err := uc.userRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserGetDTO{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Phone:      user.Phone,
		DOB:        user.DOB,
		NIC:        user.NIC,
		EmployeeID: user.EmployeeID,
		Role:       user.Role,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}, nil
}
