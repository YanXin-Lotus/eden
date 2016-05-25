package routers

import (
	"net/http"
	//"eden/config"
	"eden/models"
    "eden/controllers"
	"github.com/qor/qor"
	"github.com/qor/admin"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

var (
	Routers *echo.Echo
)

func init() {

	Routers = echo.New()
	
	Routers.SetRenderer(controllers.T)

	//------------
	// Middleware
	//------------   
	Routers.Use(middleware.Logger())
	Routers.Use(middleware.Static("/static"))
	//Routers.Use(middleware.JWT([]byte(config.Config.JwtAuthKey)))
	
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
	Routers.Post("/editinfo", controllers.DoEditInfo)
	Routers.Get("/changepw", controllers.ChangePW)
	Routers.Post("/changePW", controllers.DoChangePW)

	//other routers
	Routers.Get("/about", controllers.About)
	Routers.Get("/friendship", controllers.Friendship)

	// Initalize
	Admin := admin.New(&qor.Config{DB: models.DB})
	Admin.SetSiteName("eden admin interface")

	// Create resources from GORM-backend model
	Admin.AddResource(&models.User{},  &admin.Config{Menu: []string{"用户管理"}})
	Admin.AddResource(&models.Article{},  &admin.Config{Menu: []string{"文章管理"}})

	//todo: access control
	// amount to /admin, so visit `/admin` to view the admin interface
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
    Routers.Any("/admin/*", standard.WrapHandler(mux))
}
