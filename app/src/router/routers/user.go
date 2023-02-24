package routers

import (
	"net/http"
	"social-api/src/controller"
)

var routersUser = []Route{
	{
		URI:                  "/user",
		Method:               http.MethodPost,
		Function:             controller.CreateUser,
		RequireAuthorization: false,
	},
	{
		URI:                  "/users",
		Method:               http.MethodGet,
		Function:             controller.SearchUsers,
		RequireAuthorization: true,
	},
	{
		URI:                  "/user/{userId}",
		Method:               http.MethodGet,
		Function:             controller.SearchUserById,
		RequireAuthorization: true,
	},
	{
		URI:                  "/user/{userId}",
		Method:               http.MethodPatch,
		Function:             controller.UpdateUser,
		RequireAuthorization: true,
	},
	{
		URI:                  "/user/{userId}",
		Method:               http.MethodDelete,
		Function:             controller.DeleteUser,
		RequireAuthorization: true,
	},
	{
		URI:                  "/user/{userId}/follow",
		Method:               http.MethodPost,
		Function:             controller.FollowUser,
		RequireAuthorization: true,
	},
	{
		URI:                  "/user/{userId}/unfollow",
		Method:               http.MethodPost,
		Function:             controller.UnfollowUser,
		RequireAuthorization: true,
	},
	{
		URI:                  "/user/{userId}/followers",
		Method:               http.MethodGet,
		Function:             controller.SearchUserFollowers,
		RequireAuthorization: true,
	},
	{
		URI:                  "/user/{userId}/following",
		Method:               http.MethodGet,
		Function:             controller.SearchUserFollowing,
		RequireAuthorization: true,
	},
	{
		URI:                  "/user/{userId}/update-password",
		Method:               http.MethodPatch,
		Function:             controller.UpdateUserPassword,
		RequireAuthorization: true,
	},
}
