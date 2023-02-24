package routers

import (
	"social-api/src/controller"
)

var loginRoute = Route{
	URI:                  "/login",
	Method:               "POST",
	Function:             controller.Login,
	RequireAuthorization: false,
}
