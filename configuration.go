package golith

import (
	"github.com/copolio/golith/golithgin"
	"github.com/copolio/golith/golithgorm"
	"github.com/copolio/golith/golithswag"
)

type Configuration struct {
	Gin     *golithgin.Configuration  `yaml:"golithgin"`
	Swagger *golithswag.Configuration `yaml:"golithswag"`
	Gorm    *golithgorm.Configuration `yaml:"golithgorm"`
}

func DefaultConfiguration() Configuration {
	return Configuration{
		Gin:     golithgin.DefaultConfiguration(),
		Swagger: golithswag.DefaultConfiguration(),
		Gorm:    golithgorm.DefaultConfiguration(),
	}
}
