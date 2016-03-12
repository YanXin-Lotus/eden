package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

var BaseController echo.Context

func init() {
	BaseController.Set("DB", val)
}

func (c BaseController) Prepare() {
	return
}

func (c BaseController) Signin() {
	return
}

func (c BaseController) Signout() {
	return
}

func (c BaseController) CurrentUser() {
	return
}

func (c BaseController) Message() error {
	return c.Render(http.StatusOK, "views/index.html", "test")
}
