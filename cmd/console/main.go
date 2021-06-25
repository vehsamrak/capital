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

	configParser := config.Parser{}
	capitalConfig, isVerbose, err := configParser.Parse()
	if err != nil {
		log.WithError(err).Error("Capital app error occurred")
		return
	}

	if isVerbose {
		log.SetLevel(log.DebugLevel)
	}

	if capitalConfig == nil {
		return
	}

	log.Debug("Capital app started")

	capitalResult := app.CalculateCapital(capitalConfig)
	consoleRenderer := renderer.Console{}

	render, err := consoleRenderer.Render(capitalResult)
	if err != nil {
		log.WithError(err).Error("Render failed with an error")
		return
	}

	fmt.Printf("%s\n", render)

	log.Debug("Capital app finished")
}
