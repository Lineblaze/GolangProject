package main

import (
	"GolangProject/internal/app"
	"GolangProject/internal/config"
	"GolangProject/pkg/logging"
	"log"
)

func main() {
	log.Print("config initializing")
	cfg := config.GetConfig()

	log.Print("logger initializing")
	logging.Init(cfg.AppConfig.LogLevel)
	logger := logging.GetLogger()

	a, err := app.NewApp(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Println("Running application")
	a.Run()
}
