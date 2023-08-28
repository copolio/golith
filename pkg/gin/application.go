package gin

import (
	"fmt"
	"github.com/copolio/gin-bootify/pkg/config"
	gorm2 "github.com/copolio/gin-bootify/pkg/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type Application struct {
	Configuration config.AutoConfiguration
	Connection    *gorm.DB
}

func (application Application) Run() {
	gin.SetMode(string(application.Configuration.Gin.Mode))
	r := gin.Default()
	r.Use(BasicErrorHandler())
	application.initDatasource()

	err := r.Run(fmt.Sprintf(":%d", application.Configuration.Gin.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func (application Application) initDatasource() {
	var err error
	application.Connection, err = gorm.Open(
		mysql.Open(gorm2.GetMysqlDSN(application.Configuration.Gin.Datasource)),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			SkipDefaultTransaction:                   true,
			Logger:                                   logger.Default.LogMode(logger.Info),
		})

	if err != nil {
		log.Fatalf("Error connecting MySQL: %v\n", err)
	}
	sqlDB, err := application.Connection.DB()
	if err != nil {
		log.Fatalf("Error retrieving MySQL sql connection: %v\n", err)
	}
	sqlDB.SetMaxIdleConns(application.Configuration.Gin.Datasource.MaxIdleConn)
	sqlDB.SetMaxOpenConns(application.Configuration.Gin.Datasource.MaxOpenConn)
}
