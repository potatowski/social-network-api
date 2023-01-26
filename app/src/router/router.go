package router

import (
	"social-api/src/router/routers"

	"github.com/gorilla/mux"
)

// Generate will return a router with configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routers.Configurate(r)
}
