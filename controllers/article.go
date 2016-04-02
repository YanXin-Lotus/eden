package controllers

import (
	"net/http"

	"github.com/gocraft/web"
)

func (c *Context) Index(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "index", nil)
}

func (c *Context) Pagination(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.JSON(rw, http.StatusAccepted, "Pagination")
}

func (c *Context) Category(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "category", nil)
}

func (c *Context) Article(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "detail", nil)
}

func (c *Context) EditArticle(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "edit", nil)
}

func (c *Context) DoEditArticle(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	return
}
