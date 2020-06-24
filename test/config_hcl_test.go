package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hjdo/config"
)

type ConfigH struct {
	Service  ServiceH  `hcl:"service"`
	Database DatabaseH `hcl:"database"`
}

type ServiceH struct {
	Name   string  `hcl:"name"`
	Consul string  `hcl:"consul"`
	Manage ManageH `hcl:"manage"`
}

type ManageH struct {
	Limit   bool   `hcl:"limit"`
	Monitor string `hcl:"monitor"`
}

type DatabaseH struct {
	DBSource string `hcl:"db_source"`
	ConnStr  string `hcl:"conn_str"`
}

func TestInitH(t *testing.T) {
	// 将配置文件内容解析到Config结构体中，并实时对应配置文件的改动
	var c ConfigH
	_ = config.Init("./config.hcl", "", &c)
	for {
		// 如果修改配置文件的内容，结构体内对应字段的值会同时修改
		fmt.Printf("config is %#v\n", c)
		time.Sleep(10 * time.Second)
	}
}
