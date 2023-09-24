package main

import (
	"github.com/copolio/golith"
	_ "github.com/copolio/golith/examples/petstore/docs/swagger"
	"github.com/copolio/golith/golithgin"
	"github.com/copolio/golith/golithgorm"
	"github.com/copolio/golith/golithswag"
	"go.uber.org/fx"
)

// @title           Pet Store Example API
// @version         1.0
// @description     This is a sample server.
func main() {
	golith.App(
		golithgorm.Use(),
		fx.Provide(golithgin.DefaultConfiguration),
		fx.Provide(golithswag.DefaultConfiguration),
	).Run()
}
