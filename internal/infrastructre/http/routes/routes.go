package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/controllers"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/middlewares"
)

func InitRoutes(router *gin.Engine, userController *controllers.UserController, authenticationController *controllers.AuthenticationController) {
	router.GET("/ping", controllers.Ping)
	router.POST("/login", authenticationController.Login)

	// Apply JWT authentication middleware globally
	secretKey := "ct_sys_api_root"
	router.Use(middlewares.JWTAuthentication(secretKey))
	router.Use(middlewares.CabinMiddleware())
	protectedRoutes := router.Group("/")
	{
		userRoutes := protectedRoutes.Group("/user")
		{
			userRoutes.POST("/", userController.CreateUser)
		}
		_ = protectedRoutes.Group("/authentication")
		{
		}
	}
}
