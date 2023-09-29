package golithgorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Configuration struct {
	Datasource Datasource  `yaml:"datasource"`
	GormConfig gorm.Config `yaml:"gormConfig"`
}

func DefaultConfiguration() *Configuration {
	return &Configuration{
		Datasource: DefaultDatasource(),
		GormConfig: gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   true,
			Logger:                                   logger.Default.LogMode(logger.Info),
		},
	}
}
