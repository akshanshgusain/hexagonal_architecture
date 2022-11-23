package main

import (
	"github.com/akshanshgusain/Hexagonal-Architecture/app"
	"github.com/akshanshgusain/Hexagonal-Architecture/logger"
)

func main() {
	logger.Info("starting the application...")
	app.Start()
}
