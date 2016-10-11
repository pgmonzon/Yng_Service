package core


import (
    "fmt"
    "time"
    "net/http"
    "log"

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

func ResponderToken(w http.ResponseWriter, r *http.Request, start time.Time, response []byte, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	log.Printf("%s\t%s\t%s\t%s\t%d\t%d\t%s",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		code,
		len(response),
		time.Since(start),
	)
	if string(response) != "" {
		w.Write(response)
	}
}
