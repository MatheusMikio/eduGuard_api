package repositories

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
	"github.com/MatheusMikio/eduGuard_api/internal/repositories/base"
	"gorm.io/gorm"
)

type ISchoolRepository interface {
	base.IBaseRepository[schemas.School]
}

type SchoolRepository struct {
	base.BaseRepository[schemas.School]
}

func NewSchoolRepository(db *gorm.DB) ISchoolRepository {
	return &SchoolRepository{
		BaseRepository: base.BaseRepository[schemas.School]{
			Db: db,
		},
	}
}
