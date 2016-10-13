package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pgmonzon/Yng_Servicios/core"

	"github.com/dgrijalva/jwt-go/request"
)


func PingHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	//var todos []models.Todo
	session := core.Session.Copy()
	session.Ping()
	defer session.Close()
	respuesta, _ := json.MarshalIndent("Esta capa no tiene seguridad", "", "    ")
	core.JSONResponse(w, r, start, respuesta, http.StatusOK)
}

func SecuredPingHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	//var todos []models.Todo
	session := core.Session.Copy()
	defer session.Close()
	token := jwt.ParseFromRequest(r)
	respuesta, _ := json.MarshalIndent("Estas autenticado.", "", "    ") //Las respuestas siempre tienen que ser en JSON
	core.JSONResponse(w, r, start, respuesta, http.StatusOK)
}
