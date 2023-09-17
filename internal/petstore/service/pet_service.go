package service

import (
	"github.com/copolio/golith"
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

var _ PetUseCase = PetService{}

type PetService struct {
	petRepository *repository.PetRepository
}

func NewPetService(petRepository *repository.PetRepository) *PetService {
	return &PetService{petRepository: petRepository}
}

func (p PetService) AddPet(pet entity.Pet) (entity.Pet, error) {
	err := golith.GetApplication().GormDB.Transaction(func(tx *gorm.DB) error {
		_, err := p.petRepository.Save(tx, pet)
		return err
	})
	return pet, err
}

func (p PetService) UpdatePet(pet entity.Pet) (entity.Pet, error) {
	return p.petRepository.Save(golith.GetApplication().GormDB, pet)
}

func (p PetService) FindAllByStatus(status entity.PetStatus) ([]entity.Pet, error) {
	return p.petRepository.FindAll(golith.GetApplication().GormDB, entity.Pet{Status: status})
}

func (p PetService) FindById(petId uint64) (entity.Pet, error) {
	return p.petRepository.FindById(golith.GetApplication().GormDB, petId)
}

func (p PetService) Delete(petId uint64) error {
	return golith.GetApplication().GormDB.Transaction(func(tx *gorm.DB) error {
		pet, err := p.petRepository.FindById(golith.GetApplication().GormDB, petId)
		if err != nil {
			return err
		}
		dbErr := p.petRepository.Delete(golith.GetApplication().GormDB, pet)
		if dbErr != nil {
			return err
		}
		return nil
	})
}
