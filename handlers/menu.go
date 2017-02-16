package handlers

// TODO: Mucho copypaste, se tiene que poder simplificar.

import (
	"encoding/json"
	"net/http"
	"time"
	//"log"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"

	//"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func AgregarMenu(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var menu models.Menus
	json.NewDecoder(r.Body).Decode(&menu)
	if menu.Desc == "" {
		core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
		return
	}
	objID := bson.NewObjectId()
	menu.ID = objID
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("menues")
	err := collection.Insert(menu)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert rol", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(menu.ID.Hex()))
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
}
