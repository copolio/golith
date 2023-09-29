package golith

import (
	"github.com/copolio/golith/golithgin"
	"github.com/copolio/golith/golithgorm"
	"github.com/copolio/golith/golithswag"
)

type Configuration struct {
	Gin     golithgin.Configuration
	Gorm    golithgorm.Configuration
	Swagger golithswag.Configuration
}
