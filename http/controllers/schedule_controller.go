package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"github.com/snpiyasooriya/construction_design_api/usecase"
	"net/http"
	"strconv"
)

type ScheduleController struct {
	scheduleGetByProjectUseCase usecase.ScheduleGetByProjectUseCase
	scheduleCreateUseCase       usecase.ScheduleCreateUseCase
}

func NewScheduleController(scheduleGetByProjectUseCase usecase.ScheduleGetByProjectUseCase, scheduleCreateUseCase usecase.ScheduleCreateUseCase) *ScheduleController {
	return &ScheduleController{
		scheduleGetByProjectUseCase: scheduleGetByProjectUseCase,
		scheduleCreateUseCase:       scheduleCreateUseCase,
	}
}

// GetSchedulesByProjectID godoc
// @Summary Get schedules by project ID
// @Description Get all schedules for a specific project
// @Tags Schedules
// @Produce json
// @Param project_id query int true "Project ID"
// @Success 200 {array} dto.ScheduleGetByProjectOutputDTO "List of schedules"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/schedule/ByProject/ [get]
func (sc *ScheduleController) GetSchedulesByProjectID(c *gin.Context) {
	projectID, err := strconv.ParseUint(c.Query("project_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project_id"})
		return
	}
	schedules, err := sc.scheduleGetByProjectUseCase.Execute(uint(projectID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get schedules"})
		return
	}
	c.JSON(http.StatusOK, schedules)
}

// CreateSchedule godoc
// @Summary Create a new schedule
// @Description Create a new schedule with the provided details
// @Tags Schedules
// @Accept json
// @Produce json
// @Param schedule body dto.ScheduleCreateInputDTO true "Schedule creation data"
// @Success 201 {object} dto.ScheduleCreateOutputDTO "Schedule created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/schedule [post]
func (sc *ScheduleController) CreateSchedule(c *gin.Context) {
	var scheduleCreateInputDTO dto.ScheduleCreateInputDTO

	// Bind JSON request to DTO
	if err := c.ShouldBindJSON(&scheduleCreateInputDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the input
	if validationErrors := utils.CustomValidationErrors(&scheduleCreateInputDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	// Execute use case
	createdSchedule, err := sc.scheduleCreateUseCase.Execute(scheduleCreateInputDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdSchedule)
}
