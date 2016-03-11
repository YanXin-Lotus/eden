package controllers

import (
	"eden/models"
	"net/http"

	"github.com/labstack/echo"
)

type BaseController struct {
	*echo.Context
	Data map[string]interface{}
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

func (c BaseController) CurrentUser() (u *models.User) {
	return
}

func (c BaseController) Message() {
	return c.Render(http.StatusOK, "views/index.html", "test")
}
