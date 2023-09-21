package golithgorm

import (
	"go.uber.org/fx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(conf *Configuration) (*gorm.DB, error) {
	gormDB, err := gorm.Open(
		mysql.Open(GetMysqlDSN(conf.Datasource)),
		&conf.GormConfig)
	sqlDB, err := gormDB.DB()
	sqlDB.SetMaxIdleConns(conf.Datasource.MaxIdleConn)
	sqlDB.SetMaxOpenConns(conf.Datasource.MaxOpenConn)
	return gormDB, err
}

func Use() fx.Option { return fx.Provide(NewGorm) }
