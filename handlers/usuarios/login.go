package usuarios

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


func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*") //Porfavor no olvidarse de borrar esta porqueria
	if (ChequearSocialLogin(w, r)) {
		return
	}
	start := time.Now()
	var lista_usuarios []models.Usuario
	var usuario_crudo models.UsuarioCrudo
	json.NewDecoder(r.Body).Decode(&usuario_crudo)
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")

	err := collection.Find(bson.M{"user": usuario_crudo.Nombre, "passmd": core.HashearMD5(usuario_crudo.Pwd)}).All(&lista_usuarios)
	if err != nil {
		core.JSONError(w, r, start, "Failed to search user name", http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(lista_usuarios)
	if err != nil {
		panic(err)
	}
	if string(response) == "null" {
		core.JSONError(w, r, start, "Usuario o clave incorrectas", http.StatusBadRequest)
		return
	}
	token := core.CrearToken(lista_usuarios[0])
	response, _ = json.Marshal(token)
	log.Println(lista_usuarios[0].User, " se ha logueado satisfactoriamente")
	core.JSONResponse(w, r, start, response, http.StatusCreated)
}

func ChequearSocialLogin(w http.ResponseWriter, r *http.Request) (bool) {
	//Chequeamos si el login viene desde google o facebook
	// Por el momento chequeamos si es facebook
	var usuario_facebook models.UsuarioFacebook
	json.NewDecoder(r.Body).Decode(&usuario_facebook)

	log.Println(usuario_facebook)
	if (len(usuario_facebook.ID) == 0 ){
		log.Println("NO ES LOGIN DE FACEBOOK. PROBANDO CON LOGIN NORMAL")
		return false
	}
	FacebookLogin(w, r, usuario_facebook)
	return true
}

func FacebookLogin(w http.ResponseWriter, r *http.Request, usuario_facebook models.UsuarioFacebook) {
	//Login de facebook, el chequeo dio positivo asi que tenemos que loguear O registrar en caso de que no exista el usuario.
	var lista_usuarios_facebook []models.Usuario
	session := core.Session.Copy()
	defer session.Close()
	collection := session.DB(core.Dbname).C("usuarios")
	collection.Find(bson.M{"facebook.id": usuario_facebook.ID} ).All(&lista_usuarios_facebook)
	if (len(lista_usuarios_facebook) == 0){
		log.Println("NO SE ENCONTRÓ REGISTRO DE FACEBOOK. CREANDO...")
		RegistrarFacebook(w, r, usuario_facebook)
	}
	//log.Println(lista_usuarios_facebook[0], "ES LA LISTA DE USUARIOS CON ID: ", usuario_facebook.ID)
}
