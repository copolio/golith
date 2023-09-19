package golithswag

type Configuration struct {
	Use         bool
	RouterPath  string
	SourcePath  string
	OutputPath  string
	Title       string
	Description string
	Version     string
}

func DefaultConfiguration() *Configuration {
	return &Configuration{
		Use:         false,
		RouterPath:  "",
		Title:       "",
		Description: "",
		Version:     "",
	}
}
