package koten

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct {
	templates map[string]*template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates[name].ExecuteTemplate(w, "application", data)
}
