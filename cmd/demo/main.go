package main

import (
	"copolio.com/gin-bootify/pkg/config"
)

func main() {
	ginApplication := &config.GinApplication{
		Configuration: &config.Gin{
			Server: config.Server{
				Port: 80,
			},
			Datasource: config.Datasource{},
		},
	}
	ginApplication.Run()
}
