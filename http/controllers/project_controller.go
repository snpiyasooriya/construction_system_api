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

// Create godoc
// @Summary Create a new project
// @Description Create a new project with the provided details
// @Tags Projects
// @Accept json
// @Produce json
// @Param project body dto.ProjectCreateInputDTO true "Project creation data"
// @Success 200 {object} map[string]string "Project created successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project [post]
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

// Get godoc
// @Summary Get all projects
// @Description Get a list of all projects
// @Tags Projects
// @Produce json
// @Success 200 {object} map[string]interface{} "List of projects"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project [get]
func (pc *ProjectController) Get(c *gin.Context) {
	projects, err := pc.projectsGetUseCase.Execute()
	fmt.Println(projects)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"projects get error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// GetByID godoc
// @Summary Get a project by ID
// @Description Get a project's details by its ID
// @Tags Projects
// @Produce json
// @Param id path int true "Project ID"
// @Success 200 {object} dto.ProjectGetDTO "Project details"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project/{id} [get]
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

// Update godoc
// @Summary Update a project
// @Description Update a project's details by its ID
// @Tags Projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param project body dto.ProjectUpdateDTO true "Project update data"
// @Success 200 {object} map[string]string "Project updated successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project/{id} [put]
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

// Delete godoc
// @Summary Delete a project
// @Description Delete a project by its ID
// @Tags Projects
// @Param id path int true "Project ID"
// @Success 200 {object} map[string]string "Project deleted successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project/{id} [delete]
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

// AddUser godoc
// @Summary Add user to project
// @Description Add a user to a specific project
// @Tags Projects
// @Accept json
// @Produce json
// @Param id path int true "Project ID"
// @Param user body dto.ProjectAddUserDTO true "User to add to project"
// @Success 200 {object} map[string]interface{} "User added to project successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project/{id}/users [post]
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
