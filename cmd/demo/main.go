package main

import (
	"github.com/copolio/gin-bootify/pkg/web"
)

func main() {
	application := web.GetApplication()
	application.Start()
}
