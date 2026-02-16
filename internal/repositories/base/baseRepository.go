package base

import "gorm.io/gorm"

type IBaseRepository[T any] interface {
	GetAll(page uint64, size uint64) ([]*T, error)
	GetById(id uint64) (*T, error)
	Create(entity *T) error
	Update(entity *T) error
	Delete(entity *T) error
}

type BaseRepository[T any] struct {
	Db *gorm.DB
}

func (b *BaseRepository[T]) GetAll(page uint64, size uint64) ([]*T, error) {
	var entities []*T
	offSet := int((page - 1) * size)
	if err := b.Db.Offset(offSet).Limit(int(size)).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (b *BaseRepository[T]) GetById(id uint64) (*T, error) {
	var entity *T
	if err := b.Db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (b *BaseRepository[T]) Create(entity *T) error {
	return b.Db.Create(entity).Error
}

func (b *BaseRepository[T]) Update(entity *T) error {
	return b.Db.Save(entity).Error
}

func (b *BaseRepository[T]) Delete(entity *T) error {
	return b.Db.Delete(entity).Error
}
