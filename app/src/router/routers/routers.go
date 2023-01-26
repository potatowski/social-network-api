package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route is a struct to representation of all api routes
type Route struct {
	URI                  string
	Method               string
	Function             func(http.ResponseWriter, *http.Request)
	RequireAuthorization bool
}

// Configurate added all routes in mux router
func Configurate(r *mux.Router) *mux.Router {
	routes := routersUser

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
