package handler

import "github.com/labstack/echo/v4"

type slackHandler struct{}

func NewSlackHandler() *slackHandler {
	return &slackHandler{}
}

func (h *slackHandler) Create(ctx echo.Context) error {
	return ctx.JSON(200, echo.Map{
		"message": "Ok",
	})
}
