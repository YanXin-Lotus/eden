package controllers

import (
	"net/http"
    "eden/models"
    "eden/services"

	"github.com/gocraft/web"
)

func (c *Context) Login(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "account/login", nil)
}

func (c *Context) DoLogin(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
    var user models.User
    err := c.Parse2Struct(req, &user)
    if err != nil {
        c.Redirect(rw, req, "/login")
    }
    err = services.Login(&user)
    if err != nil {
        c.JSON(rw, http.StatusForbidden, &retJson{OK: false, Desc: "Error password"})
    }
    c.SetUser(&user, rw)
	c.Redirect(rw, req, "/")
}

func (c *Context) Signout(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	return
}

func (c *Context) Register(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "account/register", nil)
}

func (c *Context) DoRegister(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.Redirect(rw, req, "/")
}

func (c *Context) Info(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "account/info", nil)
}

func (c *Context) EditInfo(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "account/editinfo", nil)
}

func (c *Context) DoEditInfo(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.Redirect(rw, req, "/info")
}

func (c *Context) ChangePW(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "account/changepw", nil)
}

func (c *Context) DoChangePW(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.Redirect(rw, req, "/info")
}
