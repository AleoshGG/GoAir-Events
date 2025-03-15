package controllers

import (
	"GoAir-Events/Events/application/services"
	"GoAir-Events/Events/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StopVentilationController struct {
	service *services.StopVentilationService
}

func NewStopVentilationController() *StopVentilationController {
	fcm := infrastructure.GetFCM()
	service := services.NewStopVentilationService(fcm)
	return &StopVentilationController{service: service} 
}

func (svc StopVentilationController) StopVentilation(c *gin.Context) {
	svc.service.Run()
	
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}