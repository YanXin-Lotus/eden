package main

import (
    "net/http"
	"eden/routers"
)

func main() {
	//init routers
	routers.Init()

	//run
	http.ListenAndServe("localhost:3000", routers.Routers)
}
