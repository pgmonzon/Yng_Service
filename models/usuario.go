package models

import (
	"gopkg.in/mgo.v2/bson"
)

//Todo struct to todo
type Usuario struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	User      string        `json:"user"`
	Pass      string        `json:"password"`
	Email     string        `json:"email"`
	PassMD    int32        `json:"md5"`
}
