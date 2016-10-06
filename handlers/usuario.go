package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"

	//"github.com/gorilla/mux"
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
/*
//TodoShow handler to show all todos
func TodoShow(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var todo models.Todo
	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["todoID"]) != true {
		core.JSONError(w, r, start, "bad entry for id", http.StatusBadRequest)
		return
	}
	todoID := bson.ObjectIdHex(vars["todoID"])
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("todos")
	collection.Find(bson.M{"_id": todoID}).One(&todo)
	if todo.ID == "" {
		core.JSONError(w, r, start, "todo not found", http.StatusNotFound)
	} else {
		response, err := json.MarshalIndent(todo, "", "    ")
		if err != nil {
			panic(err)
		}
		core.JSONResponse(w, r, start, response, http.StatusOK)
	}
}
*/
// TodoAdd handler to add new todo
func AgregarUsuario(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuario models.Usuario
	json.NewDecoder(r.Body).Decode(&usuario)
	if usuario.User == "" {
		core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
		return
	}
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
/*
//TodoUpdate handler to update a previous todo
func TodoUpdate(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var todo models.Todo
	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["todoID"]) != true {
		core.JSONError(w, r, start, "bad entry for id", http.StatusBadRequest)
		return
	}
	json.NewDecoder(r.Body).Decode(&todo)
	if todo.Name == "" {
		core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
		return
	}
	todoID := bson.ObjectIdHex(vars["todoID"])
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("todos")
	err := collection.Update(bson.M{"_id": todoID},
		bson.M{"$set": bson.M{"name": todo.Name, "completed": todo.Completed}})
	if err != nil {
		core.JSONError(w, r, start, "Could not find Todo "+string(todoID.Hex())+" to update", http.StatusNotFound)
		return
	}
	core.JSONResponse(w, r, start, []byte{}, http.StatusNoContent)
}

//TodoDelete handler to delete a todo
func TodoDelete(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	vars := mux.Vars(r)
	todoID := bson.ObjectIdHex(vars["todoID"])
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("todos")
	err := collection.Remove(bson.M{"_id": todoID})
	if err != nil {
		core.JSONError(w, r, start, "Could not find Todo "+string(todoID.Hex())+" to delete", http.StatusNotFound)
		return
	}
	core.JSONResponse(w, r, start, []byte{}, http.StatusNoContent)
}

//TodoSearchName handler Todo by Name
func TodoSearchName(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var todo []models.Todo
	vars := mux.Vars(r)
	todoName := vars["todoName"]
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("todos")
	err := collection.Find(bson.M{"name": bson.M{"$regex": todoName}}).All(&todo)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search todo name", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(todo, "", "    ")
	if err != nil {
		panic(err)
	}
	if string(response) == "null" {
		core.JSONError(w, r, start, "Could not find any Todo containing "+todoName, http.StatusNotFound)
		return
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}

//TodoSearchStatus search todo by status (completed, not completed)
func TodoSearchStatus(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var todo []models.Todo
	vars := mux.Vars(r)
	todoStatus := vars["status"]
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("todos")
	if todoStatus == "true" {
		err := collection.Find(bson.M{"completed": bson.M{"$eq": true}}).All(&todo)
		if err != nil {
			core.JSONError(w, r, start, "Failed to search todo name", http.StatusInternalServerError)
			return
		}
		response, err := json.MarshalIndent(todo, "", "    ")
		if err != nil {
			panic(err)
		}
		if string(response) == "null" {
			core.JSONError(w, r, start, "Could not find any Todo containing status "+todoStatus, http.StatusNotFound)
			return
		}
		core.JSONResponse(w, r, start, response, http.StatusOK)
	} else if todoStatus == "false" {
		err := collection.Find(bson.M{"completed": bson.M{"$eq": false}}).All(&todo)
		if err != nil {
			core.JSONError(w, r, start, "Failed to search todo name", http.StatusInternalServerError)
			return
		}
		response, err := json.MarshalIndent(todo, "", "    ")
		if err != nil {
			panic(err)
		}
		if string(response) == "null" {
			core.JSONError(w, r, start, "Could not find any Todo containing status "+todoStatus, http.StatusNotFound)
			return
		}
		core.JSONResponse(w, r, start, response, http.StatusOK)
	} else {
		core.JSONError(w, r, start, "bad request, must be true or false, not "+todoStatus, http.StatusNotFound)
	}
}
*/
