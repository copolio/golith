package golith

import (
	"github.com/copolio/golith/pkg/golithgin"
	"github.com/copolio/golith/pkg/golithgorm"
	"github.com/copolio/golith/pkg/golithswag"
)

type Configuration struct {
	Gin     golithgin.Configuration  `yaml:"golithgin"`
	Swagger golithswag.Configuration `yaml:"golithswag"`
	Gorm    golithgorm.Configuration `yaml:"golithgorm"`
}

func DefaultConfiguration() Configuration {
	return Configuration{
		Gin:     golithgin.DefaultGinConfig(),
		Swagger: golithswag.DefaultConfiguration(),
		Gorm:    golithgorm.DefaultConfiguration(),
	}
}
