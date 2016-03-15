package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Index(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func Pagination(c *echo.Context) error {
	return c.String(http.StatusOK, "Pagination\n")
}

func Category(c *echo.Context) error {
	return c.String(http.StatusOK, "Category\n")
}
