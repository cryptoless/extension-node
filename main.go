package main

import (
	"extension-node/app/service"
	"extension-node/boot"
	"extension-node/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Log().SetAsync(true)

	boot.BootInit()
	router.RouteInit()
	service.ServiceInit()

	g.Server().Run()
}
