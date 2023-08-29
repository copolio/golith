package core

import (
	"github.com/copolio/gin-bootify/pkg/gin"
	"github.com/copolio/gin-bootify/pkg/gorm"
	"github.com/copolio/gin-bootify/pkg/swag"
)

type Configuration struct {
	Gin     gin.Configuration  `yaml:"gin"`
	Swagger swag.Configuration `yaml:"swag"`
	Gorm    gorm.Configuration `yaml:"gorm"`
}

func DefaultConfiguration() Configuration {
	return Configuration{
		Gin:     gin.DefaultGinConfig(),
		Swagger: swag.DefaultConfiguration(),
		Gorm:    gorm.DefaultConfiguration(),
	}
}
