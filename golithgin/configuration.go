package golithgin

import (
	"github.com/gin-gonic/gin"
)

type Configuration struct {
	Mode Mode `yaml:"mode"`
	Port uint `yaml:"port"`
}

type Mode string

const (
	DebugMode   Mode = gin.DebugMode
	ReleaseMode Mode = gin.ReleaseMode
	TestMode    Mode = gin.TestMode
)

func DefaultConfiguration() *Configuration {
	return &Configuration{
		Mode: DebugMode,
		Port: 8080,
	}
}
