package main

import (
	"github.com/copolio/golith"
	"github.com/copolio/golith/internal/petstore/controller"
	"github.com/copolio/golith/internal/petstore/repository"
	"github.com/copolio/golith/internal/petstore/router"
	"github.com/copolio/golith/internal/petstore/service"
)

func main() {
	application := golith.GetApplication()
	application.Run()
	v2Router := router.NewV2Router(controller.NewPetController(service.NewPetService(&repository.PetRepository{})))
	v2Router.SetV2Routes(application.GinEngine)
}
