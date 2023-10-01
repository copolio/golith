package petstore

import (
	"github.com/copolio/golith/golithgin"
	"github.com/copolio/golith/golithgorm"
	"github.com/copolio/golith/golithswag"
	"github.com/copolio/golith/golithviper"
	"github.com/copolio/golith/internal/petstore/controller"
	"github.com/copolio/golith/internal/petstore/repository"
	"github.com/copolio/golith/internal/petstore/router"
	"github.com/copolio/golith/internal/petstore/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("restapi",
	golithgin.Use(),
	golithswag.Use(),
	golithgorm.Use(),
	golithviper.Use(),
	fx.Provide(
		router.NewV2Router,
		controller.NewPetController,
		fx.Annotate(
			service.NewPetService,
			fx.As(new(service.PetUseCase)),
		),
		fx.Annotate(
			repository.NewPetGormRepository,
			fx.As(new(repository.PetRepository)),
		),
	),
	fx.Invoke(func(engine *gin.Engine, router *router.V2Router) {
		router.SetRoutes(engine)
	}),
)
