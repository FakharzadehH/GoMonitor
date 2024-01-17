package handlers

import (
	"errors"
	"github.com/FakharzadehH/GoMonitor/internal/domain"
	"github.com/FakharzadehH/GoMonitor/internal/logger"
	"github.com/FakharzadehH/GoMonitor/repository"
	"github.com/FakharzadehH/GoMonitor/service"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type Handlers struct {
	repos *repository.Repository
	svcs  *service.Service
	log   *zap.SugaredLogger
}

func New(repos *repository.Repository, svcs *service.Service) *Handlers {
	return &Handlers{
		repos: repos,
		svcs:  svcs,
		log:   logger.Logger(),
	}
}

func (h *Handlers) AddServer() echo.HandlerFunc {
	type request struct {
		domain.AddServerRequest
	}
	type response struct {
		domain.AddServerResponse
	}
	return func(c echo.Context) error {
		var req request
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, nil)
		}
		res, err := h.svcs.AddServer(req.AddServerRequest)
		if err != nil {
			return err
		}
		return c.JSON(200, &response{
			*res,
		})
	}
}

func (h *Handlers) ShowServer() echo.HandlerFunc {
	type response struct {
		domain.StatusShowResponse
	}
	return func(c echo.Context) error {
		if c.QueryParam("id") == "" {
			return c.JSON(http.StatusBadRequest, errors.New("please enter an ID"))
		}
		serverID, err := strconv.ParseUint(c.QueryParam("id"), 10, 64)
		if err != nil {
			return err
		}
		res, err := h.svcs.GetServerByID(uint(serverID))
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			*res,
		})
	}
}

func (h *Handlers) IndexServers() echo.HandlerFunc {
	type response struct {
		domain.StatusIndexResponse
	}
	return func(c echo.Context) error {
		res, err := h.svcs.GetAllServers()
		if err != nil {
			return err
		}
		return c.JSON(200, response{
			StatusIndexResponse: *res,
		})
	}
}
