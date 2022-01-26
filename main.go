package main

import (
	_ "extension-node/boot"
	_ "extension-node/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Log().SetAsync(true)

	g.Server().Run()
}
