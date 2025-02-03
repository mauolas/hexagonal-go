package main

import (
	"example.com/hexagonal/app"
	"example.com/hexagonal/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
