package core


import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/pgmonzon/Yng_Servicios/models"

	  "github.com/gorilla/context"
    "github.com/dgrijalva/jwt-go"
    //"gopkg.in/mgo.v2/bson"
)

func CrearToken(usuario models.Usuario) (string) {
    // Crea el token
    secreto := []byte("firechrome")
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": "yangeeapp@gmail.com",
			"exp": time.Now().Add(time.Hour + 1).Unix(),
      "id": usuario.ID,
		})
    // Crea una string a partir de la key (el secreto)
    tokenString, _ := token.SignedString(secreto)
    return tokenString
}

func ArmarToken(r *http.Request) (models.Token) { //En caso de que se quisiesen guardar los tokens, se hace directo de esta funcion
    var token models.Token
    user := context.Get(r, "user")
    tjson, _ := json.Marshal(user.(*jwt.Token))
    json.Unmarshal(tjson, &token)
    return token
}
