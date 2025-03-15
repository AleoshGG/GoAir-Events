package services

import "GoAir-Events/Events/application"

type StartVentilationService struct {
	ne application.EventsRepository
}

func NewStartVentilationService(ne application.EventsRepository) *StartVentilationService {
	return &StartVentilationService{ne: ne}
}

func (e StartVentilationService) Run() {
	e.ne.StartVentilation()
}