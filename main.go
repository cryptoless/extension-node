package main

import (
	"extension-node/boot"
	"extension-node/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Log().SetAsync(true)

	boot.BootInit()
	router.RouteInit()

	g.Server().Run()
}
