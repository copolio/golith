package repository

import (
	"github.com/copolio/golith"
	"github.com/copolio/golith/golithcore"
	"github.com/copolio/golith/golithgorm"
	"github.com/copolio/golith/internal/petstore/entity"
	"go.uber.org/fx"
)

func init() {
	golith.Register(fx.Provide(NewPetRepository))
}

type PetRepository interface {
	golithcore.CrudRepository[entity.Pet, uint64]
}

type PetGormRepository struct {
	golithgorm.Repository[entity.Pet, uint64]
}

func NewPetRepository() PetRepository {
	return &PetGormRepository{}
}
