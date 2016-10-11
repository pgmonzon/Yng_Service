package main


import (
    "fmt"
    "time"

    "github.com/dgrijalva/jwt-go"
)

const (
    mySigningKey = "firechrome"
)

func main() {
    createdToken, err := ExampleNew([]byte(mySigningKey))
    if err != nil {
        fmt.Println("Creating token failed")
    }
    ExampleParse(createdToken, mySigningKey)
}

func ExampleNew(mySigningKey []byte) (string, error) {
    // Create the token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"iss": "adrian.diasdacostalima@gmail.com",
			"exp": time.Now().Add(time.Hour + 1).Unix(),
		})
    // Sign and get the complete encoded token as a string
    tokenString, err := token.SignedString(mySigningKey)
		fmt.Println(token.SignedString(mySigningKey))
    return tokenString, err
}

func ExampleParse(myToken string, myKey string) {
    token, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(myKey), nil
    })

    if err == nil && token.Valid {
        fmt.Println("Your token is valid.  I like your style.")
    } else {
        fmt.Println("This token is terrible!  I cannot accept this.")
    }
}

/*func main() {
  // Create a new token object, specifying signing method and the claims
  // you would like it to contain.
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "foo": "bar",
      "nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
  })

  // Sign and get the complete encoded token as a string using the secret
	hmacSampleSecret := ""
  tokenString, err := token.SignedString(hmacSampleSecret)

  fmt.Println(tokenString, err)
}
*/
