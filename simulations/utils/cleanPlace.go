package utils

import (
	"fmt"
	"time"
)

func Clean(airQuality int) {
	FetchAPI("start", 0)
	timeSleep := airQuality / 10
	fmt.Println("Limpiando")
	time.Sleep(time.Duration(timeSleep) *time.Second)
	fmt.Println("Habitación limpia")
	FetchAPI("stop", 0)
}