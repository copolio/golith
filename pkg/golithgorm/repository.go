package golithgorm

import (
	"github.com/copolio/golith/pkg/golithdata"
	"gorm.io/gorm"
)

var _ golithdata.FpCrudRepository[any, any] = &Repository[any, any]{}
var _ golithdata.CrudRepository[any, any] = &Repository[any, any]{}

type Repository[T any, K any] struct {
	db *gorm.DB
}

func (g Repository[T, K]) FpSave(entity T) func(arg any) (T, error) {
	return func(arg any) (T, error) {
		tx := arg.(*gorm.DB)
		result := tx.Save(&entity)
		return entity, result.Error
	}
}

func (g Repository[T, K]) FpFindById(id K) func(arg any) (T, error) {
	return func(arg any) (T, error) {
		tx := arg.(*gorm.DB)
		var entity T
		result := tx.First(&entity, id)
		return entity, result.Error
	}
}

func (g Repository[T, K]) FpDelete(entity T) func(arg any) error {
	return func(arg any) error {
		tx := arg.(*gorm.DB)
		result := tx.Delete(&entity)
		return result.Error
	}
}

func (g Repository[T, K]) WithTransaction(transaction any) {
	g.db = transaction.(*gorm.DB)
}

func (g Repository[T, K]) Save(entity T) (*T, error) {
	result := g.db.Save(&entity)
	return &entity, result.Error
}

func (g Repository[T, K]) FindById(id K) (*T, error) {
	var entity T
	result := g.db.First(&entity, id)
	return &entity, result.Error
}

func (g Repository[T, K]) Delete(entity T) (err error) {
	result := g.db.Delete(&entity)
	return result.Error
}
