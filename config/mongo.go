package config

import (
	"github.com/gogf/gf/frame/g"
)

type mongoCfg struct {
	Url  string
	Db   string
	User string
	Pass string
}

func (a *mongoCfg) Validate() error {

	return nil
}
func (a *mongoCfg) Load() {
	a.Url = g.Cfg().GetString("mongo.Url")
	if a.Url == "" {
		panic("mongo have no Url")
	}

	get := g.Cfg().Get("mongo.Db")
	if get == nil {
		panic("mongo have no Db")
	}
	a.Db = get.(string)

	get = g.Cfg().Get("mongo.User")
	if get == nil {
		panic("mongo have no User")
	}
	a.User = get.(string)

	get = g.Cfg().Get("mongo.Pass")
	if get == nil {
		panic("mongo have no Pass")
	}
	a.Pass = get.(string)

}
