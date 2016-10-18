package core

// NOTA: Hay objetos de bson guardados en string. Tal vez es preferible ser consistente y guardar todos como objetos de bson
import (
    "net/http"
    "log"

    "github.com/pgmonzon/Yng_Servicios/models"
    "github.com/pgmonzon/Yng_Servicios/cfg"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)

func ChequearPermisos(r *http.Request, permisoBuscado string) (bool) {
  // esta funcion se encarga de responder SI o NO a la pregunta "¿tiene este usuario permisos para ejecutar lo que me esta pidiendo?"
  id := extraerClaim(r, "id")
  session := Session.Copy()
  defer session.Close()
  if id == ""{
    return false  //Esto sería un error más que falta de permisos (no existe el campo id en el token o es un token invalido). Hay que buscar la forma de manejar estos errores
  }
  log.Println("LOGGING EXITOSO DE: ", id)
  //A PARTIR DE ACA, ESTOY DEBUGGEANDO Y PROBANDO FUNCIONES
  user, err := extraerInfoUsuario(id.(string), session) // La tengo que convertir a string porque me devolvieron una interface{}
  if err != true {
    return false
  }
  a := extraerPermisosDelRol(user.Rol, session)
  PermisoID := extraerIDPermiso(permisoBuscado, session).ID
  for _, v := range a.IDPermisos {
    v_bson := bson.ObjectIdHex(v)
    if (v_bson == PermisoID) {
      return true
    }
  }
  //log.Println("Rol del usuario: ",extraerInfoRol(user.Rol, session).Nombre)
  return false
}

func extraerInfoRol(id string, session *mgo.Session) (models.Roles) {
  var rol []models.Roles
  if bson.IsObjectIdHex(id) != true { // Un poco de sanity.
    log.Printf("FATAL ERROR: Id rol invalida.")
    return rol[0]
  }
  id_bson := bson.ObjectIdHex(id)
  //session := Session.Copy()
  //defer session.Close()
  collection := session.DB(Dbname).C("roles")
  err := collection.Find(bson.M{"_id": id_bson}).All(&rol)
  if err != nil {
    log.Printf("FATAL ERROR: Id invalida. Lo cual significa que /login esta creando tokens con IDs rotas")
    return rol[0]
  }
  return rol[0]
}

func extraerInfoUsuario(id string, session *mgo.Session) (models.Usuario, bool) {
  //que rol tiene la id que nos pasan???
  var usuario []models.Usuario
	if bson.IsObjectIdHex(id) != true { // Un poco de sanity.
    log.Printf("FATAL ERROR: Id usuario invalida.")
		return usuario[0], false //Podria devolver la ID de un usuario especial (una especie de muñeco sin permisos)
	}
  id_bson := bson.ObjectIdHex(id)
	//session := Session.Copy()
	//defer session.Close()
	collection := session.DB(Dbname).C("usuarios")
  err := collection.Find(bson.M{"_id": id_bson}).All(&usuario)
	if err != nil {
		log.Printf("FATAL ERROR: Id invalida. Lo cual significa que /login esta creando tokens con IDs rotas")
		return usuario[0], false
	}
  log.Println(usuario[0])
  if usuario[0].Rol == cfg.GuestRol{ //Es guest
    return usuario[0], false
  }
	return usuario[0], true
}

func extraerPermisosDelRol(id string, session *mgo.Session) (models.RP){
  //le das una ID de rol a esta funcion, y te devuelve los permisos que tiene ese Rol (los devuelve en un array)
  var rp []models.RP
  if bson.IsObjectIdHex(id) != true { // Un poco de sanity.
    log.Printf("FATAL ERROR: Id rol invalida.")
    return rp[0]
  }
  //session := Session.Copy()
  //defer session.Close()
  collection := session.DB(Dbname).C("rp")
  err := collection.Find(bson.M{"idrol": id}).All(&rp)
  if err != nil {
    log.Printf("FATAL ERROR: Id invalida. Lo cual significa que /login esta creando tokens con IDs rotas")
    return rp[0]
  }
  return rp[0] //esto no es ideal, es temporal
}

func extraerIDPermiso(permiso string, session *mgo.Session) (models.Permisos) {
  var modelPermisos []models.Permisos
  collection := session.DB(Dbname).C("permisos")
  err := collection.Find(bson.M{"nombre": permiso}).All(&modelPermisos)
  if err != nil {
    log.Printf("FATAL ERROR: Permiso invalido. El permiso buscado no existe.")
    return modelPermisos[0]
  }
  return modelPermisos[0]
}

func estaActivo() (bool){
  return true
}
