package web

import (
	"fmt"
	"github.com/copolio/golith/pkg/core"
	gin2 "github.com/copolio/golith/pkg/gin"
	gorm2 "github.com/copolio/golith/pkg/gorm"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var instance *Application

type Application struct {
	Configuration core.Configuration
	GormDB        *gorm.DB
	GinEngine     *gin.Engine
}

func GetApplication() *Application {
	if instance != nil {
		return instance
	}
	instance = &Application{
		Configuration: core.DefaultConfiguration(),
		GormDB:        nil,
		GinEngine:     nil,
	}
	return instance
}

func (application Application) Start() {
	application.initGorm()
	application.initGin()
}

func (application Application) initGin() {
	gin.SetMode(string(application.Configuration.Gin.Mode))
	r := gin.Default()
	r.Use(gin2.HttpErrorHandler())

	err := r.Run(fmt.Sprintf(":%d", application.Configuration.Gin.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func (application Application) initGorm() {
	var err error
	application.GormDB, err = gorm.Open(
		mysql.Open(gorm2.GetMysqlDSN(application.Configuration.Gorm.Datasource)),
		&application.Configuration.Gorm.Configuration)

	if err != nil {
		log.Fatalf("Error connecting MySQL: %v\n", err)
	}
	sqlDB, err := application.GormDB.DB()
	if err != nil {
		log.Fatalf("Error retrieving MySQL sql connection: %v\n", err)
	}
	sqlDB.SetMaxIdleConns(application.Configuration.Gorm.Datasource.MaxIdleConn)
	sqlDB.SetMaxOpenConns(application.Configuration.Gorm.Datasource.MaxOpenConn)
}
