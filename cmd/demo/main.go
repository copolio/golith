package main

import (
	"copolio.com/gin-bootify/pkg/config"
)

func main() {
	ginApplication := &config.GinApplication{
		Configuration: &config.AutoConfiguration{
			Gin: config.Gin{
				Mode: config.DebugMode,
				Server: config.Server{
					Port: 8080,
				},
				Datasource: config.Datasource{},
			},
		},
	}
	ginApplication.Run()
}
