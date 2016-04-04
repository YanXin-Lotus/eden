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
    c.Data["list"] = list
    c.Data["user"] = c.CurrentUser
	c.HTML(rw, http.StatusOK, "index", c.Data)
}

func (c *Context) Pagination(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
    page := c.ParseParam2Int(req, "page")
    sort := c.ParseParam2Int(req, "sort")
    list, err := services.ArtList(page ,sort)
    if err != nil {
        c.NotFound(rw, req, next)
    }
    c.Data["list"] = list
    c.Data["user"] = c.CurrentUser
	c.HTML(rw, http.StatusOK, "pagination", c.Data)
}

func (c *Context) Category(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
    page := c.ParseParam2Int(req, "page")
    sort := c.ParseParam2Int(req, "sort")
    list, err := services.ArtList(page, sort)
    if err != nil {
        c.NotFound(rw, req, next)
    }
	c.Data["list"] = list
    c.Data["user"] = c.CurrentUser
	c.HTML(rw, http.StatusOK, "pagination", c.Data)
}

func (c *Context) Article(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
    id := c.ParseParam2Uint64(req, "id")
    art, err := services.ShowArt(id)
    if err != nil {
        c.NotFound(rw, req, next)
    }
	c.Data["art"] = art
    c.Data["user"] = c.CurrentUser
	c.HTML(rw, http.StatusOK, "pagination", c.Data)
}

func (c *Context) EditArticle(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HTML(rw, http.StatusOK, "edit", nil)
}

func (c *Context) DoEditArticle(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
    id := c.ParseParam2Uint64(req, "id")
	err := services.UpdateArt(id, nil)
    if err != nil {
        c.NotFound(rw, req, next)
    }
    c.HTML(rw, http.StatusOK, "detail", nil)
}
