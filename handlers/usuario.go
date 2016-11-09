package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"log"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"
	"github.com/pgmonzon/Yng_Servicios/cfg"

	//"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func IndexUsuario(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuarios []models.Usuario
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	collection.Find(bson.M{}).All(&usuarios)
	response, err := json.MarshalIndent(usuarios, "", "    ")
	if err != nil {
		panic(err)
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var usuarioDB []models.Usuario
	var usuario_crudo models.UsuarioCrudo
	json.NewDecoder(r.Body).Decode(&usuario_crudo)
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	err := collection.Find(bson.M{"user": usuario_crudo.Nombre, "passmd": core.HashearMD5(usuario_crudo.Pwd)}).All(&usuarioDB)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search user name", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(usuarioDB, "", "    ")
	if err != nil {
		panic(err)
	}
	if string(response) == "null" {
		core.JSONError(w, r, start, "Usuario o clave incorrectas", http.StatusCreated)
		return
	}
	token := core.CrearToken(usuarioDB[0])
	response, _ = json.Marshal(token)
	log.Println(usuarioDB[0].User)
	core.JSONResponse(w, r, start, response, http.StatusCreated)
}


func VerificarUsuario(w http.ResponseWriter, r *http.Request) {
	//Chequeamos si el codigo que nos estan mandando es el mismo que el guardado en la base de datos.
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var lista_usuarios []models.Usuario
	var codigo models.UsuarioCodigo
	json.NewDecoder(r.Body).Decode(&codigo)

	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	id := core.ExtraerClaim(r, "id")
	id_bson := bson.ObjectIdHex(id.(string))

	err := collection.Find(bson.M{"_id": id_bson}).All(&lista_usuarios)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search user name", http.StatusInternalServerError)
		return
	}
	if (lista_usuarios[0].Activado == true) {
		core.JSONError(w, r, start, "Este usuario ya está activo", http.StatusOK)
		return
	}

	err = collection.Update(bson.M{"_id": id_bson},
		bson.M{"$set": bson.M{"activado": true}})
	response, err := json.Marshal(lista_usuarios)
	if err != nil {
		panic(err)
	}
	if string(response) == "null" {
		core.JSONError(w, r, start, "Usuario o clave incorrectas", http.StatusCreated)
		return
	}
	response, _ = json.Marshal(lista_usuarios[0].Codigo)
	log.Println(lista_usuarios[0].Codigo)
	if (codigo.Codigo == lista_usuarios[0].Codigo) {
		log.Println(codigo.Codigo, " matchea con la base de datos")
	}	else {
		log.Println(codigo.Codigo, " no matchea con", lista_usuarios[0].Codigo)
	}
	core.JSONResponse(w, r, start, response, http.StatusCreated)
}
