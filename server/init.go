package server

import (
	"os"
	"slack-twitter-forwarder/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() error {
	e := echo.New()

	router := e.Group(os.Getenv("SERVER_VERSION"))

	healthCheckHandler := handler.NewHealthCheckHandler()
	router.GET("", healthCheckHandler.Show)

	slackHandler := handler.NewSlackHandler()
	router.POST("/slack", slackHandler.Create)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
	return nil
}
