package usuarios

import (
	"encoding/json"
	"net/http"
	"time"
	"log"

	"github.com/pgmonzon/Yng_Servicios/models"
	"github.com/pgmonzon/Yng_Servicios/core"
	"github.com/pgmonzon/Yng_Servicios/cfg"

	//"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// Registrar usuarios
func Registrar(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var usuario models.Usuario
	var lista_usuarios []models.Usuario
	var usuario_crudo models.UsuarioCrudo
	json.NewDecoder(r.Body).Decode(&usuario_crudo) //Nota: NewDecoder por alguna razon solo funciona 1 vez por handler
	if (usuario_crudo.Nombre == "" || usuario_crudo.Pwd == "" || usuario_crudo.Email == "") { //TODO: Esto a futuro debe chequear si el usuario tiene menos de 4 letras o si la contraseña es incorrecta
		core.JSONError(w, r, start, "Usuario, contraseña o email invalidas", http.StatusBadRequest)
		return
	}

	usuario.PassMD = core.HashearMD5(usuario_crudo.Pwd)
	usuario.User = usuario_crudo.Nombre
	usuario.Email = usuario_crudo.Email
	objID := bson.NewObjectId()

	usuario.Activado = false
	usuario.ID = objID
	usuario.Rol = bson.ObjectIdHex(cfg.GuestRol) // en cada creacion de usuario, se les asigna un rol que va a ser GUEST

	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")

	err := collection.Find(bson.M{"user": usuario.User}).All(&lista_usuarios)
	if err != nil {
		core.JSONError(w, r, start, "La base de datos esta caida", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(lista_usuarios)
	if err != nil {
		panic(err)
	}
	if string(response) != "null" {
		core.JSONError(w, r, start, "El usuario especificado ya existe", http.StatusNotFound)
		return
	}

	err = collection.Insert(usuario)
	if err != nil {
		core.JSONError(w, r, start, "Failed to insert user", http.StatusInternalServerError)
		return
	}

	log.Printf("CREANDO USUARIO: %s MD5: %d", usuario.User, usuario.PassMD) //Notese que la password es la md5
	core.EnviarMailDeVerificacion(usuario.Email, usuario.ID)
	response, _ = json.Marshal("Usuario creado satisfactoriamente")
	core.JSONResponse(w, r, start, response, http.StatusCreated)
}
