package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hjdo/config"
)

type ConfigE struct {
	Service  ServiceE  `env:"service"`
	Database DatabaseE `env:"database"`
}

type ServiceE struct {
	Name   string  `env:"name"`
	Consul string  `env:"consul"`
	Manage ManageE `env:"manage"`
}

type ManageE struct {
	Limit   bool   `env:"limit"`
	Monitor string `env:"monitor"`
}

type DatabaseE struct {
	DBSource string `env:"db_source"`
	ConnStr  string `env:"conn_str"`
}

func TestInitE(t *testing.T) {
	// 将配置文件内容解析到Config结构体中，并实时对应配置文件的改动
	var c ConfigE
	_ = config.Init("./config.env", "", &c)
	for {
		// 如果修改配置文件的内容，结构体内对应字段的值会同时修改
		fmt.Printf("config is %#v\n", c)
		time.Sleep(10 * time.Second)
	}
}
