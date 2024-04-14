package handler

import "github.com/labstack/echo/v4"

type healthCheckHandler struct{}

func NewHealthCheckHandler() *healthCheckHandler {
	return &healthCheckHandler{}
}

func (h *healthCheckHandler) Show(ctx echo.Context) error {
	return ctx.JSON(200, echo.Map{
		"message": "Ok",
	})
}
