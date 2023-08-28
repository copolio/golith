package main

import (
	"github.com/copolio/gin-bootify/pkg/config"
	"github.com/copolio/gin-bootify/pkg/gin"
)

func main() {
	ginApplication := &gin.Application{
		Configuration: config.DefaultAutoConfiguration(),
	}
	ginApplication.Run()
}
