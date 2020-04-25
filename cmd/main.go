package main

import (
	"fmt"
	"log"

	"github.com/metacatdud/go-boilerplate/config"
	"github.com/metacatdud/go-boilerplate/infrastructure/router"
	"github.com/metacatdud/go-boilerplate/interface/controller"
)

var (
	envArg string
)

func main() {
	config.Read()

	r := router.NewRouter(controller.NewUserController())
	fmt.Println("Server listen at http://localhost" + ":" + config.Config.Server.Address)
	if err := r.Start(":" + config.Config.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
