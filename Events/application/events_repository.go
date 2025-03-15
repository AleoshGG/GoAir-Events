package application

type EventsRepository interface {
	StartVentilation()
	StopVentilation()
	DurationVentilation()
}