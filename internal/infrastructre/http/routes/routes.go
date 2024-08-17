package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/controllers"
)

func InitRoutes(router *gin.Engine, userController *controllers.UserController) {
	router.GET("/ping", controllers.Ping)
	RegisterUserRoutes(router, userController)

	// Add more routes as needed
}
