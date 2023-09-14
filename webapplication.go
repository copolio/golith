package golith

import (
	"fmt"
	"github.com/copolio/golith/pkg/golithgin"
	"github.com/copolio/golith/pkg/golithgorm"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var instance *WebApplication

type WebApplication struct {
	Configuration Configuration
	GormDB        *gorm.DB
	GinEngine     *gin.Engine
}

func GetWebApplication() *WebApplication {
	if instance != nil {
		return instance
	}
	instance = &WebApplication{
		Configuration: DefaultConfiguration(),
		GormDB:        nil,
		GinEngine:     nil,
	}
	return instance
}

func (application WebApplication) Start() {
	application.initGorm()
	application.initGin()
}

func (application WebApplication) initGin() {
	gin.SetMode(string(application.Configuration.Gin.Mode))
	r := gin.Default()
	r.Use(golithgin.HttpErrorHandler())

	err := r.Run(fmt.Sprintf(":%d", application.Configuration.Gin.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func (application WebApplication) initGorm() {
	var err error
	application.GormDB, err = gorm.Open(
		mysql.Open(golithgorm.GetMysqlDSN(application.Configuration.Gorm.Datasource)),
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
