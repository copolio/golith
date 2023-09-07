package main

import (
	"github.com/copolio/golith/pkg/web"
)

func main() {
	application := web.GetApplication()
	application.Start()
}
