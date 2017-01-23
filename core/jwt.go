package core


import (
    //"encoding/json"
    "net/http"
    "time"
    "fmt"
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

func ParsearToken(r *http.Request) (*jwt.Token) {
    /* Parseamos los tokens.
     * Leemos los headers de *http.Request y buscamos por Authorization que es lo que vamos a usar.
     * Una vez que tenemos nuestro header, lo parseamos usando funciones de jwt de dgrijalva.
     * Por último devolvemos el token en forma de map. Un token parseado se ve de la siguiente forma:
     * ANTES  : eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODQ3NzAx... etc etc
     * DESPUES: map[exp:1.484770114e+09 id:586ea6684158607ce356dd40 iss:yangeeapp@gmail.com]
    */
    authorization := r.Header["Authorization"][0]
    token_sin_parsear := strings.Fields(authorization) // authorization es "Bearer eyJhbG.eyJle.1Jav", aca separamos las 2 palabras

    token, _ := jwt.Parse(token_sin_parsear[1], func(token *jwt.Token) (interface{}, error) {

    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
        return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
    }

    return []byte(cfg.Secreto), nil
    })

    return token
}

func ExtraerClaim(r *http.Request, rclaim string) (interface{}) {

    // Recibe el claim que estas buscando en forma de string, por ejempo "iss" o "id" y devuelve su valor en forma de interface{}
    // Esta funcion sirve para cualquier momento en el que tengas un TOKEN de algun user y quieras leer algun claim en especifico
    token := ParsearToken(r)
    token_mapeado := token.Claims.(jwt.MapClaims)
    return token_mapeado[rclaim]
}
