package repositories

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/entities"
	"github.com/MatheusMikio/eduGuard_api/internal/repositories/base"
	"gorm.io/gorm"
)

type IRoomRepository interface {
	base.IBaseRepository[entities.Room]
}

type RoomRepository struct {
	base.BaseRepository[entities.Room]
}

func NewRoomRepository(db *gorm.DB) IRoomRepository {
	return &RoomRepository{
		BaseRepository: base.BaseRepository[entities.Room]{
			Db: db,
		},
	}
}
