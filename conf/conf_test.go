package conf

import (
	"github.com/yhh-show/helpers/jsons"
	"testing"
)

type demoConf struct {
	App struct {
		Env   string `env:"APP_ENV" envDefault:"local"`
		Debug bool   `env:"APP_DEBUG" envDefault:"true"`
	}
	Web struct {
		Addr string `env:"ADDR" envDefault:":9999"`
	} `envPrefix:"WEB_"`
}

func TestLoad(t *testing.T) {
	r := ForceLoad[demoConf]()
	jsons.ToStringPretty(r)
}
