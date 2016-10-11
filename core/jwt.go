package core


import (
    "time"

    "github.com/dgrijalva/jwt-go"
)

const (
    mySigningKey = "firechrome"
)

func CrearToken(mySigningKey []byte) (string) {
    // Create the token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": "adrian.diasdacostalima@gmail.com",
			"exp": time.Now().Add(time.Hour + 1).Unix(),
		})
    // Sign and get the complete encoded token as a string
    tokenString, _ := token.SignedString(mySigningKey)
    return tokenString
}
