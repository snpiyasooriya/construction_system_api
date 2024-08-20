package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	httpDTO "github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/dto"
	"github.com/snpiyasooriya/construction_design_api/internal/usecases"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"net/http"
)

type UserController struct {
	userCreateUseCase usecases.UserCreateUseCase
}

func NewUserController(userCreateUseCase usecases.UserCreateUseCase) *UserController {
	return &UserController{userCreateUseCase: userCreateUseCase}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var userCreateHttpDTO httpDTO.UserCreateDTO
	if err := c.Bind(&userCreateHttpDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the input
	if validationErrors := utils.CustomValidationErrors(userCreateHttpDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	userCreateDTO := dto.UserCreateDTO{
		FirstName: userCreateHttpDTO.FirstName,
		LastName:  userCreateHttpDTO.LastName,
		Email:     userCreateHttpDTO.Email,
		Phone:     userCreateHttpDTO.Phone,
		DOB:       userCreateHttpDTO.DOB.ToTime(),
		NIC:       userCreateHttpDTO.NIC,
		Password:  userCreateHttpDTO.Password,
		Role:      userCreateHttpDTO.Role,
	}

	createdUser, err := uc.userCreateUseCase.Execute(userCreateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
