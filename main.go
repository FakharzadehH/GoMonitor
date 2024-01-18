package main

import (
	"github.com/FakharzadehH/GoMonitor/internal/config"
	"github.com/FakharzadehH/GoMonitor/internal/job"
	"github.com/FakharzadehH/GoMonitor/internal/logger"
	"github.com/FakharzadehH/GoMonitor/internal/server"
	"log"
)

func main() {
	if err := config.Load("config.yaml"); err != nil {
		log.Fatal("error loading config")
	}
	go job.CheckServersHealthJob()
	logger.Init()
	logger.Logger().Fatal(server.Start())
}
