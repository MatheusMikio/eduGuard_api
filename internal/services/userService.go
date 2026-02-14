package services

import "github.com/MatheusMikio/eduGuard_api/internal/repositories"

type IUserService interface{}

type UserService struct {
	Repository repositories.IUserRepository
}

func NewUSerService(repo repositories.IUserRepository) IUserService {
	return &UserService{
		Repository: repo,
	}
}
