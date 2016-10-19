package routers

import (
	"log"
	"net/http"
	//"encoding/json"

	"github.com/pgmonzon/Yng_Servicios/handlers"
	"github.com/pgmonzon/Yng_Servicios/cfg"

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
	secret := cfg.Secreto

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
			/*Extractor: jwtmiddleware.FromFirst(jwtmiddleware.FromAuthHeader,
		                                     jwtmiddleware.FromParameter("auth_code")),*/
		})
	//Todo
	r.HandleFunc("/ping", handlers.PingHandler).Methods("GET")
	r.Handle("/secured/ping", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(handlers.SecuredPingHandler)),
  ))

  //Usuarios, por ahora solo tiene 2 funciones, mostrar usuarios y agregar usuarios

  //r.HandleFunc("/api/usuarios", handlers.IndexUsuario).Methods("GET")
	r.Handle("/api/usuarios", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(handlers.IndexUsuario)),
	))
	r.HandleFunc("/api/usuarios/login", handlers.UserLogin).Methods("POST")
	r.HandleFunc("/api/usuarios/register", handlers.AgregarUsuario).Methods("POST")
	//r.HandleFunc("/api/usuarios/search/byname/{User}", handlers.UserSearchName).Methods("GET")
	//r.HandleFunc("/api/usuarios/search/byname/{User}", handlers.UserSearchNameJSON).Methods("POST")

	//############			RBAC			##############
	r.HandleFunc("/api/roles", handlers.ListarRoles).Methods("GET")
	r.HandleFunc("/api/roles", handlers.AgregarRol).Methods("POST")
	r.HandleFunc("/api/permisos", handlers.ListarPermisos).Methods("GET")
	r.HandleFunc("/api/permisos", handlers.AgregarPermiso).Methods("POST")
	//r.HandleFunc("/api/rp", handlers.ListarRP).Methods("GET")
	r.HandleFunc("/api/rp", handlers.AgregarRP).Methods("POST")

	//############		Ejemplo		##############
	r.HandleFunc("/api/ejemplos", handlers.EjemploIndex).Methods("GET")
	r.HandleFunc("/api/ejemplos", handlers.AgregarEjemplo).Methods("POST")
	//r.HandleFunc("/api/ejemplos/{ejemploID}", handlers.ModificarEjemplo).Methods("PUT")
	r.Handle("/api/ejemplos/{ejemploID}", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(handlers.ModificarEjemplo)),
	)).Methods("PUT")

  //Ejemplo de todos
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
