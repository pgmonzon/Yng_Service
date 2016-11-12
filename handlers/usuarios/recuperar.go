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

func CambiarContrasena(w http.ResponseWriter, r *http.Request) {
	//En realidad esta función no debería ir acá, pero bueno, por el momento acá se queda.
  //Llamamos a esta funcion con un token y la contraseña nueva y las cambia.
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var usuario models.UsuarioRecuperar
	json.NewDecoder(r.Body).Decode(&usuario)

	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")

  id_usuario := core.ExtraerClaim(r, "id")
  id_bson := bson.ObjectIdHex(id_usuario.(string))
  log.Println(id_bson)

  err := collection.Update(bson.M{"_id": id_bson},
		bson.M{"$set": bson.M{"passmd": core.HashearMD5(usuario.Pwd)}})
	if err != nil {
		core.JSONError(w, r, start, "La base de datos está caida", http.StatusInternalServerError)
		return
	}
  log.Println("Cambiando contraseña del usuario con id ", id_usuario)

  response, _ := json.Marshal("todo ok papa")
  core.JSONResponse(w, r, start, response, http.StatusOK)
}
