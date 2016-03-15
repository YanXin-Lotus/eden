package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index(c *echo.Context) error {
	return c.Render(http.StatusOK, "about.html", "")
}

func Pagination(c *echo.Context) error {
	return c.String(http.StatusOK, "Pagination\n")
}

func Category(c *echo.Context) error {
	return c.String(http.StatusOK, "Category\n")
}
