package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Config struct {
	Port string
	Dsn  string
}

type App struct {
	Server *echo.Echo
	Cfg    *Config
}

func NewApp(port, dsn string) *App {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	app := App{
		Server: e,
		Cfg: &Config{
			Port: port,
			Dsn:  dsn,
		},
	}

	return &app
}
