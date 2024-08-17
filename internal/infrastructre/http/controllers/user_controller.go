package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/internal/domain/entities"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	"github.com/snpiyasooriya/construction_design_api/internal/usecases"
	"net/http"
)

type UserController struct {
	userCreateUseCase usecases.UserCreateUseCase
}

func NewUserController(userCreateUseCase usecases.UserCreateUseCase) *UserController {
	return &UserController{userCreateUseCase: userCreateUseCase}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userCreateDTO := dto.UserCreateDTO{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		DOB:       dto.Date{},
		NIC:       user.NIC,
		Password:  user.Password,
		Role:      user.Role,
	}
	createdUser, err := uc.userCreateUseCase.Execute(userCreateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
