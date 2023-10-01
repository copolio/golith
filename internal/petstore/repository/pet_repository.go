package repository

import (
	"github.com/copolio/golith/golithgorm"
	"github.com/copolio/golith/internal/petstore/entity"
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
