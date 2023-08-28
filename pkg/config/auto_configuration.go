package config

type GinMode string

const (
	DebugMode   GinMode = "debug"
	ReleaseMode GinMode = "release"
	TestMode    GinMode = "test"
)

type AutoConfiguration struct {
	Gin     GinConfiguration     `yaml:"gin"`
	Swagger SwaggerConfiguration `yaml:"swagger"`
}

func DefaultAutoConfiguration() AutoConfiguration {
	return AutoConfiguration{
		Gin:     DefaultGinConfig(),
		Swagger: DefaultSwaggerConfiguration(),
	}
}

type GinConfiguration struct {
	Mode       GinMode
	Server     ServerConfiguration
	Datasource Datasource
}

func DefaultGinConfig() GinConfiguration {
	return GinConfiguration{
		Mode:       DebugMode,
		Server:     DefaultServerConfig(),
		Datasource: DefaultDatasource(),
	}
}

type ServerConfiguration struct {
	Port uint `yaml:"port"`
}

func DefaultServerConfig() ServerConfiguration {
	return ServerConfiguration{Port: 8080}
}
