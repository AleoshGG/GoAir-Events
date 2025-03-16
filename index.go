package main

import (
	"GoAir-Events/Events/infrastructure"
	"GoAir-Events/Events/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	infrastructure.GoDependences()
	r := gin.Default()
	routes.RegisterRouter(r)
	
	r.Run()
}