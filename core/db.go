package core

import "gopkg.in/mgo.v2"

//Host establece el path de la base de datos
var Host = "localhost:27017"

//dbName establece el nombre de la base de datos
var Dbname = "yangee"

//Session Establish the main session
var Session = NewConnection()

// NewConnection create connection to DB
func NewConnection() *mgo.Session {
	session, err := mgo.Dial(Host)

	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
