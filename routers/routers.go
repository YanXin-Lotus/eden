package routers

import (
	"eden/controllers"

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
	t := &Template {
    	templates: template.Must(template.ParseGlob("views/*.html")),
	}
	Routers.SetRenderer(t)
	Routers.Get("/", controllers.MainController.Index)
	Routers.Get("/page/:id", controllers.MainController.Pagination)
	Routers.Get("/cat/:cat", controllers.MainController.Category)

	Routers.Get("/message", controllers.BaseController.Message)
}
