package controllers

import (
	"GoAir-Events/Events/application/services"
	"GoAir-Events/Events/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StartVentilationController struct {
	service *services.StartVentilationService
}

func NewStartVentilationController() *StartVentilationController {
	fcm := infrastructure.GetFCM()
	service := services.NewStartVentilationService(fcm)
	return &StartVentilationController{service: service} 
}

func (svc StartVentilationController) StartVentilation(c *gin.Context) {
	svc.service.Run()
	
	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}