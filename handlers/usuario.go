package handlers

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

func IndexUsuario(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	var usuarios []models.Usuario
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	collection.Find(bson.M{}).All(&usuarios)
	response, err := json.MarshalIndent(usuarios, "", "    ")
	if err != nil {
		panic(err)
	}
	core.JSONResponse(w, r, start, response, http.StatusOK)
}

// Registrar usuarios
func AgregarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var usuario models.Usuario			// Se usan dos models porque uno sirve para parsear el r.Body que se pasa en la llamada http
	var usuarioDB []models.Usuario // mientras que el otro se arma para chequear la base de datos (usuarioDB)
	var usuario_crudo models.UsuarioCrudo
	json.NewDecoder(r.Body).Decode(&usuario_crudo) //Nota: NewDecoder por alguna razon solo funciona 1 vez por handler
	if (usuario_crudo.Nombre == "" || usuario_crudo.Pwd == "") { //TODO: Esto a futuro debe chequear si el usuario tiene menos de 4 letras o si la contraseña es incorrecta
		core.JSONError(w, r, start, "Usuario o contraseña invalidas", http.StatusBadRequest)
		return
	}

	usuario.PassMD = core.HashearMD5(usuario_crudo.Pwd)
	usuario.User = usuario_crudo.Nombre
	usuario.Email = usuario_crudo.Email
	objID := bson.NewObjectId()

	usuario.Activado = 0 //Desactivado. Tenemos que mandarle un mail de activacion
	usuario.ID = objID
	usuario.Rol = bson.ObjectIdHex(cfg.GuestRol) // en cada creacion de usuario, se les asigna un rol que va a ser GUEST

	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")

	err := collection.Find(bson.M{"user": usuario.User}).All(&usuarioDB)
	if err != nil {
		core.JSONError(w, r, start, "La base de datos esta caida", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(usuarioDB, "", "    ")
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

	//Ya podemos mandar nuestro mail de verificacion usando core.EnviarMail(vars)
	w.Header().Set("Location", r.URL.Path+"/"+string(usuario.ID.Hex())) //que es esto?
	core.JSONResponse(w, r, start, []byte{}, http.StatusCreated)
}


func UserLogin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	start := time.Now()
	var usuarioDB []models.Usuario
	var usuario_crudo models.UsuarioCrudo
	json.NewDecoder(r.Body).Decode(&usuario_crudo)
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	err := collection.Find(bson.M{"user": usuario_crudo.Nombre, "passmd": core.HashearMD5(usuario_crudo.Pwd)}).All(&usuarioDB)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search user name", http.StatusInternalServerError)
		return
	}
	response, err := json.MarshalIndent(usuarioDB, "", "    ")
	if err != nil {
		panic(err)
	}
	if string(response) == "null" {
		core.JSONError(w, r, start, "Usuario o clave incorrectas", http.StatusCreated)
		return
	}
	token := core.CrearToken(usuarioDB[0])
	response, _ = json.Marshal(token)
	log.Println(usuarioDB[0].User)
	core.JSONResponse(w, r, start, response, http.StatusCreated)
}
