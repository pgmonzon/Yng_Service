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

func Auditar(r *http.Request, session *mgo.Session, id_bson bson.ObjectId, nombre_tabla string) (bool){
  var ejemplos []models.Ejemplo
  var auditoria models.Auditoria
  id_user := extraerClaim(r, "id")
  log.Println("la id del usuario es la siguiente, string o interface?", id_user)
  if id_user == ""{
    return false  //Esto sería un error más que falta de permisos (no existe el campo id en el token o es un token invalido). Hay que buscar la forma de manejar estos errores
  }
  collection := session.DB(Dbname).C(nombre_tabla)
  collection.Find(bson.M{"_id": id_bson}).All(&ejemplos)
  if (len(ejemplos) == 0) {
    log.Printf("ERROR: Id invalida.")
    return false
  }
  auditoria.ID = bson.NewObjectId()
  auditoria.NombreTabla = nombre_tabla
  auditoria.Tabla = ejemplos[0]
  auditoria.Fecha = time.Now()
  //auditoria.IDReq = ejemplos[0].ID
  auditoria.IDUser = bson.ObjectIdHex(id_user.(string))

  collection = session.DB(Dbname).C("auditoria") //Aca se pueden guardar en tablas diferentes segun el tipo de cosa que estemos auditando.
  err := collection.Insert(auditoria)
  if err != nil {
    log.Println("ooooooooopssssssss")
    return false
  }
  return true

}
