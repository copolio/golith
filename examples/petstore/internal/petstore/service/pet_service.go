package service

import (
	"github.com/copolio/golith/examples/petstore/internal/petstore/dto"
	entity2 "github.com/copolio/golith/examples/petstore/internal/petstore/entity"
	"github.com/copolio/golith/examples/petstore/internal/petstore/repository"
	"gorm.io/gorm"
)

type PetUseCase interface {
	AddPet(request dto.CreatePet) (entity2.Pet, error)
	UpdatePet(request dto.UpdatePet) (entity2.Pet, error)
	FindAllByStatus(status entity2.PetStatus) ([]entity2.Pet, error)
	FindById(petId uint64) (entity2.Pet, error)
	Delete(petId uint64) error
}

type PetService struct {
	db            *gorm.DB
	petRepository repository.PetRepository
}

func NewPetService(db *gorm.DB, petRepository repository.PetRepository) *PetService {
	return &PetService{
		db:            db,
		petRepository: petRepository,
	}
}

func (p PetService) AddPet(request dto.CreatePet) (entity2.Pet, error) {
	tags := make([]entity2.Tag, len(request.Tags))
	photoUrls := make([]entity2.PhotoUrl, len(request.PhotoUrls))
	for idx, tag := range request.Tags {
		tags[idx] = entity2.Tag{
			Model: gorm.Model{
				ID: tag.Id,
			},
			Name: tag.Name,
		}
	}
	for idx, photoUrl := range request.PhotoUrls {
		photoUrls[idx] = entity2.PhotoUrl{
			Url: photoUrl,
		}
	}
	pet := entity2.Pet{
		Category: entity2.Category{
			Model: gorm.Model{
				ID: request.Category.Id,
			},
			Name: request.Category.Name,
		},
		Name:      request.Name,
		PhotoUrls: photoUrls,
		Tags:      tags,
		Status:    request.Status,
	}
	err := p.db.Transaction(func(tx *gorm.DB) error {
		_, err := p.petRepository.Save(tx, pet)
		return err
	})
	return pet, err
}

func (p PetService) UpdatePet(request dto.UpdatePet) (entity2.Pet, error) {
	tags := make([]entity2.Tag, len(request.Tags))
	photoUrls := make([]entity2.PhotoUrl, len(request.PhotoUrls))
	for idx, tag := range request.Tags {
		tags[idx] = entity2.Tag{
			Model: gorm.Model{
				ID: tag.Id,
			},
			Name: tag.Name,
		}
	}
	for idx, photoUrl := range request.PhotoUrls {
		photoUrls[idx] = entity2.PhotoUrl{
			Url: photoUrl,
		}
	}
	pet := entity2.Pet{
		Category: entity2.Category{
			Model: gorm.Model{
				ID: request.Category.Id,
			},
			Name: request.Category.Name,
		},
		Name:      request.Name,
		PhotoUrls: photoUrls,
		Tags:      tags,
		Status:    request.Status,
	}
	err := p.db.Transaction(func(tx *gorm.DB) error {
		_, err := p.petRepository.Save(tx, pet)
		return err
	})
	return pet, err
}

func (p PetService) FindAllByStatus(status entity2.PetStatus) ([]entity2.Pet, error) {
	return p.petRepository.FindAll(p.db, entity2.Pet{Status: status})
}

func (p PetService) FindById(petId uint64) (entity2.Pet, error) {
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
