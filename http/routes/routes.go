package routes

import (
	"github.com/gin-gonic/gin"
	controllers2 "github.com/snpiyasooriya/construction_design_api/http/controllers"
)

func InitRoutes(
	router *gin.Engine,
	userController *controllers2.UserController,
	authenticationController *controllers2.AuthenticationController,
	projectController *controllers2.ProjectController,
	scheduleController *controllers2.ScheduleController,
	projectTypeController *controllers2.ProjectTypeController,
	shapeController *controllers2.ShapeController,
) {
	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/ping", controllers2.Ping)
		apiRoutes.GET("/health", controllers2.Health)
		apiRoutes.POST("/login", authenticationController.Login)

		// Apply JWT authentication middleware globally
		//secretKey := "ct_sys_api_root"
		//apiRoutes.Use(middlewares.JWTAuthentication(secretKey))
		//apiRoutes.Use(middlewares.CabinMiddleware())
		protectedRoutes := apiRoutes.Group("/")
		{
			userRoutes := protectedRoutes.Group("/users")
			{
				userRoutes.POST("/", userController.CreateUser)
				userRoutes.GET("/", userController.GetUsers)
				userRoutes.GET("/:id", userController.GetUserByID)
				userRoutes.PUT("/:id", userController.UpdateUser)
				userRoutes.DELETE("/:id", userController.DeleteUser)
			}
			projectTypeRoutes := protectedRoutes.Group("/project-type")
			{
				projectTypeRoutes.POST("/", projectTypeController.Create)
				projectTypeRoutes.GET("/", projectTypeController.GetAll)
				projectTypeRoutes.GET("/:id", projectTypeController.Get)
				projectTypeRoutes.DELETE("/:id", projectTypeController.Delete)
				projectTypeRoutes.PUT("/:id", projectTypeController.Update)
			}
			projectRoutes := protectedRoutes.Group("/project")
			{
				projectRoutes.POST("/", projectController.Create)
				projectRoutes.GET("/", projectController.Get)
				projectRoutes.GET("/:id", projectController.GetByID)
				projectRoutes.PUT("/:id", projectController.Update)
				projectRoutes.DELETE("/:id", projectController.Delete)
				projectRoutes.POST("/:id/users", projectController.AddUser)
			}
			scheduleRoutes := protectedRoutes.Group("/schedule")
			{
				scheduleRoutes.GET("/ByProject/", scheduleController.GetSchedulesByProjectID)
			}
			shapeRoutes := protectedRoutes.Group("/shapes")
			{
				shapeRoutes.POST("/", shapeController.Create)
				shapeRoutes.GET("/", shapeController.Get)
				shapeRoutes.GET("/:id", shapeController.GetByID)
				shapeRoutes.DELETE("/:id", shapeController.Delete)
			}
			_ = protectedRoutes.Group("/authentication")
			{
			}
		}
	}
}
