package gorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Configuration struct {
	Datasource    Datasource
	Configuration gorm.Config
}

func DefaultConfiguration() Configuration {
	return Configuration{
		Datasource: DefaultDatasource(),
		Configuration: gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   true,
			Logger:                                   logger.Default.LogMode(logger.Info),
		},
	}
}
