package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	"github.com/snpiyasooriya/construction_design_api/internal/usecases"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"net/http"
)

type AuthenticationController struct {
	loginUseCase usecases.LoginUseCase
	//logoutUseCase usecases.LogoutUseCase
}

func NewAuthenticationController(loginUseCase usecases.LoginUseCase) *AuthenticationController {
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
	// Validate the input
	if validationErrors := utils.CustomValidationErrors(&userLoginDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	authentication, err := ac.loginUseCase.Execute(userLoginDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	domain := "yourdomain.com"
	secure := true

	if gin.Mode() == gin.DebugMode { // or some other environment check
		domain = "localhost"
		secure = true
	}

	// Set the JWT as a HttpOnly, Secure cookie if in production
	c.SetCookie("jwt", authentication.Token, 3600*24, "/", domain, secure, true) // Secure & HttpOnly
	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}
