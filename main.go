package main

import (
	"log"
	"net/http"

	"github.com/pgmonzon/Yng_Servicios/routers"
	"github.com/pgmonzon/Yng_Servicios/core"
)

func main() {
	log.Printf("Yangee API Service v 0.1")
	router := routers.NewRouter() // this func is in router.go
	defer core.Session.Close() // related to Session in db.go
	log.Fatal(http.ListenAndServe(":3113", router))
}
