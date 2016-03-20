package main

import (
	"eden/routers"
    "github.com/labstack/echo/engine/standard"
)

func main() {
	//init routers
	routers.Init()

	//run
	routers.Routers.Run(standard.New(":2333"))
}
