package main

import (
	"github.com/FakharzadehH/GoMonitor/internal/config"
	"github.com/FakharzadehH/GoMonitor/internal/logger"
	"github.com/FakharzadehH/GoMonitor/internal/server"
	"log"
)

func main() {
	if err := config.Load("config.yaml"); err != nil {
		log.Fatal("error loading config")
	}
	logger.Init()
	logger.Logger().Fatal(server.Start())
}
