package main

import (
    "net/http"
	_ "eden/routers"
    _ "eden/models"
)

func main() {
	//run
	http.ListenAndServe("localhost:3000", routers.Routers)
}
