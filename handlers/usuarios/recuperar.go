package usuarios

/*
 * Todas las funciones necesarias para la recuperacion de contraseñas
 */

 import (
 	"encoding/json"
 	"net/http"
 	"time"
 	"log"

 	"github.com/pgmonzon/Yng_Servicios/models"
 	"github.com/pgmonzon/Yng_Servicios/core"
 	//"github.com/pgmonzon/Yng_Servicios/cfg"

 	//"github.com/gorilla/mux"
 	"gopkg.in/mgo.v2/bson"
 )

func RecuperarPassword(w http.ResponseWriter, r *http.Request) {
	//Recibe un mail, le decimos cual es su usuario y le damos un codigo para que pueda crear una nueva password
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var lista_usuarios []models.Usuario
	var usuario models.UsuarioRecuperar
  json.NewDecoder(r.Body).Decode(&usuario)
  log.Println(usuario, usuario.Email)

  session := core.Session.Copy()
  defer session.Close()
  collection := session.DB(core.Dbname).C("usuarios")

  err := collection.Find(bson.M{"email": usuario.Email}).All(&lista_usuarios)
  if err != nil {
    core.JSONError(w, r, start, "El mail no existe", http.StatusInternalServerError)
    return
  }
  if (len(lista_usuarios) == 0){
    core.JSONError(w, r, start, "El mail no existe en la base de datos", http.StatusInternalServerError)
    return
  }

  core.EnviarMailDeVerificacion(lista_usuarios[0].Email, lista_usuarios[0].ID)

  core.JSONError(w, r, start, "nais", http.StatusOK)
}


func RecibirCodigoDeRecuperacion(w http.ResponseWriter, r *http.Request) {
	//Chequeamos si el codigo que nos estan mandando es el mismo que el guardado en la base de datos.
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var lista_usuarios []models.Usuario
	var usuario models.UsuarioRecuperar
	json.NewDecoder(r.Body).Decode(&usuario)

	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")

  err := collection.Find(bson.M{"email": usuario.Email}).All(&lista_usuarios)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search user name", http.StatusInternalServerError)
		return
	}
  if (len(lista_usuarios) == 0){
    core.JSONError(w, r, start, "El mail no existe en la base de datos", http.StatusInternalServerError)
    return
  }

	if (usuario.Codigo == lista_usuarios[0].Codigo) {
    log.Println("Usuario ", lista_usuarios[0].User, "recuperando contraseña")
		token := core.CrearToken(lista_usuarios[0])
    response, _ := json.Marshal(token)
    core.JSONResponse(w, r, start, response, http.StatusOK)
	}	else {
    core.JSONError(w, r, start, "El código no es correcto", http.StatusInternalServerError)
	}
}
