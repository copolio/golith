package dto

import (
	"github.com/copolio/golith/internal/petstore/entity"
	"mime/multipart"
)

type UploadPetImage struct {
	AdditionalMetaData string                `form:"additionalMetaData"`
	File               *multipart.FileHeader `form:"file"`
}

type UploadPetImageResult struct {
	Code    uint
	Type    string
	Message string
}

type UpdatePet struct {
	Name   string           `form:"name"`
	Status entity.PetStatus `form:"status"`
}
