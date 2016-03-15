package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Login(c *echo.Context) error {
	return c.Render(http.StatusOK, "account/login.html", nil)
}

func DoLogin(c *echo.Context) error {
	return c.Redirect(http.StatusContinue, "/")
}

func Register(c *echo.Context) error {
	return c.Render(http.StatusOK, "views/account/register.html", nil)
}

func DoRegister(c *echo.Context) error {
	return c.Redirect(http.StatusContinue, "/")
}

func Info(c *echo.Context) error {
	return c.Render(http.StatusOK, "views/account/info.html", nil)
}

func EditInfo(c *echo.Context) error {
	return c.Render(http.StatusOK, "views/account/editinfo.html", nil)
}

func DoEdit(c *echo.Context) error {
	return c.Redirect(http.StatusContinue, "/info")
}

func ChangePW(c *echo.Context) error {
	return c.Render(http.StatusOK, "views/account/changepw.html", nil)
}

func DoChange(c *echo.Context) error {
	return c.Redirect(http.StatusContinue, "/info")
}
