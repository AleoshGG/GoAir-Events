package controllers

import (
	"GoAir-Events/Events/application/services"
	"GoAir-Events/Events/infrastructure"

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

}