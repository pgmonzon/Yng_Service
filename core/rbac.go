package core

// NOTA: Hay objetos de bson guardados en string. Tal vez es preferible ser consistente y guardar todos como objetos de bson
import (
    "net/http"
    "log"
    "errors"

    "github.com/pgmonzon/Yng_Servicios/models"
    "github.com/pgmonzon/Yng_Servicios/cfg"
    "gopkg.in/mgo.v2/bson"
)

func ChequearPermisos(r *http.Request, permisoBuscado string) (bool) {
  // esta funcion se encarga de responder SI o NO a la pregunta "¿tiene este usuario permisos para ejecutar lo que me esta pidiendo?"
  id := ExtraerClaim(r, "id")
  permiso, err := ExtraerInfoPermiso(permisoBuscado)
  if (err != nil) { return false }
  if (!permiso.Activo || permiso.Borrado){
    return false
  }
  if id == ""{
    return false  //Esto sería un error más que falta de permisos (no existe el campo id en el token o es un token invalido). Hay que buscar la forma de manejar estos errores
  }
  user, err := ExtraerInfoUsuario(id.(string)) // La tengo que convertir a string porque me devolvieron una interface{}
  if (err != nil) { return false }
  if user.Rol == cfg.GuestRol{
    return false //Es guest
  }

  RP, err := ExtraerPermisosDelRol(user.Rol)
  if (err != nil) { return false }
  for _, v := range RP.IDPermisos {
    v_bson := bson.ObjectIdHex(v)
    if (v_bson == permiso.ID) {
      log.Println("Acceso permitido de:",user.User,"a:",permisoBuscado)
      return true
    }
  }
  return false
}

func ExtraerInfoRol(id string) (models.Roles, error) {
  var modelRol []models.Roles
  var modelError models.Roles
  if bson.IsObjectIdHex(id) != true { // Un poco de sanity.
    log.Println("ERROR: Id rol invalida.", id)
    return modelError, errors.New("La id del Rol no es un objeto bson")
  }
  id_bson := bson.ObjectIdHex(id)
  session := Session.Copy()
  defer session.Close()
  collection := session.DB(Dbname).C("roles")
  collection.Find(bson.M{"_id": id_bson}).All(&modelRol)
  if (len(modelRol) == 0) {
    log.Println("ERROR: Id invalida. El usuario", id, "tiene un Rol que no existe")
    return modelError, errors.New("El usuario tiene un rol que no existe")
  }
  return modelRol[0], nil
}

func ExtraerInfoUsuario(id string) (models.Usuario, error) {
  //que rol tiene la id que nos pasan???
  var usuario []models.Usuario
  var modelError models.Usuario
	if bson.IsObjectIdHex(id) != true { // Un poco de sanity.
    log.Printf("FATAL ERROR: Id usuario invalida.")
		return modelError, errors.New("Id usuario invalida") //Podria devolver la ID de un usuario especial (una especie de muñeco sin permisos)
	}
  id_bson := bson.ObjectIdHex(id)
	session := Session.Copy()
	defer session.Close()
	collection := session.DB(Dbname).C("usuarios")
  collection.Find(bson.M{"_id": id_bson}).All(&usuario)
	if (len(usuario) == 0) {
		log.Printf("ERROR: Id invalida. Lo cual significa que /login esta creando tokens con IDs rotas o la db esta corrupta")
		return modelError, errors.New("Id usuario invalida")
	}
	return usuario[0], nil
}


func ExtraerPermisosDelRol(id bson.ObjectId) (models.RP, error){
  var rp []models.RP
  var modelError models.RP

  id_string := bson.ObjectId.Hex(id)
  session := Session.Copy()
  defer session.Close()
  collection := session.DB(Dbname).C("rp")
  collection.Find(bson.M{"idrol": id_string}).All(&rp)
  if (len(rp) == 0) {
    return modelError, errors.New("Id no tiene permisos.")
  }
  return rp[0], nil
}

func ExtraerInfoPermiso(permiso string) (models.Permisos, error) {
  //Nota: En caso que sea necesario, se puede hacer un case switch si "permiso string" es una ID o el nombre del permiso
  var modelPermisos []models.Permisos
  var modelError models.Permisos
  session := Session.Copy()
  defer session.Close()
  collection := session.DB(Dbname).C("permisos")
  collection.Find(bson.M{"nombre": permiso}).All(&modelPermisos)
  if (len(modelPermisos) == 0) {
    log.Printf("ERROR: Permiso invalido. El permiso buscado no existe.")
    return modelError, errors.New("El permiso buscado no existe")
  }
  return modelPermisos[0], nil
}

func EstaPermisoActivo(permiso string) (bool){
  per, err := ExtraerInfoPermiso(permiso)
  if err == nil {
    if per.Activo == true{
      return true
    }
  }
  return false
}
