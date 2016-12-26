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

type Auditoria struct {
  ID        bson.ObjectId `bson:"_id" json:"id"`
  NombreTabla string      `json:"nombretabla"`
  Tabla     interface{}   `json:"tabla"`
  Fecha     time.Time     `json:"fecha"`
  //IDReq     bson.ObjectId `json:"idrequested"`
  IDUser    bson.ObjectId `json:"iduser"`
}
