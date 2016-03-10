package main

import (
	"eden/routers"
	//"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func main() {
	//init routers
	routers.Init()

	// Middleware
	routers.Routers.Use(mw.Logger())
	routers.Routers.Use(mw.Recover())

	//run
	routers.Routers.Run(":2333")
}
