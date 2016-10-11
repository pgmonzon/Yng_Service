package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pgmonzon/Yng_Servicios/core"
)


func PingHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	//var todos []models.Todo
	session := core.Session.Copy()
	defer session.Close()
	respuesta, _ := json.MarshalIndent("Esta capa no tiene seguridad", "", "    ")
	core.JSONResponse(w, r, start, respuesta, http.StatusOK)
}

func SecuredPingHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	//var todos []models.Todo
	session := core.Session.Copy()
	defer session.Close()
	respuesta, _ := json.MarshalIndent("Estas autenticado.", "", "    ")
	core.JSONResponse(w, r, start, respuesta, http.StatusOK)
}
