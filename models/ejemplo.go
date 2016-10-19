package models

import (
  "time"

	"gopkg.in/mgo.v2/bson"
)

type Ejemplo struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Nombre    string        `json:"nombre"`
	Importe   int           `json:"importe"`
	Activo    bool          `json:"activo"`
	Borrado   bool          `json:"borrado"`
}

type EjemploAuditoria struct {
  ID        bson.ObjectId `bson:"_id" json:"id"`
  Nombre    string        `json:"nombre"`
  Importe   int           `json:"importe"`
  Activo    bool          `json:"activo"`
  Borrado   bool          `json:"borrado"`
  Fecha     time.Time     `json:"fecha"`
  IDReq     bson.ObjectId `json:"idejemplo"`
  IDUser    bson.ObjectId `json:"iduser"`
}
