package controllers

import (
	"eden/models"
	"net/http"

	"github.com/labstack/echo"
)

type BaseController struct {
	*echo.Context
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

func (c BaseController) Message(msg string) {
	return c.HTML(http.StatusFound, html)
}
