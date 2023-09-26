package entity

import "gorm.io/gorm"

type PetStatus string

const (
	AVAILABLE PetStatus = "available"
	PENDING   PetStatus = "pending"
	SOLD      PetStatus = "sold"
)

type Pet struct {
	gorm.Model
	Category  Category
	Name      string
	PhotoUrls []string
	Tags      []Tag
	Status    PetStatus
}
