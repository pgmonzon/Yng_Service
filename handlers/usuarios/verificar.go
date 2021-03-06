package usuarios

/*
 *  Modulo encargado de verificar al usuario luego de la registración.
*/

import (
	"encoding/json"
	"net/http"
	"time"
	"log"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"

	"gopkg.in/mgo.v2/bson"
)

func Verificar(w http.ResponseWriter, r *http.Request) {
	//Verifica el codigo de activacion luego de registrarnos.
	//Esto se hace chequeando si el codigo que nos estan mandando es el mismo que el guardado en la base de datos.
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var lista_usuarios []models.Usuario
	var codigo models.UsuarioCodigo
	json.NewDecoder(r.Body).Decode(&codigo)

	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	id := core.ExtraerClaim(r, "id")
	id_bson := bson.ObjectIdHex(id.(string))

	err := collection.Find(bson.M{"_id": id_bson}).All(&lista_usuarios)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search user name", http.StatusInternalServerError)
		return
	}
	if (lista_usuarios[0].Activado == true) {
		core.JSONError(w, r, start, "Este usuario ya está activo", http.StatusOK)
		return
	}

	if (codigo.Codigo == lista_usuarios[0].Codigo) {
		log.Println(lista_usuarios[0].User, "acaba de activar su cuenta con el codigo", codigo.Codigo, ".")
	} else {
		log.Println(codigo.Codigo, "no matchea con", lista_usuarios[0].Codigo)
		core.JSONError(w, r, start, "Codigo incorrecto", http.StatusNoContent)
		return
	}



	err = collection.Update(bson.M{"_id": id_bson},
		bson.M{"$set": bson.M{"activado": true}})
	response, err := json.Marshal(lista_usuarios) //hace algo esto??
	if err != nil {
		panic(err)
	}
	if string(response) == "null" {
		core.JSONError(w, r, start, "Usuario o clave incorrectas", http.StatusInternalServerError)
		return
	}
	response, _ = json.Marshal(lista_usuarios[0].Activado)
	core.JSONResponse(w, r, start, response, http.StatusCreated)
}
