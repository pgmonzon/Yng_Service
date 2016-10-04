package routers

import (
	"log"
	"net/http"
	//"encoding/json"

	"github.com/pgmonzon/Yng_Servicios/handlers"

  "github.com/auth0/go-jwt-middleware"
  "github.com/codegangsta/negroni"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

//NotFound responses to routes not defined
func NotFound(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s\t%s\t%s\t%s\t%d\t%d\t%d",
		r.RemoteAddr,
		r.Method,
		r.RequestURI,
		r.Proto,
		http.StatusNotFound,
		0,
		0,
	)
	w.WriteHeader(http.StatusNotFound)
}

//NewRouter creates the router
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(false)

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("secreto"), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})
	//Todo
	r.HandleFunc("/ping", handlers.PingHandler).Methods("GET")
	r.Handle("/secured/ping", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(handlers.SecuredPingHandler)),
  ))



	r.HandleFunc("/api/todos", handlers.TodoIndex).Methods("GET")
	r.HandleFunc("/api/todos/{todoID}", handlers.TodoShow).Methods("GET")
	r.HandleFunc("/api/todos", handlers.TodoAdd).Methods("POST")
	r.HandleFunc("/api/todos/{todoID}", handlers.TodoUpdate).Methods("PUT")
	r.HandleFunc("/api/todos/{todoID}", handlers.TodoDelete).Methods("DELETE")
	r.HandleFunc("/api/todos/search/byname/{todoName}", handlers.TodoSearchName).Methods("GET")
	r.HandleFunc("/api/todos/search/bystatus/{status}", handlers.TodoSearchStatus).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(NotFound)
	return r
}
