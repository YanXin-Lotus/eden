package main

import (
	_ "eden/models"
	"eden/routers"
)

func main() {
	e := routers.Routers
	e.Logger.Fatal(e.Start(":8089"))
}
