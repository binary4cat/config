package main

import (
	"fmt"

	"github.com/hjdo/config/example/config"
)

func main() {
	config.Init("./config/config.yml")

	// Complete configuration object
	fmt.Println(config.Config)
	// Individually bound configuration items
	fmt.Println(config.Manage)
}
