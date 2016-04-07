package routers

import (
	"github.com/gocraft/web"
    "eden/controllers"
)

var (
	Routers *web.Router
)

func init() {

	Routers = web.New(controllers.Context{})

	//------------
	// Middleware
	//------------
    Routers.Middleware(web.LoggerMiddleware).Middleware(web.ShowErrorsMiddleware)    

	//article routers
	Routers.Get("/", (*controllers.Context).Index)
	Routers.Get("/art/:id", (*controllers.Context).Article)
	Routers.Get("/art/:id/edit", (*controllers.Context).EditArticle)
	Routers.Post("/art/:id/edit", (*controllers.Context).DoEditArticle)
	Routers.Get("/page/:id", (*controllers.Context).Pagination)
	Routers.Get("/cat/:cat", (*controllers.Context).Category)

	//account routers
	Routers.Get("/login", (*controllers.Context).Login)
	Routers.Post("/login", (*controllers.Context).DoLogin)
	Routers.Get("/signout", (*controllers.Context).Signout)
	Routers.Get("/register", (*controllers.Context).Register)
	Routers.Post("/register", (*controllers.Context).DoRegister)
	Routers.Get("/info", (*controllers.Context).Info)
	Routers.Get("/editinfo", (*controllers.Context).EditInfo)
	Routers.Post("/editinfo", (*controllers.Context).DoEditInfo)
	Routers.Get("/changepw", (*controllers.Context).ChangePW)
	Routers.Post("/changePW", (*controllers.Context).DoChangePW)

	//other routers
	Routers.Get("/about", (*controllers.Context).About)
	Routers.Get("/friendship", (*controllers.Context).Friendship)
}
