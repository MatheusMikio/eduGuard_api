package services

import "github.com/MatheusMikio/eduGuard_api/internal/repositories"

type IEventService interface{}

type EventService struct {
	Repository repositories.IEventRepository
}

func NewEventService(repo repositories.IEventRepository) IEventService {
	return &EventService{
		Repository: repo,
	}
}
