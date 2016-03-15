package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

var BC echo.Context

func Prepare() {
	return
}

func About(c *echo.Context) error {
	return c.Render(http.StatusOK, "about.html", "")
}

func Friendship(c *echo.Context) error {
	return c.String(http.StatusOK, "Friendship")
}

func Signin() {
	return
}

func Signout() {
	return
}

func CurrentUser() {
	return
}
