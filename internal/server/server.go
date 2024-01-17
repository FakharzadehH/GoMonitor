package server

import (
	"github.com/FakharzadehH/GoMonitor/internal/config"
	"github.com/FakharzadehH/GoMonitor/internal/server/handlers"
	"github.com/FakharzadehH/GoMonitor/repository"
	"github.com/FakharzadehH/GoMonitor/service"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

func Start() error {
	e := echo.New()
	e.Use(echoMiddleware.Logger())
	e.HTTPErrorHandler = ErrorHandler()
	db, err := config.NewGORMConnection(config.GetConfig())
	if err != nil {
		return err
	}

	repos := repository.NewRepository(db)
	svcs := service.NewService(repos)
	handler := handlers.New(repos, svcs)
	routes(e, handler)
	return e.Start(":1323")
}
