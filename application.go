package golith

import (
	"fmt"
	"github.com/copolio/golith/golithgin"
	"github.com/copolio/golith/golithgorm"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var instance *Application = nil

type Application struct {
	Configuration Configuration
	GormDB        *gorm.DB
	GinEngine     *gin.Engine
}

func NewApplication() (webApplication *Application) {
	if instance != nil {
		return instance
	}
	instance = &Application{
		Configuration: DefaultConfiguration(),
		GormDB:        nil,
		GinEngine:     nil,
	}
	initGin()
	initGorm()
	return instance
}

func (application Application) Run() {
	err := application.GinEngine.Run(fmt.Sprintf(":%d", application.Configuration.Gin.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}

func initGin() {
	gin.SetMode(string(instance.Configuration.Gin.Mode))
	instance.GinEngine = gin.Default()
	instance.GinEngine.Use(golithgin.HttpErrorHandler())
}

func initGorm() {
	var err error
	instance.GormDB, err = gorm.Open(
		mysql.Open(golithgorm.GetMysqlDSN(instance.Configuration.Gorm.Datasource)),
		&instance.Configuration.Gorm.Configuration)

	if err != nil {
		log.Fatalf("Error connecting MySQL: %v\n", err)
	}
	sqlDB, err := instance.GormDB.DB()
	if err != nil {
		log.Fatalf("Error retrieving MySQL sql connection: %v\n", err)
	}
	sqlDB.SetMaxIdleConns(instance.Configuration.Gorm.Datasource.MaxIdleConn)
	sqlDB.SetMaxOpenConns(instance.Configuration.Gorm.Datasource.MaxOpenConn)
}
