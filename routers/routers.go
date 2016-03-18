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
    
    //static files
    Routers.Use(mw.Static("public"))

	// Logger
	Routers.Use(mw.Logger())

	// Recover
	Routers.Use(mw.Recover())

	//article routers
	Routers.Get("/", controllers.Index)
	Routers.Get("/art/:id", controllers.Article)
	Routers.Get("/art/:id/edit", controllers.EditArticle)
	Routers.Post("/art/:id/edit", controllers.DoEditArticle)
	Routers.Get("/page/:id", controllers.Pagination)
	Routers.Get("/cat/:cat", controllers.Category)

	//account routers
	Routers.Get("/login", controllers.Login)
	Routers.Post("/login", controllers.DoLogin)
	Routers.Get("/signout", controllers.Signout)
	Routers.Get("/register", controllers.Register)
	Routers.Post("/register", controllers.DoRegister)
	Routers.Get("/info", controllers.Info)
	Routers.Get("/editinfo", controllers.EditInfo)
	Routers.Post("/editinfo", controllers.DoEdit)
	Routers.Get("/changepw", controllers.ChangePW)
	Routers.Post("/changePW", controllers.DoChange)

	//other routers
	Routers.Get("/about", controllers.About)
	Routers.Get("/friendship", controllers.Friendship)
}
