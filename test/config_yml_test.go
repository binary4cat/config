package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hjdo/config"
)

type ConfigY struct {
	Service  ServiceY  `yml:"service"`
	Database DatabaseY `yml:"database"`
}

type ServiceY struct {
	Name   string  `yml:"name"`
	Consul string  `yml:"consul"`
	Manage ManageY `yml:"manage"`
}

type ManageY struct {
	Limit   bool   `yml:"limit"`
	Monitor string `yml:"monitor"`
}

type DatabaseY struct {
	DBSource string `yml:"db_source"`
	ConnStr  string `yml:"conn_str"`
}

func TestInitY(t *testing.T) {
	// 将配置文件内容解析到Config结构体中，并实时对应配置文件的改动
	var c ConfigY
	_ = config.Init("./config.yml", "", &c)
	for {
		// 如果修改配置文件的内容，结构体内对应字段的值会同时修改
		fmt.Printf("config is %#v\n", c)
		time.Sleep(10 * time.Second)
	}
}
