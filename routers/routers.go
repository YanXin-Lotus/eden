package routers

import (
    "eden/controllers"
    "github.com/labstack/echo"
)

var Routers *echo.Echo

func init() {
    Routers := echo.New()
    Routers.Get("/", controllers.Index)
}