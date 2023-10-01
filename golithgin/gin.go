package golithgin

import (
	"context"
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func NewGin(lc fx.Lifecycle, conf *Configuration) *gin.Engine {
	logger, _ := zap.NewProduction()
	gin.SetMode(string(conf.Mode))
	r := gin.Default()
	r.Use(HttpErrorHandler())
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	server := &http.Server{Addr: fmt.Sprintf(":%d", conf.Port), Handler: r}
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
	return r
}

func Use() fx.Option {
	return fx.Provide(NewGin)
}

func Run() fx.Option {
	return fx.Invoke(func(r *gin.Engine, conf *Configuration) {
		r.Run(fmt.Sprintf(":%d", conf.Port))
	})
}
