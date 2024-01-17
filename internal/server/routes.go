package server

import (
	"github.com/FakharzadehH/GoMonitor/internal/server/handlers"
	"github.com/labstack/echo/v4"
)

func routes(e *echo.Echo, h *handlers.Handlers) {
	api := e.Group("/api")
	api.POST("/server", h.AddServer())
	api.GET("/server", h.ShowServer())
	api.GET("/server/all", h.IndexServers())
}
