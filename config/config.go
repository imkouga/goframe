package config

import (
	"github.com/imkouga/gocore/cfg/conf"
)

func Init(cfgPath string) error {
	return conf.LoadConfig(cfgPath)
}

func Reload() error {
	return nil
}
