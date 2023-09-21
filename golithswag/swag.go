package golithswag

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

func Use() fx.Option {
	return fx.Invoke(func(config *Configuration, router *gin.Engine) {
		if config.Use == false {
			return
		}
		router.GET(config.RouterPath, ginSwagger.WrapHandler(swaggerfiles.Handler))
	})
}
