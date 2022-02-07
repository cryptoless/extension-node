package config

var MongoCfg mongoCfg
var RateCfg rateCfg

func init() {

	(&MongoCfg).Load()
	(&RateCfg).Load()
}
