package routes

import (
	"GoAir-Events/Events/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	eventsRouter := r.Group("/events") 
	{
		eventsRouter.POST("/start", controllers.NewStartVentilationController().StartVentilation)
		eventsRouter.POST("/stop", controllers.NewStopVentilationController().StopVentilation)
		eventsRouter.POST("/duration", controllers.NewDurationVentilationController().DurationVentilation)
	}
}