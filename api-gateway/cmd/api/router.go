package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func AddRouter(server *echo.Echo) {
	api := server.Group("/api")
	version := api.Group("/v1")

	version.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
}
