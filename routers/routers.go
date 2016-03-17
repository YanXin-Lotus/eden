package routers

import (
	"eden/controllers"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

var (
	Routers *echo.Echo
)

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
