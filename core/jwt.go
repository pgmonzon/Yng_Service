package core


import (
    "encoding/json"
    "net/http"
    "time"
    "log"
    "strings"

    "github.com/pgmonzon/Yng_Servicios/models"
    "github.com/pgmonzon/Yng_Servicios/cfg"

    "github.com/dgrijalva/jwt-go"
    //"gopkg.in/mgo.v2/bson"
)

func CrearToken(usuario models.Usuario) (interface{}) {
    // Crea el token
    secreto := []byte(cfg.Secreto)
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": "yangeeapp@gmail.com",
			"exp": time.Now().Add(time.Hour + 1).Unix(), //Expira en 1 hora
			"id": usuario.ID,
		})
    tokenString, _ := token.SignedString(secreto)
    token_json := map[string]string{"token": tokenString}
    return token_json
}

func ArmarToken(r *http.Request) (models.Token) { //En caso de que se quisiesen guardar los tokens, se hace directo de esta funcion
    var token models.Token
    //user := context.Get(r, "Bearer")
    authorization := r.Header["Authorization"][0]
    token_sin_parsear := strings.Fields(authorization) // authorization es "Bearer eyJhbG.eyJle.1Jav", aca separamos las 2 palabras
    log.Println(token_sin_parsear[1])

    tjson, _ := json.Marshal(token_sin_parsear[1].(*jwt.Token))
    json.Unmarshal(tjson, &token)
    return token
}

func ExtraerClaim(r *http.Request, rclaim string) (interface{}) { //Recibe el claim que estas buscando en forma de string, por ejempo "iss" o "id" y devuelve su valor en forma de interface{}
    // Esta funcion sirve para cualquier momento en el que tengas un TOKEN de algun user y quieras leer algun claim en especifico
    token := ArmarToken(r)
    tokenmap, _ := token.Claims.(map[string]interface{})
    return tokenmap[rclaim]
}
