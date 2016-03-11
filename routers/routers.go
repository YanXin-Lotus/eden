package routers

import (
	"eden/controllers"
	"io"
	"text/template"

	"github.com/labstack/echo"
)

var Routers *echo.Echo

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init() {
	Routers = echo.New()
	Routers.Get("/", controllers.MainController.Index)
	Routers.Get("/page/:id", controllers.MainController.Pagination)
	Routers.Get("/cat/:cat", controllers.MainController.Category)

	Routers.Get("/message", controllers.BaseController.Message)
}
