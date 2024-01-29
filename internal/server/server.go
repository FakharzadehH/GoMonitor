package server

import (
	"github.com/FakharzadehH/GoMonitor/internal/config"
	"github.com/FakharzadehH/GoMonitor/internal/server/handlers"
	"github.com/FakharzadehH/GoMonitor/repository"
	"github.com/FakharzadehH/GoMonitor/service"
	"github.com/labstack/echo-contrib/echoprometheus"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"os"
)

func Start() error {
	e := echo.New()

	e.Use(echoprometheus.NewMiddleware("echo"))
	e.GET("/metrics", echoprometheus.NewHandler())

	e.Use(echoMiddleware.Logger())
	e.HTTPErrorHandler = ErrorHandler()
	writeDB, err := config.NewGORMConnection(config.GetConfig().DB.GetWriteDSN())
	if err != nil {
		return err
	}
	readDB, err := config.NewGORMConnection(config.GetConfig().DB.GetReadDSN())
	if err != nil {
		return err
	}

	writeRepos := repository.NewRepository(writeDB)
	readRepos := repository.NewRepository(readDB)
	svcs := service.NewService(writeRepos, readRepos)
	handler := handlers.New(svcs)
	routes(e, handler)
	return e.Start(":" + os.Getenv("APP_PORT"))
}
