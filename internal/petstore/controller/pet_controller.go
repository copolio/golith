package controller

import (
	"errors"
	"github.com/copolio/golith/internal/petstore/dto"
	"github.com/copolio/golith/internal/petstore/entity"
	"github.com/copolio/golith/internal/petstore/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
)

type PetController struct {
	petService *service.PetService
}

func NewPetController(petService *service.PetService) *PetController {
	return &PetController{petService: petService}
}

func (p PetController) Create() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := entity.Pet{}
		if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
			ctx.Error(err)
			return
		}
		response, err := p.petService.AddPet(request)
		if err != nil {
			ctx.Error(err)
			return
		}
		ctx.JSON(http.StatusOK, response)
	}
}

func (p PetController) Update() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := entity.Pet{}
		if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
			ctx.Error(err)
			return
		}
	}
}

func (p PetController) Delete() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		petId, err := strconv.ParseUint(ctx.Param("petId"), 10, 64)
		if err != nil {
			ctx.AbortWithError(400, errors.New("Invalid ID supplied"))
			return
		}
		apiKey := ctx.GetHeader("api_key")
		if apiKey == "" {
			ctx.AbortWithError(400, errors.New("Invalid Key supplied"))
			return
		}
		err = p.petService.Delete(petId)
		if err != nil {
			ctx.Error(err)
			return
		}
	}
}

func (p PetController) UpdateByForm() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		petId, err := strconv.ParseUint(ctx.Param("petId"), 10, 64)
		if err != nil {
			ctx.AbortWithError(400, errors.New("Invalid ID supplied"))
			return
		}
		var formData dto.UpdatePet
		ctx.Bind(&formData)
		p.petService.UpdatePet(entity.Pet{
			Id:     petId,
			Name:   formData.Name,
			Status: formData.Status,
		})
	}
}

func (p PetController) FindById() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		petId, err := strconv.ParseUint(ctx.Param("petId"), 10, 64)
		if err != nil {
			ctx.AbortWithError(400, errors.New("Invalid ID supplied"))
			return
		}
		pet, err := p.petService.FindById(petId)
		if err != nil {
			ctx.AbortWithError(400, errors.New("Pet not found"))
			return
		}
		ctx.JSON(200, pet)
	}
}

func (p PetController) FindByStatus() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		status := entity.PetStatus(ctx.Query("status"))
		results, err := p.petService.FindAllByStatus(status)
		if err != nil {
			ctx.AbortWithError(400, errors.New("Invalid status value"))
			return
		}
		ctx.JSON(200, results)
	}
}
