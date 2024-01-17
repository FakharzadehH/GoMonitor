package server

import (
	"github.com/FakharzadehH/GoMonitor/internal/config"
	"github.com/FakharzadehH/GoMonitor/repository"
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
	//svcs := service.NewServices(repos)
	//middlewares := middleware.NewMiddlewares(svcs, repos)
	//handler := handlers.New(repos, svcs)
	//routes(e, handler, middlewares)
	return e.Start(":1323")
}
g