package handlers

import (
  "encoding/json"
  "time"
  "net/http"
  "log"
  "fmt"

	"github.com/pgmonzon/Yng_Servicios/models"
  "github.com/pgmonzon/Yng_Servicios/core"

  "gopkg.in/mgo.v2/bson"
)

const indexPage = `
<h1>Login</h1>
<form method="post" action="/login">
    <label for="name">User name</label>
    <input type="text" id="name" name="name">
    <label for="password">Password</label>
    <input type="password" id="password" name="password">
    <button type="submit">Login</button>
</form>
`
const internalPage = `
<h1>Internal</h1>
<hr>
<small>Hola !!</small>
<form method="post" action="/logout">
<button type="submit">Logout</button>
</form>
`

func IndexLogin(w http.ResponseWriter, r *http.Request) {
  start := time.Now()
  session := core.Session.Copy()
	defer session.Close()
  //w.Header().Set("Content-Type", "application/json; charset=utf-8")
  log.Printf("%s\t%s\t%s\t%s\t%d",
    r.RemoteAddr,
    r.Method,
    r.RequestURI,
    r.Proto,
    time.Since(start),
  )
  fmt.Fprintf(w, indexPage)
}

/*func IndexLogued(w http.ResponseWriter, r *http.Request) {
  start := time.Now()
  session := core.Session.Copy()
  defer session.Close()
  log.Printf("%s\t%s\t%s\t%s\t%d",
    r.RemoteAddr,
    r.Method,
    r.RequestURI,
    r.Proto,
    time.Since(start),
  )
	//name := w.Header().Set("Usuario", "guest")
	//if name == "guest" {
		// .. check credentials ..
		//setSession(name, response)
		//core.JSONError(w, r, start, r.Body, http.StatusInternalServerError)
	//}
	//http.Redirect(response, request, redirectTarget, 302)
  core.JSONResponse(w, r, start, cuerpo, http.StatusInternalServerError)
}*/

func IndexLogued(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuario models.Usuario
  var usuarioDB models.Usuario //juntar estos 2
	json.NewDecoder(r.Body).Decode(&usuario)
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")

  err2 := collection.Find(bson.M{"user": usuario.User}).All(&usuarioDB)
  if err2 != nil {
    core.JSONError(w, r, start, "Failed to search todo name", http.StatusInternalServerError)
    return
  }
  b, err2 := json.Marshal(usuarioDB)

  if usuario.User == "guest" {
    core.JSONResponse(w, r, start, b, http.StatusOK)
    return
  }
  if usuario.User != usuarioDB.User {
    core.JSONError(w, r, start, "QUE HACES QUERIDO!!!", http.StatusBadRequest)
    return
  }
	err := collection.Insert(usuario)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(usuario.ID.Hex()))
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
}
