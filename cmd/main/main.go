package main

import (
	"auth_reg/config"
	"auth_reg/internal/app"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal("Error in parsing config")
	}

	app.Run(cfg)
}
