package config

var MongoCfg mongoCfg

func init() {

	(&MongoCfg).Load()
}
