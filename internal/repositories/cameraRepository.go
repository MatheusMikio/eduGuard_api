package repositories

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
	"github.com/MatheusMikio/eduGuard_api/internal/repositories/base"
	"gorm.io/gorm"
)

type ICameraRepository interface {
	base.IBaseRepository[schemas.Camera]
}

type CameraRepository struct {
	base.BaseRepository[schemas.Camera]
}

func NewCameraRepository(db *gorm.DB) ICameraRepository {
	return &CameraRepository{
		BaseRepository: base.BaseRepository[schemas.Camera]{
			Db: db,
		},
	}
}
