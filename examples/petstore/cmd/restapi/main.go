package main

import (
	"github.com/copolio/golith"
	_ "github.com/copolio/golith/examples/petstore/cmd/restapi/docs/swagger"
	"github.com/copolio/golith/examples/petstore/internal/petstore"
)

// @title           Pet Store Example API
// @version         1.0
// @description     This is a sample server.
// @BasePath /v2
// @tag.name pet
// @tag.description Everything about your Pets
func main() {
	app := golith.App(petstore.Module)
	app.Run()
}
