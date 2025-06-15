package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"github.com/snpiyasooriya/construction_design_api/usecase"
)

type ScheduleController struct {
	scheduleGetByProjectUseCase usecase.ScheduleGetByProjectUseCase
	scheduleCreateUseCase       usecase.ScheduleCreateUseCase
	scheduleGetByIDUseCase      usecase.ScheduleGetByIDUseCase
	scheduleUpdateUseCase       usecase.ScheduleUpdateUseCase
}

func NewScheduleController(
	scheduleGetByProjectUseCase usecase.ScheduleGetByProjectUseCase,
	scheduleCreateUseCase usecase.ScheduleCreateUseCase,
	scheduleGetByIDUseCase usecase.ScheduleGetByIDUseCase,
	scheduleUpdateUseCase usecase.ScheduleUpdateUseCase,
) *ScheduleController {
	return &ScheduleController{
		scheduleGetByProjectUseCase: scheduleGetByProjectUseCase,
		scheduleCreateUseCase:       scheduleCreateUseCase,
		scheduleGetByIDUseCase:      scheduleGetByIDUseCase,
		scheduleUpdateUseCase:       scheduleUpdateUseCase,
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
	projectID, err := strconv.ParseUint(c.Param("project_id"), 10, 64)
	fmt.Println(projectID)
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
// @Success 201 {object} dto.ScheduleCreateInputDTO "Schedule created successfully"
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
	err := sc.scheduleCreateUseCase.Execute(&scheduleCreateInputDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, scheduleCreateInputDTO)
}

// GetScheduleByID godoc
// @Summary Get a schedule by ID
// @Description Get a schedule's details by its ID
// @Tags Schedules
// @Produce json
// @Param id path int true "Schedule ID"
// @Success 200 {object} models.Schedule "Schedule details"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 404 {object} map[string]string "Schedule not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/schedule/{id} [get]
func (sc *ScheduleController) GetScheduleByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	// You'll need to implement a scheduleGetByIDUseCase in the controller struct
	schedule, err := sc.scheduleGetByIDUseCase.Execute(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	c.JSON(http.StatusOK, schedule)
}

// UpdateSchedule godoc
// @Summary Update a schedule
// @Description Update a schedule's details by its ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param id path int true "Schedule ID"
// @Param schedule body dto.ScheduleUpdateDTO true "Schedule update data"
// @Success 200 {object} dto.ScheduleUpdateOutputDTO "Schedule updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request"
// @Failure 404 {object} map[string]string "Schedule not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Security Bearer
// @Router /api/schedule/{id} [put]
func (sc *ScheduleController) UpdateSchedule(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	var scheduleUpdateDTO dto.ScheduleUpdateDTO
	if err := c.ShouldBindJSON(&scheduleUpdateDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the ID from the URL parameter
	scheduleUpdateDTO.ID = uint(id)

	// Validate the input
	if validationErrors := utils.CustomValidationErrors(&scheduleUpdateDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	// Execute the update use case
	updatedSchedule, err := sc.scheduleUpdateUseCase.Execute(scheduleUpdateDTO)
	if err != nil {
		if err.Error() == "schedule not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedSchedule)
}
