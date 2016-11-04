package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

//Todo struct to todo
type Usuario struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	User      string        `json:"user"`
	Email     string        `json:"email"`
	PassMD    int32         `json:"md5"`
	Activado  bool		`json:"activado"`
	Codigo    string	`json:"codigo"`
	Creacion  time.Time	`json:"creacion"`
	Rol	  bson.ObjectId `json:"rol"`
}

type UsuarioCrudo struct {
	Nombre		string		`json:"user"`
	Pwd		string		`json:"password"`
	Email		string		`json:"email"`
}

type UsuarioCodigo struct {
	Codigo	string	`json:"codigo"`
}
