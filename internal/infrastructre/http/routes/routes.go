package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/controllers"
)

func InitRoutes(
	router *gin.Engine,
	userController *controllers.UserController,
	authenticationController *controllers.AuthenticationController,
	projectController *controllers.ProjectController,
	scheduleController *controllers.ScheduleController,
) {
	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/ping", controllers.Ping)
		apiRoutes.POST("/login", authenticationController.Login)

		// Apply JWT authentication middleware globally
		//secretKey := "ct_sys_api_root"
		//apiRoutes.Use(middlewares.JWTAuthentication(secretKey))
		//apiRoutes.Use(middlewares.CabinMiddleware())
		protectedRoutes := apiRoutes.Group("/")
		{
			userRoutes := protectedRoutes.Group("/user")
			{
				userRoutes.POST("/", userController.CreateUser)
			}
			//projectTypeRoutes := protectedRoutes.Group("/project_type")
			//{
			//	projectTypeRoutes.POST("/", projectTypeController.Create)
			//}
			projectRoutes := protectedRoutes.Group("/projects")
			{
				projectRoutes.POST("/", projectController.Create)
				projectRoutes.GET("/", projectController.Get)
			}
			scheduleRoutes := protectedRoutes.Group("/schedules")
			{
				scheduleRoutes.GET("/ByProject/", scheduleController.GetSchedulesByProjectID)
			}
			_ = protectedRoutes.Group("/authentication")
			{
			}
		}
	}
}
