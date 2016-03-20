package controllers

import (
	"net/http"

	"github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

func Index(c echo.Context) error {
	return Render.HTML(c.Response().(*standard.Response).ResponseWriter, http.StatusOK, "index", nil)
}

func Pagination(c echo.Context) error {
	return c.String(http.StatusOK, "Pagination\n")
}

func Category(c echo.Context) error {
	return c.String(http.StatusOK, "Category\n")
}

func Article(c echo.Context) error {
	return nil
}

func EditArticle(c echo.Context) error {
	return nil
}

func DoEditArticle(c echo.Context) error {
	return nil
}
