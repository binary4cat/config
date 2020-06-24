package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hjdo/config"
)

type ConfigI struct {
	Service  ServiceI  `ini:"service"`
	Database DatabaseI `ini:"database"`
}

type ServiceI struct {
	Name   string  `ini:"name"`
	Consul string  `ini:"consul"`
	Manage ManageI `ini:"manage"`
}

type ManageI struct {
	Limit   bool   `ini:"limit"`
	Monitor string `ini:"monitor"`
}

type DatabaseI struct {
	DBSource string `ini:"db_source"`
	ConnStr  string `ini:"conn_str"`
}

func TestInitI(t *testing.T) {
	// 将配置文件内容解析到Config结构体中，并实时对应配置文件的改动
	var c ConfigI
	_ = config.Init("./config.ini", "", &c)
	for {
		// 如果修改配置文件的内容，结构体内对应字段的值会同时修改
		fmt.Printf("config is %#v\n", c)
		time.Sleep(10 * time.Second)
	}
}
