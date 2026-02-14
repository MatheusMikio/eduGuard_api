package services

import "github.com/MatheusMikio/eduGuard_api/internal/repositories"

type IRoomService interface{}

type RoomService struct {
	Repository repositories.IRoomRepository
}

func NewRoomService(repo repositories.IRoomRepository) IRoomService {
	return &RoomService{
		Repository: repo,
	}
}
