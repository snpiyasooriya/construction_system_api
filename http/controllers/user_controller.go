package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"github.com/snpiyasooriya/construction_design_api/usecase"
	"net/http"
	"strconv"
)

type UserController struct {
	userCreateUseCase  usecase.UserCreateUseCase
	userGetUseCase     usecase.UserGetUseCase
	userGetByIDUseCase usecase.UserGetByIDUseCase
	userUpdateUseCase  usecase.UserUpdateUseCase
	userDeleteUseCase  usecase.UserDeleteUseCase
}

func NewUserController(
	userCreateUseCase usecase.UserCreateUseCase,
	userGetUseCase usecase.UserGetUseCase,
	userGetByIDUseCase usecase.UserGetByIDUseCase,
	userUpdateUseCase usecase.UserUpdateUseCase,
	userDeleteUseCase usecase.UserDeleteUseCase,
) *UserController {
	return &UserController{
		userCreateUseCase:  userCreateUseCase,
		userGetUseCase:     userGetUseCase,
		userGetByIDUseCase: userGetByIDUseCase,
		userUpdateUseCase:  userUpdateUseCase,
		userDeleteUseCase:  userDeleteUseCase,
	}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.UserCreateDTO true "User creation data"
// @Success 201 {object} dto.UserGetDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	var userCreateDTO dto.UserCreateDTO
	if err := c.ShouldBindJSON(&userCreateDTO); err != nil {
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

// GetUsers godoc
// @Summary Get all users
// @Description Get a list of all users
// @Tags users
// @Produce json
// @Success 200 {object} dto.UsersGetDTO
// @Failure 500 {object} map[string]interface{}
// @Router /api/users [get]
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.userGetUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByID godoc
// @Summary Get a user by ID
// @Description Get a user's details by their ID
// @Tags users
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} dto.UserGetDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := uc.userGetByIDUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// UpdateUser godoc
// @Summary Update a user
// @Description Update a user's details by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body dto.UserUpdateDTO true "User update data"
// @Success 200 {object} dto.UserGetDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var userUpdateDTO dto.UserUpdateDTO
	if err := c.ShouldBindJSON(&userUpdateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the input
	if validationErrors := utils.CustomValidationErrors(userUpdateDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	updatedUser, err := uc.userUpdateUseCase.Execute(uint(id), userUpdateDTO)
	if err != nil {
		if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

// DeleteUser godoc
// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /api/users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = uc.userDeleteUseCase.Execute(uint(id))
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
