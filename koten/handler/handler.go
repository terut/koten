package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func ListHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprint("show repositories!"))
}

func TreeHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprintf("tree %s!", c.Param("name")))
}
