package gorm

import (
	"github.com/copolio/gin-bootify/pkg/data"
	"github.com/copolio/gin-bootify/pkg/web"
	"gorm.io/gorm"
)

var _ data.FpCrudRepository[any, any] = &GormRepository[any, any]{}
var _ data.CrudRepository[any, any] = &GormRepository[any, any]{}

type GormRepository[T any, K any] struct {
}

func (g GormRepository[T, K]) FpSave(entity T) func(arg any) (T, error) {
	return func(arg any) (T, error) {
		tx := arg.(*gorm.DB)
		result := tx.Save(&entity)
		return entity, result.Error
	}
}

func (g GormRepository[T, K]) FpFindById(id K) func(arg any) (T, error) {
	return func(arg any) (T, error) {
		tx := arg.(*gorm.DB)
		var entity T
		result := tx.First(&entity, id)
		return entity, result.Error
	}
}

func (g GormRepository[T, K]) FpDelete(entity T) func(arg any) error {
	return func(arg any) error {
		tx := arg.(*gorm.DB)
		result := tx.Delete(&entity)
		return result.Error
	}
}

func (g GormRepository[T, K]) WithTransaction(transaction any) {
	web.GetApplication().GormDB = transaction.(*gorm.DB)
}

func (g GormRepository[T, K]) Save(entity T) (*T, error) {
	result := web.GetApplication().GormDB.Save(&entity)
	return &entity, result.Error
}

func (g GormRepository[T, K]) FindById(id K) (*T, error) {
	var entity T
	result := web.GetApplication().GormDB.First(&entity, id)
	return &entity, result.Error
}

func (g GormRepository[T, K]) Delete(entity T) (err error) {
	result := web.GetApplication().GormDB.Delete(&entity)
	return result.Error
}
