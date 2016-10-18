package models

import (
  "time"

	"gopkg.in/mgo.v2/bson"
)

type Ejemplo struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Nombre    string        `json:"nombre"`
	Importe   int           `json:"importe"`
  Fecha     time.Time     `json:"fecha"`
	Activo    bool          `json:"activo"`
	Borrado   bool          `json:"borrado"`
}
