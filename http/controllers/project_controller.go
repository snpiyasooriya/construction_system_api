package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/services"
	"github.com/snpiyasooriya/construction_design_api/usecase"
	"net/http"
)

type ProjectController struct {
	projectService     *services.ProjectService
	projectsGetUseCase usecase.ProjectsGetUseCase
}

func NewProjectController(projectService *services.ProjectService, projectsGetUseCase usecase.ProjectsGetUseCase) *ProjectController {
	return &ProjectController{
		projectService:     projectService,
		projectsGetUseCase: projectsGetUseCase,
	}
}

func (pc *ProjectController) Create(c *gin.Context) {
	var projectCreateInputDTO dto.ProjectCreateInputDTO
	err := c.Bind(&projectCreateInputDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"project create error": err.Error()})
	}
	err = pc.projectService.CreateProject(projectCreateInputDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"project create error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project Create Success"})
}

func (pc *ProjectController) Get(c *gin.Context) {
	projects, err := pc.projectsGetUseCase.Execute()
	fmt.Println(projects)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"projects get error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"projects": projects})
}
