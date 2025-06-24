package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/dto"
	"github.com/snpiyasooriya/construction_design_api/pkg/utils"
	"github.com/snpiyasooriya/construction_design_api/usecase"
)

type ScheduleItemController struct {
	scheduleItemCreateUseCase        usecase.ScheduleItemCreateUseCase
	scheduleItemGetByIDUseCase       usecase.ScheduleItemGetByIDUseCase
	scheduleItemGetByScheduleUseCase usecase.ScheduleItemGetByScheduleUseCase
	scheduleItemUpdateUseCase        usecase.ScheduleItemUpdateUseCase
	scheduleItemDeleteUseCase        usecase.ScheduleItemDeleteUseCase
}

func NewScheduleItemController(
	scheduleItemCreateUseCase usecase.ScheduleItemCreateUseCase,
	scheduleItemGetByIDUseCase usecase.ScheduleItemGetByIDUseCase,
	scheduleItemGetByScheduleUseCase usecase.ScheduleItemGetByScheduleUseCase,
	scheduleItemUpdateUseCase usecase.ScheduleItemUpdateUseCase,
	scheduleItemDeleteUseCase usecase.ScheduleItemDeleteUseCase,
) *ScheduleItemController {
	return &ScheduleItemController{
		scheduleItemCreateUseCase:        scheduleItemCreateUseCase,
		scheduleItemGetByIDUseCase:       scheduleItemGetByIDUseCase,
		scheduleItemGetByScheduleUseCase: scheduleItemGetByScheduleUseCase,
		scheduleItemUpdateUseCase:        scheduleItemUpdateUseCase,
		scheduleItemDeleteUseCase:        scheduleItemDeleteUseCase,
	}
}

// CreateScheduleItem creates a new schedule item
// @Summary Create a new schedule item
// @Description Create a new schedule item with the provided details
// @Tags schedule-items
// @Accept json
// @Produce json
// @Param scheduleItem body dto.ScheduleItemCreateInputDTO true "Schedule Item data"
// @Success 201 {object} dto.ScheduleItemCreateOutputDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/schedule-items [post]
func (sic *ScheduleItemController) CreateScheduleItem(c *gin.Context) {
	var scheduleItemCreateInputDTO dto.ScheduleItemCreateInputDTO

	// Bind JSON request to DTO
	if err := c.ShouldBindJSON(&scheduleItemCreateInputDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the input
	if validationErrors := utils.CustomValidationErrors(&scheduleItemCreateInputDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	// Execute use case
	result, err := sic.scheduleItemCreateUseCase.Execute(scheduleItemCreateInputDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create schedule item: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GetScheduleItemByID retrieves a schedule item by ID
// @Summary Get schedule item by ID
// @Description Get a schedule item by its ID
// @Tags schedule-items
// @Accept json
// @Produce json
// @Param id path int true "Schedule Item ID"
// @Success 200 {object} dto.ScheduleItemGetByIDOutputDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/schedule-items/{id} [get]
func (sic *ScheduleItemController) GetScheduleItemByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	result, err := sic.scheduleItemGetByIDUseCase.Execute(uint(id))
	if err != nil {
		if err.Error() == "schedule item not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get schedule item: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetScheduleItemsByScheduleID retrieves schedule items by schedule ID
// @Summary Get schedule items by schedule ID
// @Description Get all schedule items for a specific schedule
// @Tags schedule-items
// @Accept json
// @Produce json
// @Param scheduleId path int true "Schedule ID"
// @Success 200 {array} dto.ScheduleItemGetByScheduleOutputDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/schedules/{scheduleId}/items [get]
func (sic *ScheduleItemController) GetScheduleItemsByScheduleID(c *gin.Context) {
	scheduleIDParam := c.Param("scheduleId")
	scheduleID, err := strconv.ParseUint(scheduleIDParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID format"})
		return
	}

	result, err := sic.scheduleItemGetByScheduleUseCase.Execute(uint(scheduleID))
	if err != nil {
		if err.Error() == "schedule not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get schedule items: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// UpdateScheduleItem updates an existing schedule item
// @Summary Update schedule item
// @Description Update an existing schedule item with the provided details
// @Tags schedule-items
// @Accept json
// @Produce json
// @Param id path int true "Schedule Item ID"
// @Param scheduleItem body dto.ScheduleItemUpdateInputDTO true "Schedule Item update data"
// @Success 200 {object} dto.ScheduleItemUpdateOutputDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/schedule-items/{id} [put]
func (sic *ScheduleItemController) UpdateScheduleItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var scheduleItemUpdateInputDTO dto.ScheduleItemUpdateInputDTO

	// Bind JSON request to DTO
	if err := c.ShouldBindJSON(&scheduleItemUpdateInputDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set the ID from the URL parameter
	scheduleItemUpdateInputDTO.ID = uint(id)

	// Validate the input
	if validationErrors := utils.CustomValidationErrors(&scheduleItemUpdateInputDTO); validationErrors != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": validationErrors})
		return
	}

	// Execute use case
	result, err := sic.scheduleItemUpdateUseCase.Execute(scheduleItemUpdateInputDTO)
	if err != nil {
		if err.Error() == "schedule item not found" || err.Error() == "shape not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule item: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

// DeleteScheduleItem deletes a schedule item by ID
// @Summary Delete schedule item
// @Description Delete a schedule item by its ID
// @Tags schedule-items
// @Accept json
// @Produce json
// @Param id path int true "Schedule Item ID"
// @Success 200 {object} dto.ScheduleItemDeleteOutputDTO
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Security Bearer
// @Router /api/schedule-items/{id} [delete]
func (sic *ScheduleItemController) DeleteScheduleItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	result, err := sic.scheduleItemDeleteUseCase.Execute(uint(id))
	if err != nil {
		if err.Error() == "schedule item not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete schedule item: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
