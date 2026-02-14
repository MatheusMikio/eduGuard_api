package repositories

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/entities"
	"github.com/MatheusMikio/eduGuard_api/internal/repositories/base"
	"gorm.io/gorm"
)

type IEventRepository interface {
	base.IBaseRepository[entities.Event]
}

type EventRepository struct {
	base.BaseRepository[entities.Event]
}

func NewEventRepository(db *gorm.DB) IEventRepository {
	return &EventRepository{
		BaseRepository: base.BaseRepository[entities.Event]{
			Db: db,
		},
	}
}
