package usuarios

import (
	"encoding/json"
	"net/http"
	"time"
	"log"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"
	//"github.com/pgmonzon/Yng_Servicios/cfg"

	//"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)


func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var lista_usuarios []models.Usuario
	var usuario_crudo models.UsuarioCrudo
	json.NewDecoder(r.Body).Decode(&usuario_crudo)
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")

	err := collection.Find(bson.M{"user": usuario_crudo.Nombre, "passmd": core.HashearMD5(usuario_crudo.Pwd)}).All(&lista_usuarios)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search user name", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(lista_usuarios)
	if err != nil {
		panic(err)
	}
	if string(response) == "null" {
		core.JSONError(w, r, start, "Usuario o clave incorrectas", http.StatusCreated)
		return
	}
	token := core.CrearToken(lista_usuarios[0])
	response, _ = json.Marshal(token)
	log.Println(lista_usuarios[0].User, " se ha logueado satisfactoriamente")
	core.JSONResponse(w, r, start, response, http.StatusCreated)
}
