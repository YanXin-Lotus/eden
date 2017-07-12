package routers

import (
	"net/http"
	//"eden/config"
	"eden/controllers"
	"eden/models"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

var (
	Routers *echo.Echo
)

func init() {

	Routers = echo.New()

	Routers.Renderer = controllers.T

	//------------
	// Middleware
	//------------
	Routers.Use(middleware.Logger())
	Routers.Use(middleware.Static("/statics"))
	//Routers.Use(middleware.JWT([]byte(config.Config.JwtAuthKey)))

	//article routers
	Routers.GET("/", controllers.Index)
	Routers.GET("/art/:id", controllers.Article)
	Routers.GET("/art/:id/edit", controllers.EditArticle)
	Routers.POST("/art/:id/edit", controllers.DoEditArticle)
	Routers.GET("/page/:id", controllers.Pagination)
	Routers.GET("/cat/:cat", controllers.Category)

	//account routers
	Routers.GET("/login", controllers.Login)
	Routers.POST("/login", controllers.DoLogin)
	Routers.GET("/signout", controllers.Signout)
	Routers.GET("/register", controllers.Register)
	Routers.POST("/register", controllers.DoRegister)
	Routers.GET("/info", controllers.Info)
	Routers.GET("/editinfo", controllers.EditInfo)
	Routers.POST("/editinfo", controllers.DoEditInfo)
	Routers.GET("/changepw", controllers.ChangePW)
	Routers.POST("/changePW", controllers.DoChangePW)

	//other routers
	Routers.GET("/about", controllers.About)
	Routers.GET("/friendship", controllers.Friendship)

	// Initalize
	Admin := admin.New(&qor.Config{DB: models.DB})
	Admin.SetAuth(&controllers.Auth{})
	Admin.SetSiteName("eden admin interface")

	// Create resources from GORM-backend model
	Admin.AddResource(&models.User{}, &admin.Config{Menu: []string{"用户管理"}})
	Admin.AddResource(&models.Article{}, &admin.Config{Menu: []string{"文章管理"}})

	//todo: access control
	// amount to /admin, so visit `/admin` to view the admin interface
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	Routers.Any("/admin/*", echo.WrapHandler(mux))
}
