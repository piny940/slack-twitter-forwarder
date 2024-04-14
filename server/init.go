package server

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() error {
	e := echo.New()

	// router := e.Group(os.Getenv("SERVER_VERSION"))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":" + os.Getenv("SERVER_PORT")))
	return nil
}
