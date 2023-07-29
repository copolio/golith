package config

type AppConfig struct {
	Server     Server     `yaml:"server"`
	Datasource Datasource `yaml:"datasource"`
}
