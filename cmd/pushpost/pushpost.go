package main

import (
	"log"
	"pushpost/internal/app"
	"pushpost/internal/config"
)

func main() {
	conf, err := config.LoadYamlConfig("configs/user_service.yaml")

	if err != nil {

		log.Println(err)
	}

	app.Run(conf)

}
