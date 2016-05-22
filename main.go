package main

import (
	"eden/routers"
    _ "eden/models"
	
	"github.com/labstack/echo/engine/standard"
)

func main() {
	//run
	routers.Routers.Run(standard.New(":2333"))
}
