package controller

import (
	"errors"
	"fmt"
	"github.com/copolio/golith/examples/petstore/internal/petstore/dto"
	"github.com/copolio/golith/examples/petstore/internal/petstore/entity"
	"github.com/copolio/golith/examples/petstore/internal/petstore/service"
	"github.com/copolio/golith/golithgin"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"time"
)

type PetController struct {
	petUseCase service.PetUseCase
}

func NewPetController(petUseCase service.PetUseCase) *PetController {
	return &PetController{petUseCase: petUseCase}
}

// Create
// @Tags         pet
// @Summary      Adds a new pet to the store
// @Description  Adds a new pet to the store
// @Accept       json
// @Produce      json
// @Param        body body      dto.CreatePet  true  "Pet object that needs to be added to the store"
// @Success      201  {object}  entity.Pet
// @Failure      405  {object}  golithgin.HttpError "Invalid input"
// @Router       /pets [post]
func (p PetController) Create(ctx *gin.Context) {
	request := dto.CreatePet{}
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
	ctx.JSON(http.StatusCreated, response)
}

// Update
// @Tags         pet
// @Summary      Update an existing pet
// @Description  Update an existing pet
// @Accept       json
// @Produce      json
// @Param        body body      dto.UpdatePet  true  "Pet object that needs to be added to the store"
// @Success      200  {object}  entity.Pet
// @Failure      400  {object}  golithgin.HttpError "Invalid ID supplied"
// @Failure      404  {object}  golithgin.HttpError "Pet not found"
// @Failure      405  {object}  golithgin.HttpError "Validation exception"
// @Router       /pets [put]
func (p PetController) Update(ctx *gin.Context) {
	request := dto.UpdatePet{}
	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		ctx.Error(err)
		return
	}
	response, err := p.petUseCase.UpdatePet(request)
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

// Delete
// @Tags         pet
// @Summary      Deletes a pet
// @Description  Deletes a pet
// @Accept       json
// @Produce      json
// @Param        api_key header string  false  "api_key"
// @Param        petId path uint64  true  "Pet id to delete"
// @Failure      400  {object}  golithgin.HttpError "Invalid ID supplied"
// @Failure      404  {object}  golithgin.HttpError "Pet not found"
// @Router       /pets/{petId} [delete]«
func (p PetController) Delete(ctx *gin.Context) {
	petId, err := strconv.ParseUint(ctx.Param("petId"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid ID supplied: %w", err))
		return
	}
	apiKey := ctx.GetHeader("api_key")
	if apiKey == "" {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("invalid Key supplied"))
		return
	}
	err = p.petUseCase.Delete(petId)
	if err != nil {
		ctx.Error(err)
		return
	}
}

// FindById
// @Tags         pet
// @Summary      Find pet by ID
// @Description  Returns a single pet
// @Accept       json
// @Produce      json
// @Param        petId path uint64  true  "ID of pet to return"
// @Success      200  {object}  entity.Pet "successful operation"
// @Failure      400  {object}  golithgin.HttpError "Invalid ID supplied"
// @Failure      404  {object}  golithgin.HttpError "Pet not found"
// @Router       /pets/{petId} [get]
func (p PetController) FindById(ctx *gin.Context) {
	petId, err := strconv.ParseUint(ctx.Param("petId"), 10, 64)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid ID supplied: %w", err))
		return
	}
	pet, err := p.petUseCase.FindById(petId)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("pet not found: %w", err))
		return
	}
	ctx.JSON(http.StatusOK, pet)
}

// FindByStatus
// @Tags         pet
// @Summary      Finds Pets by status
// @Description  Multiple status values can be provided with comma separated strings
// @Accept       json
// @Produce      json
// @Param        status query entity.PetStatus  true  "Status values that need to be considered for filter"
// @Success      200  {object}  []entity.Pet "successful operation"
// @Failure      400  {object}  golithgin.HttpError "Invalid ID supplied"
// @Failure      404  {object}  golithgin.HttpError "Pet not found"
// @Router       /pets/findByStatus [get]
func (p PetController) FindByStatus(ctx *gin.Context) {
	status := entity.PetStatus(ctx.Query("status"))
	results, err := p.petUseCase.FindAllByStatus(status)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid status value: %w", err))
		return
	}
	ctx.JSON(http.StatusOK, results)
}
