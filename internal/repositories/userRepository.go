package repositories

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/entities"
	"github.com/MatheusMikio/eduGuard_api/internal/repositories/base"
	"gorm.io/gorm"
)

type IUserRepository interface {
	base.IBaseRepository[entities.User]
}

type UserRepository struct {
	base.BaseRepository[entities.User]
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		BaseRepository: base.BaseRepository[entities.User]{
			Db: db,
		},
	}
}
