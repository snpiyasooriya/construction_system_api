package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/services"
	"github.com/snpiyasooriya/construction_design_api/usecase"
	"net/http"
	"strconv"
)

type ProjectController struct {
	projectService        *services.ProjectService
	projectsGetUseCase    usecase.ProjectsGetUseCase
	projectGetByIDUseCase usecase.ProjectGetByIDUseCase
}

func NewProjectController(
	projectService *services.ProjectService,
	projectsGetUseCase usecase.ProjectsGetUseCase,
	projectGetByIDUseCase usecase.ProjectGetByIDUseCase,
) *ProjectController {
	return &ProjectController{
		projectService:        projectService,
		projectsGetUseCase:    projectsGetUseCase,
		projectGetByIDUseCase: projectGetByIDUseCase,
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
		return
	}
	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

func (pc *ProjectController) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	project, err := pc.projectGetByIDUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, project)
}

func (pc *ProjectController) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updateDTO dto.ProjectUpdateDTO
	if err := c.ShouldBindJSON(&updateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateDTO.ID = uint(id)

	err = pc.projectService.UpdateProject(updateDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project updated successfully"})
}

func (pc *ProjectController) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = pc.projectService.DeleteProject(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Project deleted successfully"})
}

func (pc *ProjectController) AddUser(c *gin.Context) {
	// Get project ID from URL parameter
	projectID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID format"})
		return
	}

	// Bind user ID from request body
	var addUserDTO dto.ProjectAddUserDTO
	if err := c.ShouldBindJSON(&addUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set project ID from URL parameter
	addUserDTO.ProjectID = uint(projectID)

	result, err := pc.projectService.AddUserToProject(addUserDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
