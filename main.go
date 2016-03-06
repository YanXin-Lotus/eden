package main

import (
    //"net/http"
    "eden/routers"
    //"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main()  {
	// Middleware
	routers.Routers.Use(mw.Logger())
	routers.Routers.Use(mw.Recover())
    routers.Routers.Run(":2333")
}