package main

import (
	"github.com/snpiyasooriya/construction_design_api/config"
	interfaces2 "github.com/snpiyasooriya/construction_design_api/internal/domain/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database/migrations"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/database/repository"
	"github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/controllers"
	server2 "github.com/snpiyasooriya/construction_design_api/internal/infrastructre/http/server"
	"github.com/snpiyasooriya/construction_design_api/internal/interfaces"
	"github.com/snpiyasooriya/construction_design_api/internal/usecases"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgres(conf).GetDb()
	migrations.AutoMigrate(db)
	var userRepo interfaces2.UserRepository
	userRepo = repository.NewGormUserRepository(db)
	// Initialize your use cases
	userCreateUseCase := usecases.NewUserCreateUseCase(userRepo)

	// Initialize your controllers
	userController := controllers.NewUserController(userCreateUseCase)

	var server interfaces.Server
	server = server2.NewGinServer(conf, userController)
	server.Start()

}
