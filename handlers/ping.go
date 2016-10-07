package handlers

import (
	"net/http"
	"time"

	"github.com/pgmonzon/Yng_Servicios/core"
)


func PingHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	//var todos []models.Todo
	session := core.Session.Copy()
	defer session.Close()
	core.JSONResponse(w, r, start, []byte("Todo en orden"), http.StatusOK)
}

func SecuredPingHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	//var todos []models.Todo
	session := core.Session.Copy()
	defer session.Close()
	core.JSONResponse(w, r, start, []byte("Felicitaciones! Estas autenticado"), http.StatusOK)
}
