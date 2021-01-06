package main

import (
	"fmt"

	"github.com/hjdo/config/example/config"
)

func main() {
	config.Init("example/config/config.yml")

	// Complete configuration object
	fmt.Printf("%+v\n", config.Config)
	// Individually bound configuration items
	fmt.Printf("%+v\n", config.Manage)
}
