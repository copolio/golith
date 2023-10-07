package entity

import "gorm.io/gorm"

type PhotoUrl struct {
	gorm.Model
	PetID uint
	Pet   Pet
	Url   string
}
