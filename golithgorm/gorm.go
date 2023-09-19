package golithgorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(conf *Configuration) error {
	gormDB, err := gorm.Open(
		mysql.Open(GetMysqlDSN(conf.Datasource)),
		&conf.GormConfig)

	if err != nil {
		return err
	}
	sqlDB, err := gormDB.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(conf.Datasource.MaxIdleConn)
	sqlDB.SetMaxOpenConns(conf.Datasource.MaxOpenConn)
	return nil
}
