package dto

import (
	"github.com/copolio/golith/examples/petstore/internal/petstore/entity"
)

type CreatePet struct {
	Id       uint `json:"id"`
	Category struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
	Status entity.PetStatus `json:"status"`
}

type UpdatePet struct {
	Id       uint `json:"id"`
	Category struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Name      string   `json:"name"`
	PhotoUrls []string `json:"photoUrls"`
	Tags      []struct {
		Id   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"tags"`
	Status entity.PetStatus `json:"status"`
}
