package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hjdo/config"
)

type ConfigJ struct {
	Service  ServiceJ  `json:"service"`
	Database DatabaseJ `json:"database"`
}

type ServiceJ struct {
	Name   string  `json:"name"`
	Consul string  `json:"consul"`
	Manage ManageJ `json:"manage"`
}

type ManageJ struct {
	Limit   bool   `json:"limit"`
	Monitor string `json:"monitor"`
}

type DatabaseJ struct {
	DBSource string `json:"db_source"`
	ConnStr  string `json:"conn_str"`
}

func TestInitJ(t *testing.T) {
	// 将配置文件内容解析到Config结构体中，并实时对应配置文件的改动
	var c ConfigJ
	_ = config.Init("./config.json", "", &c)
	for {
		// 如果修改配置文件的内容，结构体内对应字段的值会同时修改
		fmt.Printf("config is %#v\n", c)
		time.Sleep(10 * time.Second)
	}
}
