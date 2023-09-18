package golithgorm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Configuration struct {
	Datasource Datasource
	GormConfig gorm.Config
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
