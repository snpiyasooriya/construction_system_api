package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/snpiyasooriya/construction_design_api/internal/usecases"
	"net/http"
	"strconv"
)

type ScheduleController struct {
	scheduleGetByProjectUseCase usecases.ScheduleGetByProjectUseCase
}

func NewScheduleController(scheduleGetByProjectUseCase usecases.ScheduleGetByProjectUseCase) *ScheduleController {
	return &ScheduleController{
		scheduleGetByProjectUseCase: scheduleGetByProjectUseCase,
	}
}

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
