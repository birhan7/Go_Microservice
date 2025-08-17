package main

import (
	"fmt"
	"go-ecommerce/config"
	"go-ecommerce/internal/api"
	"log"
)

func main() {
	cfg, err := config.EnvSetup()
	if err != nil {
		log.Fatalln(err)
	}
	api.StartServer(cfg)

	fmt.Println("Basic folder structure for microservice based Ecommerce application.")
}
