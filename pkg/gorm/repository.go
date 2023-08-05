package gorm

import (
	"github.com/copolio/gin-bootify/pkg/database"
	"gorm.io/gorm"
)

var _ database.CrudRepository[any, any] = &GormRepository[any, any]{}

type GormRepository[T any, K any] struct {
	tx *gorm.DB
}

func (g GormRepository[T, K]) WithTransaction(transaction any) {
	g.tx = transaction.(*gorm.DB)
}

func (g GormRepository[T, K]) Save(entity T) (*T, error) {
	result := g.tx.Save(&entity)
	return &entity, result.Error
}

func (g GormRepository[T, K]) FindById(id K) (*T, error) {
	var entity T
	result := g.tx.First(&entity, id)
	return &entity, result.Error
}

func (g GormRepository[T, K]) Delete(entity T) (err error) {
	result := g.tx.Delete(&entity)
	return result.Error
}
