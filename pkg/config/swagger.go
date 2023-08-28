package config

type SwaggerConfiguration struct {
	Use         bool
	Path        string
	Title       string
	Description string
	Version     string
}

func DefaultSwaggerConfiguration() SwaggerConfiguration {
	return SwaggerConfiguration{
		Use:         false,
		Path:        "",
		Title:       "",
		Description: "",
		Version:     "",
	}
}
