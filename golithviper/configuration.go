package golithviper

import (
	"fmt"
	"github.com/copolio/golith"
	"github.com/copolio/golith/golithgin"
	"github.com/copolio/golith/golithgorm"
	"github.com/copolio/golith/golithswag"
	"github.com/spf13/viper"
	"log"
)

func LoadConfiguration() (*golithgorm.Configuration, *golithgin.Configuration, *golithswag.Configuration) {
	log.Default().Println("Init AppConfig")
	var config *golith.Configuration
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("/config/") // path to look for the config file in
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Errorf("fatal error config unmarshall: %w", err))
	}

	return &config.Gorm, &config.Gin, &config.Swagger
}
