package main

import (
	"github.com/copolio/gin-bootify/pkg"
	"github.com/copolio/gin-bootify/pkg/config"
	"github.com/copolio/gin-bootify/pkg/database/ddl"
)

func main() {
	ginApplication := &pkg.GinApplication{
		Configuration: &config.AutoConfiguration{
			Gin: config.Gin{
				Mode: config.DebugMode,
				Server: config.Server{
					Port: 8080,
				},
				Datasource: config.Datasource{
					Ddl:         ddl.NONE,
					Host:        "localhost:3306",
					Protocol:    "tcp",
					Schema:      "demo",
					User:        "root",
					Password:    "test",
					MaxIdleConn: 0,
					MaxOpenConn: 0,
				},
			},
		},
	}
	ginApplication.Run()
}
