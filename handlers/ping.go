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
	session.Ping()
	defer session.Close()
	//respuesta, _ := json.Marshal("{Esta capa no tiene seguridad}")
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta mierda
	core.JSONError(w, r, start, "sin seguridad", http.StatusOK)
}

func SecuredPingHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	if (!core.ChequearPermisos(r, "SecuredPing")){
		core.JSONError(w, r, start, "Este usuario no tiene permisos o hubo un error procesando tu request. Se ha contactado a un administrador.", http.StatusInternalServerError)
		return
	}
	session := core.Session.Copy()
	defer session.Close()
	respuesta, _ := json.MarshalIndent("Estas autenticado.", "", "    ") //Las respuestas siempre tienen que ser en JSON
	core.JSONResponse(w, r, start, respuesta, http.StatusOK)
}