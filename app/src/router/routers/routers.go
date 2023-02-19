package routers

import (
	"net/http"
	"social-api/src/middleware"

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
	routes = append(routes, loginRoute)
	routes = append(routes, routersPost...)

	for _, route := range routes {
		if route.RequireAuthorization {
			r.HandleFunc(route.URI,
				middleware.Logger(middleware.Auth(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middleware.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
