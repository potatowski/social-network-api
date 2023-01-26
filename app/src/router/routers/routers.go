package routers

import "net/http"

// Route is a struct to representation of all api routes
type Route struct {
	URI                  string
	Method               string
	Function             func(http.ResponseWriter, *http.Request)
	RequireAuthorization bool
}
