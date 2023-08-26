package main

import (
	"github.com/copolio/gin-bootify/pkg/config"
	"github.com/copolio/gin-bootify/pkg/data/ddl"
	"github.com/copolio/gin-bootify/pkg/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	ginApplication := &gin.GinApplication{
		Configuration: &config.AutoConfiguration{
			Gin: config.Gin{
				Mode: config.DebugMode,
				Server: config.Server{
					Port: 8080,
				},
				Datasource: config.Datasource{
					Ddl:         ddl.NONE,
					Host:        "localhost:3306",
					Protocol:    "tcp",
					Schema:      "demo",
					User:        "root",
					Password:    "test",
					MaxIdleConn: 10,
					MaxOpenConn: 10,
					Charset:     "utf8",
					ParseTime:   true,
					Loc:         "Local",
				},
			},
		},
	}
	var err error
	ginApplication.Connection, err = gorm.Open(
		mysql.Open(ginApplication.Configuration.Gin.Datasource.GetMysqlDSN()),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   true,
			Logger:                                   logger.Default.LogMode(logger.Info),
		})

	if err != nil {
		log.Fatalf("Error connecting MySQL: %v\n", err)
	}
	sqlDB, err := ginApplication.Connection.DB()
	if err != nil {
		log.Fatalf("Error retrieving MySQL sql connection: %v\n", err)
	}
	sqlDB.SetMaxIdleConns(ginApplication.Configuration.Gin.Datasource.MaxIdleConn)
	sqlDB.SetMaxOpenConns(ginApplication.Configuration.Gin.Datasource.MaxOpenConn)

	ginApplication.Run()
}
