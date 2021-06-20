package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/vehsamrak/capital/internal/app"
	"github.com/vehsamrak/capital/internal/logger"
	"github.com/vehsamrak/capital/internal/renderer"
)

func main() {
	log.SetFormatter(&logger.TextFormatter{})

	log.Info("Capital app started")

	capitalResult := app.CalculateCapital()
	consoleRenderer := renderer.Console{}

	consoleRenderer.Render(capitalResult)

	log.Info("Capital app finished")
}
