package routers

import (
	"net/http"
	"social-api/src/controllers"
)

var routersUser = []Route{
	{
		URI:                  "/user",
		Method:               http.MethodPost,
		Function:             controllers.CreateUser,
		RequireAuthorization: false,
	},
	{
		URI:                  "/users",
		Method:               http.MethodGet,
		Function:             controllers.SearchUsers,
		RequireAuthorization: false,
	},
	{
		URI:                  "/user/{userId}",
		Method:               http.MethodGet,
		Function:             controllers.SearchUserById,
		RequireAuthorization: false,
	},
	{
		URI:                  "/user/{userId}",
		Method:               http.MethodPatch,
		Function:             controllers.UpdateUser,
		RequireAuthorization: false,
	},
	{
		URI:                  "/user/{userId}",
		Method:               http.MethodDelete,
		Function:             controllers.DeleteUser,
		RequireAuthorization: false,
	},
}
