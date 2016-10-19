package core

// NOTA: Hay objetos de bson guardados en string. Tal vez es preferible ser consistente y guardar todos como objetos de bson
import (
    "net/http"
    "log"
    "time"
    //"errors"

    "github.com/pgmonzon/Yng_Servicios/models"
    "gopkg.in/mgo.v2"
    //"github.com/pgmonzon/Yng_Servicios/cfg"
    "gopkg.in/mgo.v2/bson"
)

func Auditar(r *http.Request, session *mgo.Session, id_ejemplo bson.ObjectId) (bool){
  var ejemplos []models.Ejemplo
  var auditoria models.EjemploAuditoria
  id_user := extraerClaim(r, "id")
  log.Println("la id del usuario es la siguiente, string o interface?", id_user)
  if id_user == ""{
    return false  //Esto sería un error más que falta de permisos (no existe el campo id en el token o es un token invalido). Hay que buscar la forma de manejar estos errores
  }
  collection := session.DB(Dbname).C("ejemplos")
  collection.Find(bson.M{"_id": id_ejemplo}).All(&ejemplos)
  if (len(ejemplos) == 0) {
    log.Printf("ERROR: Id invalida. Lo cual significa que /login esta creando tokens con IDs rotas o la db esta corrupta")
    return false
  }
  auditoria.ID = bson.NewObjectId()
  auditoria.Nombre = ejemplos[0].Nombre
  auditoria.Importe = ejemplos[0].Importe
  auditoria.Activo = ejemplos[0].Activo
  auditoria.Borrado = ejemplos[0].Borrado
  auditoria.Fecha = time.Now()
  auditoria.IDReq = ejemplos[0].ID
  auditoria.IDUser = bson.ObjectIdHex(id_user.(string))

  collection = session.DB(Dbname).C("ejemploauditoria")
  err := collection.Insert(auditoria)
  if err != nil {
    log.Println("ooooooooopssssssss")
    return false
  }
  return true


  //log.Println(m, " -+- ", id)

}

/*start := time.Now()
var todo models.Todo
json.NewDecoder(r.Body).Decode(&todo)
if todo.Name == "" {
  core.JSONError(w, r, start, "Incorrect body", http.StatusBadRequest)
  return
}
objID := bson.NewObjectId()
todo.ID = objID
todo.Created = time.Now()
session := core.Session.Copy()
defer session.Close()
collection := session.DB(core.Dbname).C("todos")
err := collection.Insert(todo)
if err != nil {
  core.JSONError(w, r, start, "Failed to insert todo", http.StatusInternalServerError)
  return
}
w.Header().Set("Location", r.URL.Path+"/"+string(todo.ID.Hex()))*/
