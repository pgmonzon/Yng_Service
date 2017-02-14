package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Roles struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Nombre    string        `json:"rol"`
	Activo    bool          `json:"activo"`
	Borrado   bool          `json:"borrado"`
}

type Permisos struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Nombre    string        `json:"nombre"`
	Activo    bool          `json:"activo"`
	Borrado   bool          `json:"borrado"`
	Link	  string	`json:"link"`
}

type RP struct {
	ID             bson.ObjectId	`bson:"_id" json:"id"`
	IDRol          bson.ObjectId	`json:"rol"`
	IDPermisos     []bson.ObjectId	`json:"permisos"`
}

type Menus struct {
	ID		bson.ObjectId	`bson:"_id" json:"id"`
	Desc		string		`json:"desc"`
	IDPadre		bson.ObjectId	`json:"padre"`
	EsMenu		bool		`json:"esmenu"`
	Url		string		`json:"url"`
}
