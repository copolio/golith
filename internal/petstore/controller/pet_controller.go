package controller

import (
	"errors"
	"github.com/copolio/golith/golithgin"
	"github.com/copolio/golith/internal/petstore/entity"
	"github.com/copolio/golith/internal/petstore/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"time"
)

type PetController struct {
	petUseCase service.PetUseCase
}

func NewPetController(petService *service.PetService) *PetController {
	return &PetController{petUseCase: petService}
}

// Create godoc
// @Summary      Adds a new pet to the store
// @Description  Adds a new pet to the store
// @Accept       json
// @Produce      json
// @Param        body body      entity.Pet  true  "Pet object that needs to be added to the store"
// @Success      200  {object}  entity.Pet
// @Failure      405  {object}  golithgin.HttpError "Invalid input"
// @Router       /v2/pets [post]
func (p PetController) Create() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := entity.Pet{}
		if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
			ctx.Error(err)
			return
		}
		response, err := p.petUseCase.AddPet(request)
		if err != nil {
			ctx.Error(golithgin.HttpError{
				Timestamp: time.Now(),
				Status:    http.StatusMethodNotAllowed,
				Meta:      err,
				Message:   "Invalid input",
			})
			return
		}
		ctx.JSON(http.StatusOK, response)
	}
}

// Update godoc
// @Summary      Update an existing pet
// @Description  Update an existing pet
// @Accept       json
// @Produce      json
// @Param        body body      entity.Pet  true  "Pet object that needs to be added to the store"
// @Success      200  {object}  entity.Pet
// @Failure      400  {object}  golithgin.HttpError "Invalid ID supplied"
// @Failure      404  {object}  golithgin.HttpError "Pet not found"
// @Failure      405  {object}  golithgin.HttpError "Validation exception"
// @Router       /v2/pets [put]
func (p PetController) Update() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		request := entity.Pet{}
		if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
			ctx.Error(err)
			return
		}
	}
}

// Delete godoc
// @Summary      Deletes a pet
// @Description  Deletes a pet
// @Accept       json
// @Produce      json
// @Param        api_key header string  false  "api_key"
// @Param        petId path uint64  true  "Pet id to delete"
// @Failure      400  {object}  golithgin.HttpError "Invalid ID supplied"
// @Failure      404  {object}  golithgin.HttpError "Pet not found"
// @Router       /v2/pets/{petId} [delete]
func (p PetController) Delete() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		petId, err := strconv.ParseUint(ctx.Param("petId"), 10, 64)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid ID supplied"))
			return
		}
		apiKey := ctx.GetHeader("api_key")
		if apiKey == "" {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid Key supplied"))
			return
		}
		err = p.petUseCase.Delete(petId)
		if err != nil {
			ctx.Error(err)
			return
		}
	}
}

// FindById godoc
// @Summary      Find pet by ID
// @Description  Returns a single pet
// @Accept       json
// @Produce      json
// @Param        petId path uint64  true  "ID of pet to return"
// @Success      200  {object}  entity.Pet "successful operation"
// @Failure      400  {object}  golithgin.HttpError "Invalid ID supplied"
// @Failure      404  {object}  golithgin.HttpError "Pet not found"
// @Router       /v2/pets/{petId} [get]
func (p PetController) FindById() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		petId, err := strconv.ParseUint(ctx.Param("petId"), 10, 64)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid ID supplied"))
			return
		}
		pet, err := p.petUseCase.FindById(petId)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("Pet not found"))
			return
		}
		ctx.JSON(http.StatusOK, pet)
	}
}

// FindByStatus godoc
// @Summary      Finds Pets by status
// @Description  Multiple status values can be provided with comma separated strings
// @Accept       json
// @Produce      json
// @Param        status query entity.PetStatus  true  "Status values that need to be considered for filter"
// @Success      200  {object}  []entity.Pet "successful operation"
// @Failure      400  {object}  golithgin.HttpError "Invalid ID supplied"
// @Failure      404  {object}  golithgin.HttpError "Pet not found"
// @Router       /v2/pets/findByStatus [get]
func (p PetController) FindByStatus() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		status := entity.PetStatus(ctx.Query("status"))
		results, err := p.petUseCase.FindAllByStatus(status)
		if err != nil {
			ctx.AbortWithError(http.StatusBadRequest, errors.New("Invalid status value"))
			return
		}
		ctx.JSON(http.StatusOK, results)
	}
}
