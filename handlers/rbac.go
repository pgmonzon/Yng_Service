package handlers

// TODO: Mucho copypaste, se tiene que poder simplificar.

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"

	//"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func ListarRoles(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var roles []models.Roles
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("roles")
	collection.Find(bson.M{}).All(&roles)
	response, err := json.MarshalIndent(roles, "", "    ")
	if err != nil {
		panic(err)
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}

func ListarPermisos(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var permisos []models.Permisos
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("permisos")
	collection.Find(bson.M{}).All(&permisos)
	response, err := json.MarshalIndent(permisos, "", "    ")
	if err != nil {
		panic(err)
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}




func AgregarRol(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var rol models.Roles
	json.NewDecoder(r.Body).Decode(&rol)
	if rol.Nombre == "" {
		core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
		return
	}
	objID := bson.NewObjectId()
	rol.ID = objID
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("roles")
	err := collection.Insert(rol)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert rol", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(rol.ID.Hex()))
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
}

func AgregarPermiso(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var permiso models.Permisos
	json.NewDecoder(r.Body).Decode(&permiso)
	if permiso.Nombre == "" {
		core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
		return
	}
	objID := bson.NewObjectId()
	permiso.ID = objID
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("permisos")
	err := collection.Insert(permiso)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert permiso", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(permiso.ID.Hex()))
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
}

func AgregarRP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var rp models.RP
	json.NewDecoder(r.Body).Decode(&rp)
	if rp.IDRol == "" {
		core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
		return
	}
	objID := bson.NewObjectId()
	rp.ID = objID
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("rp")
	err := collection.Insert(rp)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert rp", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(rp.ID.Hex()))
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
}
