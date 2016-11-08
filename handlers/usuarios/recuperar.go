package usuarios

/*
 * Todas las funciones necesarias para la recuperacion de contraseñas
 */

 import (
 	//"encoding/json"
 	"net/http"
 	"time"
 	"log"

 	//"github.com/pgmonzon/Yng_Servicios/models"
 	"github.com/pgmonzon/Yng_Servicios/core"
 	//"github.com/pgmonzon/Yng_Servicios/cfg"

 	//"github.com/gorilla/mux"
 	//"gopkg.in/mgo.v2/bson"
 )

func EnviarMailConCodigo(w http.ResponseWriter, r *http.Request) {
	//Recibe un mail, le decimos cual es su usuario y le damos un codigo para que pueda crear una nueva password
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var lista_usuarios []models.Usuario
	var usuario models.UsuarioCrudo
  json.NewDecoder(r.Body).Decode(&usuario)
  nuevoCodigo := core.CrearCodigoDeVerificacion(6)

  session := core.Session.Copy()
  defer session.Close()
  collection := session.DB(core.Dbname).C("usuarios")

  err := collection.Find(bson.M{"email": usuario.Email}).All(&lista_usuarios)
  if err != nil {
    core.JSONError(w, r, start, "El mail no existe", http.StatusInternalServerError)
    return
  }

  //falta updatear el codigo yyyyyyyyyyyyyyyyyy mandar el mail

  core.JSONError(w, r, start, "nais", http.StatusOK)
}
