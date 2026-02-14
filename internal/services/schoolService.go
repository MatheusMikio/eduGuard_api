package services

import "github.com/MatheusMikio/eduGuard_api/internal/repositories"

type ISchoolService interface{}

type SchoolService struct {
	Repository repositories.ISchoolRepository
}

func NewSchoolService(repo repositories.ISchoolRepository) ISchoolService {
	return &SchoolService{
		Repository: repo,
	}
}
