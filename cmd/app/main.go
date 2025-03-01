package main

import (
	"log"

	"github.com/auth-service/internal/app"
	"github.com/auth-service/internal/config"
)

func main() {
	appConfig, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	if err := app.Run(appConfig); err != nil {
		log.Fatal(err)
	}
}
