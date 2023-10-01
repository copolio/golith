package golithgin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

func NewGin(lc fx.Lifecycle, conf *Configuration) *gin.Engine {
	gin.SetMode(string(conf.Mode))
	router := gin.Default()
	router.Use(HttpErrorHandler())

	server := &http.Server{Addr: fmt.Sprintf(":%d", conf.Port), Handler: router}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			fmt.Println("[golith] Starting Gin HTTP Server at", server.Addr)
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("[golith] Shutting down Gin server")
			server.Shutdown(ctx)
			return nil
		},
	})
	return router
}

func Use() fx.Option {
	return fx.Provide(NewGin)
}

func Run() fx.Option {
	return fx.Invoke(func(r *gin.Engine, conf *Configuration) {
		r.Run(fmt.Sprintf(":%d", conf.Port))
	})
}
