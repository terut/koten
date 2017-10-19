package koten

import (
	"log"

	"github.com/labstack/echo"
	"github.com/terut/koten/koten/handler"
)

func registerHandlers(e *echo.Echo) {
	e.GET("/", handler.ListHandler)
	e.GET("/repos/:name", handler.TreeHandler)
}

func Run() {
	e := echo.New()
	registerHandlers(e)

	log.Fatal(e.Start(":8080"))
}
