package main

import (
	"log"
	"net/http"

	"github.com/pgmonzon/Yng_Servicios/routers"
	"github.com/pgmonzon/Yng_Servicios/core"
)

func main() {
	log.Printf("Yangee API Service v0.2")
	router := routers.NewRouter()
	defer core.Session.Close()
	log.Fatal(http.ListenAndServe(":3113", router))
}
