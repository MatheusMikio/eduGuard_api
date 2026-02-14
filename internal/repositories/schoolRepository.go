package repositories

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/entities"
	"github.com/MatheusMikio/eduGuard_api/internal/repositories/base"
	"gorm.io/gorm"
)

type ISchoolRepository interface {
	base.IBaseRepository[entities.School]
}

type SchoolRepository struct {
	base.BaseRepository[entities.School]
}

func NewSchoolRepository(db *gorm.DB) ISchoolRepository {
	return &SchoolRepository{
		BaseRepository: base.BaseRepository[entities.School]{
			Db: db,
		},
	}
}
