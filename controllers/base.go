package controllers

import (
	"html/template"
	"net/http"

	"github.com/labstack/echo"
	"github.com/unrolled/render"
    "github.com/labstack/echo/engine/standard"
)

var Render *render.Render

func Prepare() {
	return
}

func About(c echo.Context) error {
	return Render.HTML(c.Response().(*standard.Response).ResponseWriter, http.StatusOK, "about", nil)
}


func Friendship(c echo.Context) error {
	return Render.HTML(c.Response().(*standard.Response).ResponseWriter, http.StatusOK, "friendship", nil)

}

func CurrentUser() {
	return
}

func init() {
	Render = render.New(render.Options{
		Directory:  "public/views",
		Extensions: []string{".html"},
		Funcs:      []template.FuncMap{},
	})
}
