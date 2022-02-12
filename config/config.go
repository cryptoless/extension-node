package config

var MongoCfg mongoCfg
var RateCfg rateCfg
var SSLCfg sslCfg

func init() {

	(&MongoCfg).Load()
	(&RateCfg).Load()
	(&SSLCfg).Load()
}
