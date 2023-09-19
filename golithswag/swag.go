package golithswag

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	"os/exec"
)

func NewSwag(config *Configuration, router *gin.Engine) error {
	if config.Use == false {
		return nil
	}

	// Build Swagger document
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}
	cmd := exec.Command(fmt.Sprintf("swag init --parseDependency --parseInternal -g %s -o %s", config.SourcePath, config.OutputPath))
	cmd.Dir = pwd
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Command execution error: %v", err)
	}

	// Define router for publishing generated document
	router.GET(config.RouterPath, ginSwagger.WrapHandler(swaggerfiles.Handler))

	return nil
}
