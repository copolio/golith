package main

import (
	"github.com/copolio/golith"
	_ "github.com/copolio/golith/examples/petstorefx/docs/swagger"
	"github.com/copolio/golith/golithgorm"
	"github.com/copolio/golith/internal/petstore/controller"
	"github.com/copolio/golith/internal/petstore/repository"
	"github.com/copolio/golith/internal/petstore/router"
	"github.com/copolio/golith/internal/petstore/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// @title           Pet Store Example API
// @version         1.0
// @description     This is a sample server.
func main() {
	golith.App(
		golithgorm.Use(),
		fx.Provide(router.NewV2Router),
		fx.Provide(controller.NewPetController),
		fx.Provide(service.NewPetService),
		fx.Provide(repository.NewPetRepository),
		fx.Invoke(func(engine *gin.Engine, router *router.V2Router) {
			router.SetV2Routes(engine)
		}),
	).Run()
}
