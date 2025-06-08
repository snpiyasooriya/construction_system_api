package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/interfaces/usecase"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"strconv"

	"net/http"
)

type ProjectTypeController struct {
	projectTypeCreateUseCase  usecase.CreateProjectTypeUseCaseInterface
	getAllProjectTypesUseCase usecase.GetAllProjectTypesUseCaseInterface
	getProjectTypeUseCase     usecase.GetProjectTypeUseCaseInterface
	deleteProjectTypeUseCase  usecase.DeleteProjectTypeUseCaseInterface
	updateProjectTypeUseCase  usecase.UpdateProjectTypeUseCaseInterface
}

func NewProjectTypeController(projectTypeCreateUseCase usecase.CreateProjectTypeUseCaseInterface, getAllProjectTypesUseCase usecase.GetAllProjectTypesUseCaseInterface, getProjectTypeUseCase usecase.GetProjectTypeUseCaseInterface, deleteProjectTypeUseCase usecase.DeleteProjectTypeUseCaseInterface, updateProjectTypeUseCase usecase.UpdateProjectTypeUseCaseInterface) *ProjectTypeController {
	return &ProjectTypeController{
		projectTypeCreateUseCase:  projectTypeCreateUseCase,
		getAllProjectTypesUseCase: getAllProjectTypesUseCase,
		getProjectTypeUseCase:     getProjectTypeUseCase,
		deleteProjectTypeUseCase:  deleteProjectTypeUseCase,
		updateProjectTypeUseCase:  updateProjectTypeUseCase,
	}
}

// Create godoc
// @Summary Create a new project type
// @Description Create a new project type with the provided details
// @Tags Project Types
// @Accept json
// @Produce json
// @Param projectType body dto.ProjectTypeCreateDTO true "Project type creation data"
// @Success 201 {object} dto.ProjectTypeCreateDTO "Project type created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project-type [post]
func (pt *ProjectTypeController) Create(c *gin.Context) {
	var createProjectTypeDto dto.ProjectTypeCreateDTO
	if err := c.ShouldBind(&createProjectTypeDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate the input
	if validationErrors := utils.CustomValidationErrors(&createProjectTypeDto); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	err := pt.projectTypeCreateUseCase.Execute(&createProjectTypeDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createProjectTypeDto)
}

// GetAll godoc
// @Summary Get all project types
// @Description Get a list of all project types
// @Tags Project Types
// @Produce json
// @Success 200 {array} dto.ProjectTypeGetDto "List of project types"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project-type [get]
func (pt *ProjectTypeController) GetAll(c *gin.Context) {
	projectTypes, err := pt.getAllProjectTypesUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projectTypes)
}

// Get godoc
// @Summary Get a project type by ID
// @Description Get a project type's details by its ID
// @Tags Project Types
// @Produce json
// @Param id path int true "Project Type ID"
// @Success 200 {object} dto.ProjectTypeGetDto "Project type details"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Project type not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project-type/{id} [get]
func (pt *ProjectTypeController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}
	projectType, err := pt.getProjectTypeUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if projectType == (dto.ProjectTypeGetDto{}) {
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("project type with id %d not found", id)})
		return
	}
	c.JSON(http.StatusOK, projectType)
}

// Delete godoc
// @Summary Delete a project type
// @Description Delete a project type by its ID
// @Tags Project Types
// @Param id path int true "Project Type ID"
// @Success 200 {object} map[string]string "Project type deleted successfully"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project-type/{id} [delete]
func (pt *ProjectTypeController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}
	err = pt.deleteProjectTypeUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "project type with id " + strconv.Itoa(id) + " deleted"})
}

// Update godoc
// @Summary Update a project type
// @Description Update a project type's details by its ID
// @Tags Project Types
// @Accept json
// @Produce json
// @Param id path int true "Project Type ID"
// @Param projectType body dto.ProjectTypeUpdateDTO true "Project type update data"
// @Success 200 {object} dto.ProjectTypeUpdateDTO "Project type updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/project-type/{id} [put]
func (pt *ProjectTypeController) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id must be an integer"})
		return
	}

	var updateProjectTypeDto dto.ProjectTypeUpdateDTO
	if err := c.ShouldBind(&updateProjectTypeDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateProjectTypeDto.ID = uint(id)

	// Validate the input
	if validationErrors := utils.CustomValidationErrors(&updateProjectTypeDto); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	err = pt.updateProjectTypeUseCase.Execute(&updateProjectTypeDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateProjectTypeDto)
}
