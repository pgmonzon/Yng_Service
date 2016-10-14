package core


import (
    "net/http"
    "log"

    "github.com/pgmonzon/Yng_Servicios/models"

    "gopkg.in/mgo.v2/bson"
)

func ChequearPermisos(r *http.Request) (bool) {
  // esta funcion se encarga de responder SI o NO a la pregunta "¿tiene este usuario permisos para ejecutar lo que me esta pidiendo"
  id := extraerClaim(r, "id")
  if id == nil{
    return false  //Esto sería un error más que falta de permisos (no existe el campo id en el token o es un token invalido). Hay que buscar la forma de manejar estos errores
  }
  log.Println("LOGGING EXITOSO DE: ", id)
  extraerRolDelUsuario(id)
  return true
}

func extraerRolDelUsuario(id interface{}) (string) {
  //que rol tiene la id que nos pasan???
  var usuario []models.Usuario
	if bson.IsObjectIdHex(id.(string)) != true { // Un poco de sanity.
    log.Printf("FATAL ERROR: Id invalida.")
		return
	}
  id_bson := bson.ObjectIdHex(id.(string))
	session := Session.Copy()
	defer session.Close()
	collection := session.DB(Dbname).C("usuarios")
  err := collection.Find(bson.M{"_id": id_bson}).All(&usuario)
	if err != nil {
		log.Printf("FATAL ERROR: Id invalida. Lo cual significa que /login esta creando tokens con IDs rotas")
		return
	}
	return usuario[0].Rol
}
