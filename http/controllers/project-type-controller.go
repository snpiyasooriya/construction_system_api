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

func (pt *ProjectTypeController) GetAll(c *gin.Context) {
	projectTypes, err := pt.getAllProjectTypesUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, projectTypes)
}

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
