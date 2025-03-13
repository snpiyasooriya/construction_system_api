package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"github.com/snpiyasooriya/construction_design_api/usecase"
	"net/http"
)

type AuthenticationController struct {
	loginUseCase usecase.LoginUseCase
	//logoutUseCase usecase.LogoutUseCase
}

func NewAuthenticationController(loginUseCase usecase.LoginUseCase) *AuthenticationController {
	return &AuthenticationController{
		loginUseCase: loginUseCase,
		//logoutUseCase: logoutUseCase,
	}
}

func (ac *AuthenticationController) Login(c *gin.Context) {
	var userLoginDTO dto.LoginInputDTO
	if err := c.ShouldBind(&userLoginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(userLoginDTO)

	// Validate the input
	if validationErrors := utils.CustomValidationErrors(&userLoginDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	authentication, err := ac.loginUseCase.Execute(userLoginDTO)
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	domain := "ratko.co"
	secure := true

	if gin.Mode() == gin.DebugMode { // or some other environment check
		domain = "localhost"
		secure = true
	}

	// Set the JWT as a HttpOnly, Secure cookie if in production
	c.SetCookie("jwt", authentication.Token, 3600*24, "/", domain, secure, true) // Secure & HttpOnly
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
