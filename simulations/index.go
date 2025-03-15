package main

import (
	"fmt"
	"simulations/utils"
)

func main() {
	var airQuality int

	for {
		fmt.Print("Ingrese la calidad de aire (0-100): ")

		_, err := fmt.Scan(&airQuality)

		if err != nil {
			fmt.Println("Cantidad no aceptable. Ingrese un número válido.")
			fmt.Scanln()
			continue
		}
		if airQuality < 0 || airQuality > 100 {
			fmt.Println("Cantidad no aceptable. Debe estar entre 0 y 100.")
			continue
		}

		if airQuality < 45 {
			fmt.Println("La habitación no necesita ventilación")
		}

		utils.Clean(airQuality)
	}
}
