package repositories

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
	"github.com/MatheusMikio/eduGuard_api/internal/repositories/base"
	"gorm.io/gorm"
)

type IRoomRepository interface {
	base.IBaseRepository[schemas.Room]
}

type RoomRepository struct {
	base.BaseRepository[schemas.Room]
}

func NewRoomRepository(db *gorm.DB) IRoomRepository {
	return &RoomRepository{
		BaseRepository: base.BaseRepository[schemas.Room]{
			Db: db,
		},
	}
}
