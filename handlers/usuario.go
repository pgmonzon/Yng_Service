package handlers

import (
	"encoding/json"
	"crypto/md5"
	"net/http"
	"time"
  "encoding/binary"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//TodoIndex handler to route index
func IndexUsuario(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuarios []models.UsuarioHash
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
	var usuariosimple models.Usuario
	json.NewDecoder(r.Body).Decode(&usuariosimple)
	if usuariosimple.User == "" { //TODO: Esto a futuro debe chequear si el usuario tiene menos de 4 letras o si la contraseña es incorrecta
		core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
		return
	}
	md5pass := md5.New()
	md5pass.Write([]byte(usuariosimple.Pass))
	usuariosimple.Pass = binary.BigEndian.Uint64(md5pass.Sum(nil))
  var usuario models.UsuarioHash

	objID := bson.NewObjectId()
	usuario.ID = objID
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	err := collection.Insert(usuario)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(usuario.ID.Hex()))
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
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
	err := collection.Find(bson.M{"user": usuarioDB.User, "pass": usuarioDB.Pass}).All(&usuario)
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

func UserLogin(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuario []models.Usuario
	var usuarioDB models.Usuario
	json.NewDecoder(r.Body).Decode(&usuarioDB)
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	err := collection.Find(bson.M{"user": usuarioDB.User, "pass": usuarioDB.Pass}).All(&usuario)
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
