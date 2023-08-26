package gorm

import (
	"github.com/copolio/gin-bootify/pkg/data"
	"gorm.io/gorm"
)

var _ data.CrudRepository[any, any] = &GormRepositoryImpl[any, any]{}

type GormRepository[T any, K any] interface {
	data.CrudRepository[T, K]
}

type GormRepositoryImpl[T any, K any] struct {
	tx *gorm.DB
}

func (g GormRepositoryImpl[T, K]) WithTransaction(transaction any) {
	g.tx = transaction.(*gorm.DB)
}

func (g GormRepositoryImpl[T, K]) Save(entity T) (*T, error) {
	result := g.tx.Save(&entity)
	return &entity, result.Error
}

func (g GormRepositoryImpl[T, K]) FindById(id K) (*T, error) {
	var entity T
	result := g.tx.First(&entity, id)
	return &entity, result.Error
}

func (g GormRepositoryImpl[T, K]) Delete(entity T) (err error) {
	result := g.tx.Delete(&entity)
	return result.Error
}
