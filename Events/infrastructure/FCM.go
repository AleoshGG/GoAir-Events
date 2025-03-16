package infrastructure

import (
	"context"
	"fmt"
	"log"
	"strconv"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type FCM struct {
	app *firebase.App
}

func NewFCM() *FCM {
	// Ruta al archivo de credenciales JSON
	opt := option.WithCredentialsFile("C:\\Users\\spart\\OneDrive\\Documentos\\Ingenieria en Desarrollo de Software\\5to Cuatrimestre\\Programación Cliente Servidor\\GoAir-Events\\auth.json")
	
	// Inicializa la aplicación Firebase
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error al inicializar Firebase: %v", err)
	}

	return &FCM{app: app}
}

func (e *FCM) StartVentilation() {
	fmt.Println("Se prendió el ventilador")
	e.sendMessage("Sistema de ventilación", "Se prendió el ventilador")
}
func (e *FCM) StopVentilation()                        {
	fmt.Println("Se apagó el ventilador")
	e.sendMessage("Sistema de ventilación", "Se apagó el ventilador")
}
func (e *FCM) DurationVentilation(durationSeconds int) {
	fmt.Println("El ventilador trabajó en segundos: ", durationSeconds)
	dString := strconv.Itoa(durationSeconds)
	e.sendMessage("Sistema de ventilación", "El ventilador trabajó: "+dString+" segundos")
}

func (e *FCM) sendMessage(title string, body string) {
	// Crea el cliente de Messaging
	ctx := context.Background()
	client, err := e.app.Messaging(ctx)
	if err != nil {
		log.Fatalf("Error al obtener el cliente de Messaging: %v", err)
	}

	// Define el mensaje que deseas enviar
	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Topic: "GoAirNotications", // Reemplaza con el token del dispositivo o usa Topics
	}

	// Envía el mensaje
	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalf("Error al enviar el mensaje: %v", err)
	}

	log.Printf("Mensaje enviado correctamente: %s", response)
}