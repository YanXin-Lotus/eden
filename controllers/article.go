package controllers

import (
    "eden/services"
	"net/http"

	"github.com/gocraft/web"
)

func (c *Context) Index(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
    list, err := services.ArtList(0, 0)
    if err != nil {
        c.NotFound(rw, req, next)
    }
	c.HTML(rw, http.StatusOK, "index", list)
}

func (c *Context) Pagination(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
    list, err := services.ArtList(0 ,0)
    if err != nil {
        c.NotFound(rw, req, next)
    }
	c.HTML(rw, http.StatusOK, "pagination", list)
}

func (c *Context) Category(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
    list, err := services.ArtList(0, 0)
    if err != nil {
        c.NotFound(rw, req, next)
    }
	c.HTML(rw, http.StatusOK, "category", list)
}

func (c *Context) Article(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
    art, err := services.ShowArt(0)
    if err != nil {
        c.NotFound(rw, req, next)
    }
	c.HTML(rw, http.StatusOK, "detail", art)
}

func (c *Context) EditArticle(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "edit", nil)
}

func (c *Context) DoEditArticle(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	err := services.UpdateArt(0, nil)
    if err != nil {
        c.NotFound(rw, req, next)
    }
    c.HTML(rw, http.StatusOK, "detail", nil)
}
