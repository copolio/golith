package gin

import (
	"fmt"
	"github.com/copolio/gin-bootify/pkg/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type GinApplication struct {
	Configuration *config.AutoConfiguration
	Connection    *gorm.DB
}

func (app *GinApplication) Run() {
	gin.SetMode(string(app.Configuration.Gin.Mode))
	ginApplication := gin.Default()
	ginApplication.Use(BasicErrorHandler())

	err := ginApplication.Run(fmt.Sprintf(":%d", app.Configuration.Gin.Server.Port))
	if err != nil {
		log.Fatal(err)
	}
}
