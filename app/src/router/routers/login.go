package routers

import (
	"social-api/src/controllers"
)

var loginRoute = Route{
	URI:                  "/login",
	Method:               "POST",
	Function:             controllers.Login,
	RequireAuthorization: false,
}
