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
func HeroesIndex(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var heroes []models.Hero
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("heroes")
	collection.Find(bson.M{}).All(&heroes)
	response, err := json.MarshalIndent(heroes, "", "    ")
	if err != nil {
		panic(err)
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}
