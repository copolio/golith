package main

import (
	"github.com/copolio/golith"
)

func main() {
	webApplication := golith.GetWebApplication()
	webApplication.Start()
}
