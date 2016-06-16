package main

import (
	_ "eden/models"
	"eden/routers"

	"github.com/labstack/echo/engine/standard"
)

func main() {
	//run
	routers.Routers.Run(standard.New(":2333"))
}
