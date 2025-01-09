package main

import (
	"fmt"
	"pushpost/internal/config"
	"pushpost/internal/setup"
)

func main() {
	conf, err := config.LoadConfig("configs/development.yaml")
	if err != nil {
		fmt.Println(err)
	}
	_, err = setup.Setup(*conf)

	if err != nil {
		fmt.Println(err)
	}

}
