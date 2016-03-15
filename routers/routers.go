package routers

import (
	"eden/controllers"
	"io"
	"text/template"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
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
	// Debug mode
	Routers.Debug()

	//------------
	// Middleware
	//------------

	// Logger
	Routers.Use(mw.Logger())

	// Recover
	Routers.Use(mw.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	Routers.SetRenderer(t)

	//article routers
	Routers.Get("/", controllers.Index)
	Routers.Get("/page/:id", controllers.Pagination)
	Routers.Get("/cat/:cat", controllers.Category)

	//account routers
	Routers.Get("/login", controllers.Login)

	//other routers
	Routers.Get("/about", controllers.About)
	Routers.Get("/friendship", controllers.Friendship)
}
