package service

import (
	"github.com/copolio/golith/internal/petstore/entity"
	"github.com/copolio/golith/internal/petstore/repository"
	"gorm.io/gorm"
)

type PetUseCase interface {
	AddPet(pet entity.Pet) (entity.Pet, error)
	UpdatePet(pet entity.Pet) (entity.Pet, error)
	FindAllByStatus(status entity.PetStatus) ([]entity.Pet, error)
	FindById(petId uint64) (entity.Pet, error)
	Delete(petId uint64) error
}

type PetService struct {
	db            *gorm.DB
	petRepository repository.PetRepository
}

func NewPetService(db *gorm.DB, petRepository repository.PetRepository) PetUseCase {
	return &PetService{
		db:            db,
		petRepository: petRepository,
	}
}

func (p PetService) AddPet(pet entity.Pet) (entity.Pet, error) {
	err := p.db.Transaction(func(tx *gorm.DB) error {
		_, err := p.petRepository.Save(tx, pet)
		return err
	})
	return pet, err
}

func (p PetService) UpdatePet(pet entity.Pet) (entity.Pet, error) {
	return p.petRepository.Save(p.db, pet)
}

func (p PetService) FindAllByStatus(status entity.PetStatus) ([]entity.Pet, error) {
	return p.petRepository.FindAll(p.db, entity.Pet{Status: status})
}

func (p PetService) FindById(petId uint64) (entity.Pet, error) {
	return p.petRepository.FindById(p.db, petId)
}

func (p PetService) Delete(petId uint64) error {
	return p.db.Transaction(func(tx *gorm.DB) error {
		pet, err := p.petRepository.FindById(p.db, petId)
		if err != nil {
			return err
		}
		dbErr := p.petRepository.Delete(p.db, pet)
		if dbErr != nil {
			return err
		}
		return nil
	})
}
