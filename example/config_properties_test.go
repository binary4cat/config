package example

import (
	"fmt"
	"testing"
	"time"

	"github.com/hjdo/config"
)

type ConfigP struct {
	Service  ServiceP  `properties:"service"`
	Database DatabaseP `properties:"database"`
}

type ServiceP struct {
	Name   string  `properties:"name"`
	Consul string  `properties:"consul"`
	Manage ManageP `properties:"manage"`
}

type ManageP struct {
	Limit   bool   `properties:"limit"`
	Monitor string `properties:"monitor"`
}

type DatabaseP struct {
	DBSource string `properties:"db_source"`
	ConnStr  string `properties:"conn_str"`
}

func TestInitP(t *testing.T) {
	// 将配置文件内容解析到Config结构体中，并实时对应配置文件的改动
	var c ConfigP
	_ = config.Init("./config.properties", "", &c)
	for {
		// 如果修改配置文件的内容，结构体内对应字段的值会同时修改
		fmt.Printf("config is %#v\n", c)
		time.Sleep(10 * time.Second)
	}
}
