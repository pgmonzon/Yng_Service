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

/*
type Menus struct {
	ID					bson.ObjectId		`bson:"_id" json:"id"`
	Desc				string					`json:"desc"`
	IDPadre			bson.ObjectId		`json:"padre"`
	EsMenu			bool						`json:"esmenu"`
	Url					string					`json:"url"`
}
*/

func CrearMenu(w http.ResponseWriter, r *http.Request) {
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
	collection := session.DB(core.Dbname).C("menus")
	err := collection.Insert(menu)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert menu", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(menu.ID.Hex()))
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
}

func AjustarMenu(w http.ResponseWriter, r *http.Request) {
	//NOTA: Por el momento actualiza TODAS las variables del menu. Se puede hacer más robusto para que actualice sólo las variables que fueron modificadas
	start := time.Now()
	var menu models.Menus
	json.NewDecoder(r.Body).Decode(&menu)

	if menu.IDPadre == "" {
		core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
		return
	}
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("menus")
	err := collection.Update(bson.M{"_id": menu.ID},
		bson.M{"$set": bson.M{"desc": menu.Desc, "padre": menu.IDPadre, "esmenu": menu.EsMenu, "url": menu.Url}})
	if err != nil {
		core.JSONError(w, r, start, "error ajustando menu, func AjustarMenu()", http.StatusNotFound)
		return
	}
	core.JSONResponse(w, r, start, []byte{}, http.StatusNoContent)
}

func ConseguirMenuesPadres(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var menus []models.Menus
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("menus")
	collection.Find(bson.M{}).All(&menus)
	response, err := json.MarshalIndent(menus, "", "    ")
	if err != nil {
		panic(err)
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}

func ConseguirMenuEspecifico() {

}

func BorrarMenu() {

}
