package server

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/config"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/controllers"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/routes"
	"net/http"
)

type GinServer struct {
	router *gin.Engine
	port   int
}

func NewGinServer(
	conf *config.Config,
	userController *controllers.UserController,
	authenticationController *controllers.AuthenticationController,
	projectController *controllers.ProjectController,
	scheduleController *controllers.ScheduleController,
) *GinServer {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Set-Cookie"},
		AllowCredentials: true,
	}))
	router.OPTIONS("/*any", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		c.Status(http.StatusNoContent)
	})
	routes.InitRoutes(router, userController, authenticationController, projectController, scheduleController)
	return &GinServer{
		router: router,
		port:   conf.Server.Port,
	}
}

func (server *GinServer) Start() {
	// Start the server on port 8080
	err := server.router.Run(fmt.Sprintf(":%d", server.port))
	if err != nil {
		panic(fmt.Sprintf("gin server start error: %v", err))
	}
}
