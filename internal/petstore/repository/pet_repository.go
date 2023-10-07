package repository

import (
	"github.com/copolio/golith/golithgorm"
	"github.com/copolio/golith/internal/petstore/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type PetRepository interface {
	golithgorm.CrudRepository[entity.Pet, uint64]
}

type PetGormRepository struct {
	golithgorm.Repository[entity.Pet, uint64]
}

func NewPetGormRepository() *PetGormRepository {
	return &PetGormRepository{}
}

func (p PetGormRepository) FindById(db *gorm.DB, id uint64) (entity.Pet, error) {
	var pet entity.Pet
	result := db.Preload(clause.Associations).First(&pet, id)
	return pet, result.Error
}
