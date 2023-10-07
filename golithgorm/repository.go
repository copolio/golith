package golithgorm

import (
	"gorm.io/gorm"
)

var _ CrudRepository[any, any] = &Repository[any, any]{}

type CrudRepository[T any, K any] interface {
	Save(db *gorm.DB, entity T) (T, error)
	FindAll(db *gorm.DB, where T) ([]T, error)
	FindById(db *gorm.DB, id K) (T, error)
	Delete(db *gorm.DB, entity T) error
}

type Repository[T any, K any] struct {
}

func (g Repository[T, K]) Save(db *gorm.DB, entity T) (T, error) {
	result := db.Save(&entity)
	return entity, result.Error
}

func (g Repository[T, K]) FindAll(db *gorm.DB, where T) ([]T, error) {
	var entities []T
	result := db.Where(&where).Find(&entities)
	return entities, result.Error
}

func (g Repository[T, K]) FindById(db *gorm.DB, id K) (T, error) {
	var entity T
	result := db.First(&entity, id)
	return entity, result.Error
}

func (g Repository[T, K]) Delete(db *gorm.DB, entity T) error {
	result := db.Delete(&entity)
	return result.Error
}
