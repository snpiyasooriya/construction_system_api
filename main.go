package main

import (
	"github.com/snpiyasooriya/construction_design_api/config"
	"github.com/snpiyasooriya/construction_design_api/database"
	"github.com/snpiyasooriya/construction_design_api/database/migrations"
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
	userController := controllers.NewUserController(userCreateUseCase)

	loginUseCase := usecase.NewLoginUseCaseImpl(userRepo)
	authenticationController := controllers.NewAuthenticationController(loginUseCase)
	projectRepo := repository.NewGORMProjectRepository(db)
	scheduleRepo := repository.NewGORMScheduleRepository(db)
	projectTypeRepo := repository.NewProjectTypeRepository(db)

	projectCreateUseCase := usecase.NewProjectCreateUseCaseImpl(projectRepo)
	projectsGetUseCase := usecase.NewProjectsGetUseCaseImpl(projectRepo, scheduleRepo)
	scheduleCreateUseCase := usecase.NewScheduleCreateUseCaseImpl(scheduleRepo)
	scheduleGetByProjectUseCase := usecase.NewScheduleGetByProjectUseCaseImpl(scheduleRepo)
	createProjectTypeUseCase := usecase.NewProjectTypeCreateUseCase(projectTypeRepo)
	getAllProjectTypesUseCase := usecase.NewGetAllProjectTypesUseCase(projectTypeRepo)
	getProjectTypeUsecase := usecase.NewGetProjectTypeUseCase(projectTypeRepo)
	deleteProjectTypeUseCase := usecase.NewDeleteProjectTypeUseCase(projectTypeRepo)
	updateProjectTypeUseCase := usecase.NewProjectTypeUpdateUseCase(projectTypeRepo)

	projectService := services.NewProjectCreateService(projectCreateUseCase, scheduleCreateUseCase)
	projectController := controllers.NewProjectController(projectService, projectsGetUseCase)
	scheduleController := controllers.NewScheduleController(scheduleGetByProjectUseCase)
	projectTypeController := controllers.NewProjectTypeController(createProjectTypeUseCase, getAllProjectTypesUseCase, getProjectTypeUsecase, deleteProjectTypeUseCase, updateProjectTypeUseCase)
	//var projectTypeRepo interfaces2.ProjectTypeRepository
	//projectTypeRepo = repository.NewProjectTypeGORMRepository(db)
	//projectTypeCreateUseCase := usecase.NewProjectTypeCreateUseCase(projectTypeRepo)
	//projectTypeController := controllers.NewProjectTypeController(projectTypeCreateUseCase)
	server := server2.NewGinServer(conf, userController, authenticationController, projectController, scheduleController, projectTypeController)
	server.Start()

}
