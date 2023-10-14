package entity

import "gorm.io/gorm"

type PetStatus string

func (p PetStatus) GormDataType() string {
	return "text"
}

const (
	AVAILABLE PetStatus = "available"
	PENDING   PetStatus = "pending"
	SOLD      PetStatus = "sold"
)

type Pet struct {
	gorm.Model
	CategoryID uint
	Category   Category
	Name       string
	PhotoUrls  []PhotoUrl
	Tags       []Tag
	Status     PetStatus
}
