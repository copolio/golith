package main

import (
	"fmt"
	"github.com/copolio/golith"
	"github.com/copolio/golith/examples/petstore/doc/swagger"
	"github.com/copolio/golith/internal/petstore/controller"
	"github.com/copolio/golith/internal/petstore/entity"
	"github.com/copolio/golith/internal/petstore/repository"
	"github.com/copolio/golith/internal/petstore/router"
	"github.com/copolio/golith/internal/petstore/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	swagger.SwaggerInfo.Title = "Swagger Petstore Copy"
	swagger.SwaggerInfo.Description = "This is a sample server Pet Store server that imitates a Swagger sample."
	swagger.SwaggerInfo.Version = "0.0.1"

	application := golith.NewApplication()
	err := application.GormDB.AutoMigrate(&entity.Pet{})
	if err != nil {
		fmt.Println(err)
	}
	application.GinEngine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	v2Router := router.NewV2Router(controller.NewPetController(service.NewPetService(&repository.PetRepository{})))
	v2Router.SetV2Routes(application.GinEngine)
	application.Run()
}
