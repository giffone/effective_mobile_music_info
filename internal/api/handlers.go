package api

import (
	"music_info/internal/service"

	"github.com/labstack/echo"
)

type Handler interface {
	GetInfoBy(e echo.Context) error
}

func New(service service.Service) Handler {
	return &handler{service: service}
}

type handler struct {
	service service.Service
}

func (h *handler) GetInfoBy(e echo.Context) error {
	return nil
}
