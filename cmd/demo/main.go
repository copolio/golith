package main

import (
	"github.com/copolio/golith"
)

func main() {
	application := golith.GetWebApplication()
	application.Start()
}
