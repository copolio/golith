package golithgin

import (
	"github.com/gin-gonic/gin"
)

type Configuration struct {
	Mode   Mode
	Server ServerConfiguration
}

type Mode string

const (
	DebugMode   Mode = gin.DebugMode
	ReleaseMode Mode = gin.ReleaseMode
	TestMode    Mode = gin.TestMode
)

func DefaultGinConfig() Configuration {
	return Configuration{
		Mode:   DebugMode,
		Server: DefaultServerConfig(),
	}
}

type ServerConfiguration struct {
	Port uint `yaml:"port"`
}

func DefaultServerConfig() ServerConfiguration {
	return ServerConfiguration{Port: 8080}
}
