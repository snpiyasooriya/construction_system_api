package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/config"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/controllers"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/routes"
)

type GinServer struct {
	router *gin.Engine
	port   int
}

func NewGinServer(conf *config.Config, userController *controllers.UserController, authenticationController *controllers.AuthenticationController) *GinServer {
	router := gin.Default()
	routes.InitRoutes(router, userController, authenticationController)
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
