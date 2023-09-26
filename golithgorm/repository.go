package golithgorm

import (
	"github.com/copolio/golith/golithdata"
	"gorm.io/gorm"
)

var _ golithdata.CrudRepository[any, any] = &Repository[any, any]{}

type Repository[T any, K any] struct {
}

func (g Repository[T, K]) Save(db any, entity T) (T, error) {
	tx := db.(*gorm.DB)
	result := tx.Save(&entity)
	return entity, result.Error
}

func (g Repository[T, K]) FindAll(db any, where T) ([]T, error) {
	tx := db.(*gorm.DB)
	var entities []T
	result := tx.Where(&where).Find(&entities)
	return entities, result.Error
}

func (g Repository[T, K]) FindById(db any, id K) (T, error) {
	tx := db.(*gorm.DB)
	var entity T
	result := tx.First(&entity, id)
	return entity, result.Error
}

func (g Repository[T, K]) Delete(db any, entity T) error {
	tx := db.(*gorm.DB)
	result := tx.Delete(&entity)
	return result.Error
}
