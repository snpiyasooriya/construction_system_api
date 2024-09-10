package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/internal/dto"
	"github.com/snpiyasooriya/construction_design_api/internal/usecases"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"net/http"
)

type ProjectTypeController struct {
	projectTypeCreateUseCase usecases.ProjectTypeCreateUseCase
}

func NewProjectTypeController(projectTypeCreateUseCase usecases.ProjectTypeCreateUseCase) *ProjectTypeController {
	return &ProjectTypeController{
		projectTypeCreateUseCase: projectTypeCreateUseCase,
	}
}

func (pt *ProjectTypeController) Create(c *gin.Context) {
	var projectTypeCreateIntputDTO dto.ProjectTypeCreateInputDTO
	if err := c.ShouldBind(&projectTypeCreateIntputDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validate the input
	if validationErrors := utils.CustomValidationErrors(&projectTypeCreateIntputDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	createdProject, err := pt.projectTypeCreateUseCase.Execute(projectTypeCreateIntputDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdProject)
}
