package repositories

import (
	"github.com/MatheusMikio/eduGuard_api/internal/domain/schemas"
	"github.com/MatheusMikio/eduGuard_api/internal/repositories/base"
	"gorm.io/gorm"
)

type IUserRepository interface {
	base.IBaseRepository[schemas.User]
}

type UserRepository struct {
	base.BaseRepository[schemas.User]
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		BaseRepository: base.BaseRepository[schemas.User]{
			Db: db,
		},
	}
}
