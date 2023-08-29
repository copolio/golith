package swag

type Configuration struct {
	Use         bool
	Path        string
	Title       string
	Description string
	Version     string
}

func DefaultConfiguration() Configuration {
	return Configuration{
		Use:         false,
		Path:        "",
		Title:       "",
		Description: "",
		Version:     "",
	}
}
