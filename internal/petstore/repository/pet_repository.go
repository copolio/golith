package repository

import (
	"github.com/copolio/golith/golithgorm"
	"github.com/copolio/golith/internal/petstore/entity"
)

type PetRepository struct {
	golithgorm.Repository[entity.Pet, uint64]
}

func NewPetRepository() *PetRepository {
	return &PetRepository{}
}
