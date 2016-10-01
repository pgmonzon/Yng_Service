package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Todo struct to todo
type Todo struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `json:"name"`
	Aux       string        `json:"aux"`
	Completed bool          `json:"completed"`
	Created   time.Time     `json:"createdon"`
}
