package golith

import (
	"github.com/copolio/golith/golithgin"
	"github.com/copolio/golith/golithswag"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func App(additionalOptions ...fx.Option) (app *fx.App) {
	defaultOptions := []fx.Option{
		golithgin.Use(),
		golithswag.Use(),
		fx.Provide(golithgin.DefaultConfiguration),
		fx.Provide(golithswag.DefaultConfiguration),
		fx.Invoke(func(r *gin.Engine) {
			r.Run()
		}),
	}
	options := append(additionalOptions, defaultOptions...)
	return fx.New(
		options...,
	)
}
