package repositories

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
	"github.com/MatheusMikio/eduGuard_api/internal/repositories/base"
	"gorm.io/gorm"
)

type IEventRepository interface {
	base.IBaseRepository[schemas.Event]
}

type EventRepository struct {
	base.BaseRepository[schemas.Event]
}

func NewEventRepository(db *gorm.DB) IEventRepository {
	return &EventRepository{
		BaseRepository: base.BaseRepository[schemas.Event]{
			Db: db,
		},
	}
}
