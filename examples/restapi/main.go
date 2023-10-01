package main

import (
	"github.com/copolio/golith"
	_ "github.com/copolio/golith/examples/restapi/docs/swagger"
	"github.com/copolio/golith/internal/petstore"
)

// @title           Pet Store Example API
// @version         1.0
// @description     This is a sample server.
func main() {
	app := golith.App(petstore.Module)
	app.Run()
}
