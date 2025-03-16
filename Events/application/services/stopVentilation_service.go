package services

import "GoAir-Events/Events/application"

type StopVentilationService struct {
	ne application.EventsRepository
}

func NewStopVentilationService(ne application.EventsRepository) *StopVentilationService {
	return &StopVentilationService{ne: ne}
}

func (e StopVentilationService) Run() {
	e.ne.StopVentilation()
}