package routers

import (
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
    
    //static files
    Routers.Use(mw.Static("public"))

	// Logger
	Routers.Use(mw.Logger())

	// Recover
	Routers.Use(mw.Recover())

	//article routers
	Routers.Get("/", index())
	Routers.Get("/art/:id", article())
	Routers.Get("/art/:id/edit", editArticle())
	Routers.Post("/art/:id/edit", doEditArticle())
	Routers.Get("/page/:id", pagination())
	Routers.Get("/cat/:cat", category())

	//account routers
	Routers.Get("/login", login())
	Routers.Post("/login", doLogin())
	Routers.Get("/signout", signout())
	Routers.Get("/register", register())
	Routers.Post("/register", doRegister())
	Routers.Get("/info", info())
	Routers.Get("/editinfo", editInfo())
	Routers.Post("/editinfo", doEditInfo())
	Routers.Get("/changepw", changePW())
	Routers.Post("/changePW", doChangePW())

	//other routers
	Routers.Get("/about", about())
	Routers.Get("/friendship", friendship())
}
