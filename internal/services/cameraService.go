package services

import "github.com/MatheusMikio/eduGuard_api/internal/repositories"

type ICameraService interface{}

type CameraService struct {
	Repository repositories.ICameraRepository
}

func NewCameraService(repo repositories.ICameraRepository) ICameraService {
	return &CameraService{
		Repository: repo,
	}
}
