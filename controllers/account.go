package controllers

import (
	"net/http"

	"github.com/labstack/echo"
    "github.com/labstack/echo/engine/standard"
)

func Login(c echo.Context) error {
	return Render.HTML(c.Response().(*standard.Response).ResponseWriter, http.StatusOK, "account/login", nil)
}

func DoLogin(c echo.Context) error {
	return c.Redirect(http.StatusContinue, "/")
}

func Signout(c echo.Context) error {
	return nil
}

func Register(c echo.Context) error {
	return Render.HTML(c.Response().(*standard.Response).ResponseWriter, http.StatusOK, "account/register", nil)
}

func DoRegister(c echo.Context) error {
	return c.Redirect(http.StatusContinue, "/")
}

func Info(c echo.Context) error {
	return Render.HTML(c.Response().(*standard.Response).ResponseWriter, http.StatusOK, "account/info", nil)
}

func EditInfo(c echo.Context) error {
	return Render.HTML(c.Response().(*standard.Response).ResponseWriter, http.StatusOK, "account/editinfo", nil)
}

func DoEditInfo(c echo.Context) error {
	return c.Redirect(http.StatusContinue, "/info")
}

func ChangePW(c echo.Context) error {
	return Render.HTML(c.Response().(*standard.Response).ResponseWriter, http.StatusOK, "account/changepw", nil)
}

func DoChangePW(c echo.Context) error {
	return c.Redirect(http.StatusContinue, "/info")
}
