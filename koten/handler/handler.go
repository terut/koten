package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/terut/koten/koten/repository"
)

func ListHandler(c echo.Context) error {
	return c.String(http.StatusOK, fmt.Sprint("show repositories!"))
}

func TreeHandler(c echo.Context) error {
	entry := repository.Walk(c.Param("name"), []string{".git"})
	return c.JSON(http.StatusOK, entry)
}

func CloneHandler(c echo.Context) error {
	repository.Clone()
	return c.String(http.StatusOK, fmt.Sprint("clone!"))
}
