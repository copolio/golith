package golith

import (
	"go.uber.org/fx"
)

func App(options ...fx.Option) (app *fx.App) {
	return fx.New(
		options...,
	)
}
