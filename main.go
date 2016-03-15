package main

import (
	"eden/routers"
	//"github.com/labstack/echo"
)

func main() {
	//init routers
	routers.Init()

	//run
	routers.Routers.Run(":2333")
}
