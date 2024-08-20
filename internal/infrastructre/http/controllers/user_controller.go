package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
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
	var userCreateDTO dto.UserCreateDTO
	if err := c.Bind(&userCreateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the input
	if validationErrors := utils.CustomValidationErrors(userCreateDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	createdUser, err := uc.userCreateUseCase.Execute(userCreateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
