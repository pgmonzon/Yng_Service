package core

// NOTA: Hay objetos de bson guardados en string. Tal vez es preferible ser consistente y guardar todos como objetos de bson


import (
    "net/http"
    "log"

    "github.com/pgmonzon/Yng_Servicios/models"

    "gopkg.in/mgo.v2/bson"
)

func ChequearPermisos(r *http.Request) (bool) {
  // esta funcion se encarga de responder SI o NO a la pregunta "¿tiene este usuario permisos para ejecutar lo que me esta pidiendo"
  // ESTA FUNCIÓN TODAVÍA NO ESTÁ LISTA, es un Work In Progress.
  id := extraerClaim(r, "id")
  if id == ""{
    return false  //Esto sería un error más que falta de permisos (no existe el campo id en el token o es un token invalido). Hay que buscar la forma de manejar estos errores
  }
  log.Println("LOGGING EXITOSO DE: ", id)
  //A PARTIR DE ACA, ESTOY DEBUGGEANDO Y PROBANDO FUNCIONES
  rol_user := extraerRolDelUsuario(id.(string)) // La tengo que convertir a string porque me devolvieron una interface{}
  a := extraerPermisosDelRol(rol_user)
  extraerNombresDePermisos(a)
  return true
}

func extraerRolDelUsuario(id string) (string) {
  //que rol tiene la id que nos pasan???
  var usuario []models.Usuario
	if bson.IsObjectIdHex(id) != true { // Un poco de sanity.
    log.Printf("FATAL ERROR: Id usuario invalida.")
		return "no" //Podria devolver la ID de un usuario especial (una especie de muñeco sin permisos)
	}
  id_bson := bson.ObjectIdHex(id)
	session := Session.Copy()
	defer session.Close()
	collection := session.DB(Dbname).C("usuarios")
  err := collection.Find(bson.M{"_id": id_bson}).All(&usuario)
	if err != nil {
		log.Printf("FATAL ERROR: Id invalida. Lo cual significa que /login esta creando tokens con IDs rotas")
		return "no"
	}
	return usuario[0].Rol
}

func extraerPermisosDelRol(id string) (models.RP){
  //le das una ID de rol a esta funcion, y te devuelve los permisos que tiene ese Rol (los devuelve en un array)
  var rp []models.RP
  if bson.IsObjectIdHex(id) != true { // Un poco de sanity.
    log.Printf("FATAL ERROR: Id rol invalida.")
    return rp[0]
  }
  session := Session.Copy()
  defer session.Close()
  collection := session.DB(Dbname).C("rp")
  err := collection.Find(bson.M{"idrol": id}).All(&rp)
  if err != nil {
    log.Printf("FATAL ERROR: Id invalida. Lo cual significa que /login esta creando tokens con IDs rotas")
    return rp[0]
  }
  return rp[0] //esto no es ideal, es temporal
}

func extraerNombresDePermisos(i models.RP) {
  //NOTA: Esto funciona solo si hay 1 permiso por cada rol. Tengo que buscar documentacion de cómo funcionan los slices y hacer un append por cada permiso.

  /*for _, v := range id {
  }*/
  id := i.IDPermisos[0]
  var permisos []models.Permisos
  session := Session.Copy()
  defer session.Close()
  id_bson := bson.ObjectIdHex(id)
  collection := session.DB(Dbname).C("permisos")
  collection.Find(bson.M{"_id": id_bson}).All(&permisos)
  log.Println("Permisos que tiene el usuario: ", permisos[0].Nombre)

}
