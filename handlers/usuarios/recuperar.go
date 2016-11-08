package usuarios

/*
 * Todos los handlers necesarios para la recuperacion de contraseñas
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

func RecuperarPassword(w http.ResponseWriter, r *http.Request) {
	//Recibe un mail, le decimos cual es su usuario y le damos un codigo para que pueda crear una nueva password
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	//var lista_usuarios []models.Usuario
	//var usuario models.UsuarioCrudo
  log.Println("good")
  core.JSONError(w, r, start, "nais", http.StatusOK)
}
