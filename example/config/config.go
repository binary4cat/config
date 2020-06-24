package config

import "github.com/hjdo/config"

type rootConfig struct {
	Service  service  `yml:"service"`
	Database database `yml:"database"`
}

type service struct {
	Name   string `yml:"name"`
	Consul string `yml:"consul"`
	Manage manage `yml:"manage"`
}

type manage struct {
	Limit   bool   `yml:"limit"`
	Monitor string `yml:"monitor"`
}

type database struct {
	DBSource string `yml:"db_source"`
	ConnStr  string `yml:"conn_str"`
}

// This variable can be called externally to obtain the configuration (the configuration object is not visible to the outside)
var Config = rootConfig{}
var Manage = manage{}

func Init(file string) {
	_ = config.Init(file, "", &Config)
	_ = config.Init(file, "service.manage", &Manage)
}
