package handlers

import (
  "encoding/json"
  "time"
  "net/http"
  "log"
  "fmt"

	"github.com/pgmonzon/Yng_Servicios/models"
  "github.com/pgmonzon/Yng_Servicios/core"
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
	json.NewDecoder(r.Body).Decode(&usuario)
	if usuario.User != "guest" {
		core.JSONError(w, r, start, "Felicidades, me estas pasando algo", http.StatusBadRequest)
		return
	}
  if usuario.User == "guest" {
    core.JSONError(w, r, start, "QUE HACES QUERIDO!!!", http.StatusBadRequest)
  }
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	err := collection.Insert(usuario)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert user", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Location", r.URL.Path+"/"+string(usuario.ID.Hex()))
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
}

/*
func internalPageHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, internalPage)
}
*/
