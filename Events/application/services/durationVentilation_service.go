package services

import (
	"GoAir-Events/Events/application"
)

type DurationVentilationService struct {
	ne application.EventsRepository
}

func NewDurationVentilationService(ne application.EventsRepository) *DurationVentilationService {
	return &DurationVentilationService{ne: ne}
}

func (e DurationVentilationService) Run(durationSeconds int) {
	e.ne.DurationVentilation(durationSeconds)
}