package controllers

import (
	"eden/models"
	"eden/services"
	"net/http"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	user := currentUser(c)
	if user != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	return c.Render(http.StatusOK, "login", nil)
}

//not finish, need param bind(user)
func DoLogin(c echo.Context) error {
	var user models.User
	err := services.Login(&user)
	if err != nil {
		return c.JSON(http.StatusForbidden, &retJson{OK: false, Desc: "Error password"})
	}
	err = setUser(&user, c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &retJson{OK: false, Desc: "Error set session"})
	}
	return c.JSON(http.StatusOK, &retJson{OK: true, Desc: ""})
}

func Signout(c echo.Context) error {
	return c.Redirect(http.StatusTemporaryRedirect, "/")
}

func Register(c echo.Context) error {
	user := currentUser(c)
	if user != nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
	return c.Render(http.StatusOK, "account/register", nil)
}

func DoRegister(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func Info(c echo.Context) error {
	user := currentUser(c)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return c.Render(http.StatusOK, "info", nil)
}

func EditInfo(c echo.Context) error {
	user := currentUser(c)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return c.Render(http.StatusOK, "edit", nil)
}

func DoEditInfo(c echo.Context) error {
	user := currentUser(c)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return c.JSON(http.StatusOK, nil)
}

func ChangePW(c echo.Context) error {
	user := currentUser(c)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return c.Render(http.StatusOK, "changepw", nil)
}

func DoChangePW(c echo.Context) error {
	user := currentUser(c)
	if user == nil {
		return c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
	return c.JSON(http.StatusOK, nil)
}
