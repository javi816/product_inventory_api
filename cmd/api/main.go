package main

import (
	"fmt"
	"log"

	"github.com/javi816/product_inventory_api/internal/config"
)

func main() {

	cfg, err := config.Load()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("App: ", cfg.AppName)
	fmt.Println("Enviroment: ", cfg.AppEnv)
	fmt.Println("Port: ", cfg.AppPort)
	fmt.Println("Database URL: ", cfg.DatabaseURL)
}
