package golithswag

type Configuration struct {
	Use        bool
	RouterPath string
}

func DefaultConfiguration() *Configuration {
	return &Configuration{
		Use:        true,
		RouterPath: "/docs/swagger/*any",
	}
}
