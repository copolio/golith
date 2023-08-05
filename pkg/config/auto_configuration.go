package config

type AutoConfiguration struct {
	Gin     Gin
	Swagger Swagger
}

type Gin struct {
	Mode       GinMode
	Server     Server
	Datasource Datasource
}

type GinMode string

const (
	DebugMode   GinMode = "debug"
	ReleaseMode GinMode = "release"
	TestMode    GinMode = "test"
)

type Server struct {
	Port uint `yaml:"port"`
}
