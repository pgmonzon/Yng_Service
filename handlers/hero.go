package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"log"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"

	//"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//TodoIndex handler to route index
func HeroesIndex(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var heroes []models.Hero
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("heroes")
	collection.Find(bson.M{}).All(&heroes)
	//response, err := json.MarshalIndent(heroes, "", "    ")
	response, err := json.Marshal(heroes)
	if err != nil {
		panic(err)
	}
	w.Header().Add("Access-Control-Allow-Origin", "*")
	core.JSONResponse(w, r, start, response, http.StatusOK)
}

func HeroesUpdate(w http.ResponseWriter, r *http.Request) {
	//start := time.Now()
	var heroes models.Hero
	//vars := mux.Vars(r)
	/*if bson.IsObjectIdHex(vars["todoID"]) != true {
		core.JSONError(w, r, start, "bad entry for id", http.StatusBadRequest)
		return
	}*/
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewDecoder(r.Body).Decode(&heroes)
	log.Println(heroes)
	/*if todo.Name == "" {
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
	core.JSONResponse(w, r, start, []byte{}, http.StatusNoContent)*/
}

func HeroesOk(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Add("Access-Control-Allow-Origin", "*")
	core.JSONResponse(w, r, start, []byte{}, http.StatusOK)
}
