package koten

import (
	"html/template"
	"log"

	"github.com/labstack/echo"
	"github.com/terut/koten/koten/handler"
)

func registerHandlers(e *echo.Echo) {
	e.GET("/", handler.ListHandler)
	e.GET("/repos/:name", handler.TreeHandler)
}

func registerRenderers(e *echo.Echo) {
	templates := make(map[string]*template.Template)
	templates["a"] = template.Must(template.ParseFiles("koten/views/application.tmpl", "koten/views/hello.tmpl"))
	t := &Template{
		templates: templates,
	}
	e.Renderer = t
}

func Run() {
	e := echo.New()
	registerHandlers(e)
	registerRenderers(e)

	log.Fatal(e.Start(":8080"))
}
