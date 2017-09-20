package koten

import (
	"html/template"
	"log"

	"github.com/labstack/echo"
)

func registerHandlers(e *echo.Echo) {
	e.GET("/", ListHandler)
	e.GET("/repos/:name", TreeHandler)
}

func registerRenderers(e *echo.Echo) {
	templates := make(map[string]*template.Template)
	templates["hello"] = template.Must(template.ParseFiles("koten/views/application.html", "koten/views/hello.html"))
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
