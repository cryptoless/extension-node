package config

import (
	"fmt"
	"os"

	"github.com/gogf/gf/frame/g"
)

type sslCfg struct {
	Crt string
	Key string
}

func (a *sslCfg) Validate() error {

	if a.Crt == "" || a.Key == "" {
		return fmt.Errorf("No sslCfg file")
	}

	if _, err := os.Stat(a.Crt); err != nil {
		return fmt.Errorf("No Crt file:", a.Crt)
	}
	if _, err := os.Stat(a.Key); err != nil {
		return fmt.Errorf("No Key file:", a.Key)
	}
	return nil
}
func (a *sslCfg) Load() {
	a.Crt = g.Cfg().GetString("server.Crt")
	a.Key = g.Cfg().GetString("server.Key")
}
