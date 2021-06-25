package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/vehsamrak/capital/internal/app"
	"github.com/vehsamrak/capital/internal/app/config"
	"github.com/vehsamrak/capital/internal/logger"
	"github.com/vehsamrak/capital/internal/renderer"
)

func main() {
	log.SetFormatter(&logger.TextFormatter{})

	log.Info("Capital app started")

	configParser := config.Parser{}
	capitalConfig, err := configParser.Parse()
	if err != nil {
		return
	}

	capitalResult := app.CalculateCapital(capitalConfig)
	consoleRenderer := renderer.Console{}

	render, err := consoleRenderer.Render(capitalResult)
	if err != nil {
		log.WithError(err).Error("Render failed with an error")
		return
	}

	fmt.Printf("%s\n", render)

	log.Info("Capital app finished")
}
