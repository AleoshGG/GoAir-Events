package controllers

import (
	"GoAir-Events/Events/application/services"
	"GoAir-Events/Events/infrastructure"

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

}