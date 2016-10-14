package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Todo struct to todo
type Roles struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Nombre    string        `json:"rol"`
	Activo    bool          `json:"activo"`
	Borrado   bool          `json:"borrado"`
}

//Todo struct to todo
type Permisos struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Nombre    string        `json:"permiso"`
	Activo    bool          `json:"activo"`
	Borrado   bool          `json:"borrado"`
}

//Todo struct to todo
type RP struct {
	ID             bson.ObjectId `bson:"_id" json:"id"`
	IDRol          string        `json:"rol"`
	IDPermisos     []string      `json:"permisos"`
}
