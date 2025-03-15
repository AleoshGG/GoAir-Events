package infrastructure

import "fmt"

type FCM struct {
}

func NewFCM() *FCM {
	return &FCM{}
}

func (e *FCM) StartVentilation() {
	fmt.Println("Se prendió el ventilador")
}
func (e *FCM) StopVentilation()                        {
	fmt.Println("Se apagó el ventilador")
}
func (e *FCM) DurationVentilation(durationSeconds int) {
	fmt.Println("El ventilador trabajó en segundos: %d", durationSeconds)
}