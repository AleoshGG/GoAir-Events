package controllers

import (
	"GoAir-Events/Events/application/services"
	"GoAir-Events/Events/infrastructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DurationVentilationController struct {
	service *services.DurationVentilationService
}

func NewDurationVentilationController() *DurationVentilationController {
	fcm := infrastructure.GetFCM()
	service := services.NewDurationVentilationService(fcm)
	return &DurationVentilationController{service: service} 
}

func (svc DurationVentilationController) DurationVentilation(c *gin.Context) {
	duration := c.Param("time")
	durationSeconds, _ := strconv.ParseInt(duration, 10, 64)

	svc.service.Run(int(durationSeconds))
	
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
	
}