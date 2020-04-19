package main

import (
	"fmt"

	"github.com/metacatdud/go-boilerplate/config"
)

var (
	envArg string
)

func main() {
	config.Read()
	fmt.Printf("Hello %t", config.Config.Debug)
}
