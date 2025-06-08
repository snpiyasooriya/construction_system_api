// Package main Construction System API
//
// This is a construction management system API that provides endpoints for managing
// users, projects, schedules, project types, and shapes in a construction environment.
//
// Terms Of Service: http://swagger.io/terms/
//
// Schemes: http, https
// Host: localhost:8080
// BasePath: /api
// Version: 1.0.0
// Contact: Construction System Team <support@construction-system.com>
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// SecurityDefinitions:
// Bearer:
//
//	type: apiKey
//	name: Authorization
//	in: header
//	description: Enter the token with the `Bearer ` prefix, e.g. "Bearer abcde12345"
//
// swagger:meta
package main

import (
	"github.com/snpiyasooriya/construction_design_api/config"
	"github.com/snpiyasooriya/construction_design_api/database"
	"github.com/snpiyasooriya/construction_design_api/database/migrations"
	_ "github.com/snpiyasooriya/construction_design_api/docs"
	"github.com/snpiyasooriya/construction_design_api/http/controllers"
	server2 "github.com/snpiyasooriya/construction_design_api/http/server"
	"github.com/snpiyasooriya/construction_design_api/repository"
	"github.com/snpiyasooriya/construction_design_api/services"
	"github.com/snpiyasooriya/construction_design_api/usecase"
)

func main() {

	conf := config.GetConfig()
	db := database.NewPostgres(conf).GetDb()
	migrations.AutoMigrate(db)
	userRepo := repository.NewGormUserRepository(db)
	userCreateUseCase := usecase.NewUserCreateUseCase(userRepo)
	userGetUseCase := usecase.NewUserGetUseCase(userRepo)
	userGetByIDUseCase := usecase.NewUserGetByIDUseCase(userRepo)
	userUpdateUseCase := usecase.NewUserUpdateUseCase(userRepo)
	userDeleteUseCase := usecase.NewUserDeleteUseCase(userRepo)

	userController := controllers.NewUserController(
		userCreateUseCase,
		userGetUseCase,
		userGetByIDUseCase,
		userUpdateUseCase,
		userDeleteUseCase,
	)

	loginUseCase := usecase.NewLoginUseCaseImpl(userRepo)
	authenticationController := controllers.NewAuthenticationController(loginUseCase)
	projectRepo := repository.NewGORMProjectRepository(db)
	scheduleRepo := repository.NewGORMScheduleRepository(db)
	projectTypeRepo := repository.NewProjectTypeRepository(db)

	projectCreateUseCase := usecase.NewProjectCreateUseCaseImpl(projectRepo)
	projectsGetUseCase := usecase.NewProjectsGetUseCaseImpl(projectRepo, scheduleRepo)
	projectGetByIDUseCase := usecase.NewProjectGetByIDUseCase(projectRepo)
	projectUpdateUseCase := usecase.NewProjectUpdateUseCase(projectRepo)
	projectDeleteUseCase := usecase.NewProjectDeleteUseCase(projectRepo)
	scheduleCreateUseCase := usecase.NewScheduleCreateUseCaseImpl(scheduleRepo, projectRepo)
	scheduleGetByProjectUseCase := usecase.NewScheduleGetByProjectUseCaseImpl(scheduleRepo)
	createProjectTypeUseCase := usecase.NewProjectTypeCreateUseCase(projectTypeRepo)
	getAllProjectTypesUseCase := usecase.NewGetAllProjectTypesUseCase(projectTypeRepo)
	getProjectTypeUsecase := usecase.NewGetProjectTypeUseCase(projectTypeRepo)
	deleteProjectTypeUseCase := usecase.NewDeleteProjectTypeUseCase(projectTypeRepo)
	updateProjectTypeUseCase := usecase.NewProjectTypeUpdateUseCase(projectTypeRepo)

	projectAddUserUseCase := usecase.NewProjectAddUserUseCase(projectRepo)
	projectService := services.NewProjectCreateService(projectCreateUseCase, projectUpdateUseCase, projectDeleteUseCase, scheduleCreateUseCase, projectAddUserUseCase)
	projectController := controllers.NewProjectController(projectService, projectsGetUseCase, projectGetByIDUseCase)
	scheduleController := controllers.NewScheduleController(scheduleGetByProjectUseCase, scheduleCreateUseCase)
	projectTypeController := controllers.NewProjectTypeController(createProjectTypeUseCase, getAllProjectTypesUseCase, getProjectTypeUsecase, deleteProjectTypeUseCase, updateProjectTypeUseCase)

	// Shape components
	shapeRepo := repository.NewGORMShapeRepository(db)
	shapeCreateUseCase := usecase.NewShapeCreateUseCase(shapeRepo)
	shapeGetUseCase := usecase.NewShapeGetUseCase(shapeRepo)
	shapeGetByIDUseCase := usecase.NewShapeGetByIDUseCase(shapeRepo)
	shapeDeleteUseCase := usecase.NewShapeDeleteUseCase(shapeRepo)
	shapeService := services.NewShapeService(shapeCreateUseCase, shapeGetUseCase, shapeGetByIDUseCase, shapeDeleteUseCase)
	shapeController := controllers.NewShapeController(shapeService)

	//var projectTypeRepo interfaces2.ProjectTypeRepository
	//projectTypeRepo = repository.NewProjectTypeGORMRepository(db)
	//projectTypeCreateUseCase := usecase.NewProjectTypeCreateUseCase(projectTypeRepo)
	//projectTypeController := controllers.NewProjectTypeController(projectTypeCreateUseCase)
	server := server2.NewGinServer(conf, userController, authenticationController, projectController, scheduleController, projectTypeController, shapeController)
	server.Start()

}
