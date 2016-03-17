package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

func Login(c *echo.Context) error {
	return Render.HTML(c.Response().Writer(), http.StatusOK, "account/login", nil)
}

func DoLogin(c *echo.Context) error {
	return c.Redirect(http.StatusContinue, "/")
}

func Signout(c *echo.Context) error {
	return nil
}

func Register(c *echo.Context) error {
	return Render.HTML(c.Response().Writer(), http.StatusOK, "account/register", nil)
}

func DoRegister(c *echo.Context) error {
	return c.Redirect(http.StatusContinue, "/")
}

func Info(c *echo.Context) error {
	return Render.HTML(c.Response().Writer(), http.StatusOK, "account/info", nil)
}

func EditInfo(c *echo.Context) error {
	return Render.HTML(c.Response().Writer(), http.StatusOK, "account/editinfo", nil)
}

func DoEdit(c *echo.Context) error {
	return c.Redirect(http.StatusContinue, "/info")
}

func ChangePW(c *echo.Context) error {
	return Render.HTML(c.Response().Writer(), http.StatusOK, "account/changepw", nil)
}

func DoChange(c *echo.Context) error {
	return c.Redirect(http.StatusContinue, "/info")
}
