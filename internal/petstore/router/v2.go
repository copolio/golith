package router

import (
	"github.com/copolio/golith/internal/petstore/controller"
	"github.com/gin-gonic/gin"
)

type V2Router struct {
	petController *controller.PetController
}

func NewV2Router(petController *controller.PetController) *V2Router {
	return &V2Router{petController: petController}
}

func (router V2Router) SetV2Routes(r *gin.Engine) {
	v2 := r.Group("/v2")
	{
		pet := v2.Group("/pet")
		{
			pet.POST("", router.petController.Create())
			pet.PUT("", router.petController.Update())
			pet.GET("/findByStatus", router.petController.FindByStatus())
			pet.GET("/:petId", router.petController.FindById())
			pet.DELETE("/:petId", router.petController.Delete())
		}
	}
}
