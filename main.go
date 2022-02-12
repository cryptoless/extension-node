package main

import (
	"extension-node/app/service"
	"extension-node/boot"
	"extension-node/config"
	"extension-node/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Log().SetAsync(true)

	boot.BootInit()
	router.RouteInit()
	service.ServiceInit()

	// g.Server().SetServerRoot(gfile.MainPkgPath())
	if err := config.SSLCfg.Validate(); err == nil {
		g.Server().EnableHTTPS(config.SSLCfg.Crt, config.SSLCfg.Key)
	} else {
		g.Log().Error("EnableHttps err:", err)
	}
	g.Server().Run()
}
