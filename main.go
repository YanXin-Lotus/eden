package main

import (
	_ "eden/models"
	"eden/routers"
)

func main() {
	//run
	routers.Routers.Run(routers.Routers)
}
