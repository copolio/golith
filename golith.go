package golith

import (
	"github.com/copolio/golith/golithgin"
	"github.com/copolio/golith/golithswag"
	"go.uber.org/fx"
)

var dependencies = []fx.Option{
	golithgin.Use(),
	golithswag.Use(),
	golithgin.Run(),
}

func App(additionalOptions ...fx.Option) (app *fx.App) {
	options := append(additionalOptions, dependencies...)
	return fx.New(
		options...,
	)
}

func Register(options ...fx.Option) {
	dependencies = append(options, dependencies...)
}
