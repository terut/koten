package koten

import (
	"log"

	"github.com/labstack/echo"
)

func registerHandlers(e *echo.Echo) {
	e.GET("/", ListHandler)
	e.GET("/repos/:name", TreeHandler)
}

func Run() {
	e := echo.New()
	registerHandlers(e)
	log.Fatal(e.Start(":8080"))
}
