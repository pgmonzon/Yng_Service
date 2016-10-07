package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"log"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//TodoIndex handler to route index
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

// TodoAdd handler to add new todo
func AgregarUsuario(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuario models.Usuario
	var usuariodos []models.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)
	if (usuario.User == "" || usuario.Pass == "") { //TODO: Esto a futuro debe chequear si el usuario tiene menos de 4 letras o si la contraseña es incorrecta
		core.JSONError(w, r, start, "Usuario o contraseña invalidas", http.StatusBadRequest)
		return
	}
	usuario.PassMD = core.HashearMD5(usuario.Pass)

	objID := bson.NewObjectId()
	usuario.ID = objID
	session := core.Session.Copy()
	defer session.Close()
	//log.Printf("La ID es: %s")
	collection := session.DB(core.Dbname).C("usuarios")

	err := collection.Find(bson.M{"user": usuario.User}).All(&usuariodos)
	if err != nil {
		core.JSONError(w, r, start, "La base de datos esta caida", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(usuariodos, "", "    ")
	if err != nil {
		panic(err)
	}
	if string(response) != "null" {
		core.JSONError(w, r, start, "El usuario especificado ya existe", http.StatusNotFound)
		return
	}

	err = collection.Insert(usuario)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert user", http.StatusInternalServerError)
		return
	}
	log.Printf("CREANDO USUARIO: %s PASSWORD: %d", usuario.User, usuario.PassMD)
	w.Header().Set("Location", r.URL.Path+"/"+string(usuario.ID.Hex()))
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
}


func UserLogin(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuario []models.Usuario
	var usuarioDB models.Usuario
	json.NewDecoder(r.Body).Decode(&usuarioDB)
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	err := collection.Find(bson.M{"user": usuarioDB.User, "passmd": core.HashearMD5(usuarioDB.Pass)}).All(&usuario)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search user name", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(usuario, "", "    ")
	if err != nil {
		panic(err)
	}
	if string(response) == "null" {
		core.JSONError(w, r, start, "Usuario o clave incorrectas", http.StatusNotFound)
		return
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}

func UserSearchNameJSON(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuario []models.Usuario
	var usuarioDB models.Usuario
	json.NewDecoder(r.Body).Decode(&usuarioDB)
	vars := mux.Vars(r)
	User := vars["User"]
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	err := collection.Find(bson.M{"user": usuarioDB.User, "md5": core.HashearMD5(usuarioDB.Pass)}).All(&usuario)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search user name", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(usuario, "", "    ")
	if err != nil {
		panic(err)
	}
	if string(response) == "null" {
		core.JSONError(w, r, start, "Could not find any User containing "+usuarioDB.User+User, http.StatusNotFound)
		return
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}

func UserSearchName(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuario []models.Usuario
	vars := mux.Vars(r)
	User := vars["User"]
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	err := collection.Find(bson.M{"user": "santi"}).All(&usuario)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search user name", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(usuario, "", "    ")
	if err != nil {
		panic(err)
	}
	if string(response) == "null" {
		core.JSONError(w, r, start, "Could not find any User containing "+User, http.StatusNotFound)
		return
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}
