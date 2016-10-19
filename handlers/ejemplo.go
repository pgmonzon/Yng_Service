package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//TodoIndex handler to route index
func EjemploIndex(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var ejemplo []models.Ejemplo
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("ejemplos")
	collection.Find(bson.M{}).All(&ejemplo)
	response, err := json.MarshalIndent(ejemplo, "", "    ")
	if err != nil {
		panic(err)
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}

// TodoAdd handler to add new todo
func AgregarEjemplo(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var ejemplo models.Ejemplo
	json.NewDecoder(r.Body).Decode(&ejemplo)
	if ejemplo.Nombre == "" {
		core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
		return
	}
	objID := bson.NewObjectId()
	ejemplo.ID = objID
	//ejemplo.Fecha = time.Now()
	//log.Println(ejemplo.Fecha.Format(time.RFC1123Z)) //Standard RFC 1123Z: "Mon, 02 Jan 2006 15:04:05 -0700")
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("ejemplos")
	err := collection.Insert(ejemplo)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert todo", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(ejemplo.ID.Hex()))
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
}

func ModificarEjemplo(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var ejemplo models.Ejemplo
	vars := mux.Vars(r)
	if bson.IsObjectIdHex(vars["ejemploID"]) != true {
		core.JSONError(w, r, start, "bad entry for id", http.StatusBadRequest)
		return
	}
	json.NewDecoder(r.Body).Decode(&ejemplo)
	if ejemplo.Nombre == "" {
		core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
		return
	}
	ejemploID := bson.ObjectIdHex(vars["ejemploID"])
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("ejemplos")
	err := collection.Update(bson.M{"_id": ejemploID},
		bson.M{"$set": bson.M{"nombre": ejemplo.Nombre, "importe": ejemplo.Importe, "activo": ejemplo.Activo, "borrado": ejemplo.Borrado}})
	if err != nil {
		core.JSONError(w, r, start, "Could not find Todo "+string(ejemploID.Hex())+" to update", http.StatusNotFound)
		return
	}
	if(!core.Auditar(r, session, ejemploID)){
		core.JSONError(w, r, start, "Error interno. Porfavor contactar a un operador", http.StatusInternalServerError)
	}

	core.JSONResponse(w, r, start, []byte{}, http.StatusNoContent)
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
}*/
