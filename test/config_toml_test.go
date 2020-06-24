package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hjdo/config"
)

type ConfigT struct {
	Service  ServiceT  `toml:"service"`
	Database DatabaseT `toml:"database"`
}

type ServiceT struct {
	Name   string  `toml:"name"`
	Consul string  `toml:"consul"`
	Manage ManageT `toml:"manage"`
}

type ManageT struct {
	Limit   bool   `toml:"limit"`
	Monitor string `toml:"monitor"`
}

type DatabaseT struct {
	DBSource string `toml:"db_source"`
	ConnStr  string `toml:"conn_str"`
}

func TestInitT(t *testing.T) {
	// 将配置文件内容解析到Config结构体中，并实时对应配置文件的改动
	var c ConfigT
	_ = config.Init("./config.toml", "", &c)
	for {
		// 如果修改配置文件的内容，结构体内对应字段的值会同时修改
		fmt.Printf("config is %#v\n", c)
		time.Sleep(10 * time.Second)
	}
}
