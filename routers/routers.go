package routers

import (
	"eden/controllers"

	"github.com/labstack/echo"
)

var Routers *echo.Echo

func Init() {
	Routers = echo.New()
	Routers.Get("/", controllers.Index)
}
