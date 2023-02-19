package routers

import (
	"net/http"
	"social-api/src/controllers"
)

var routersPost = []Route{
	{
		URI:                  "/post",
		Method:               http.MethodPost,
		Function:             controllers.CreatePost,
		RequireAuthorization: true,
	},
	{
		URI:                  "/posts",
		Method:               http.MethodGet,
		Function:             controllers.SearchPosts,
		RequireAuthorization: true,
	},
	{
		URI:                  "/post/{uuid}",
		Method:               http.MethodGet,
		Function:             controllers.SearchPostByUuid,
		RequireAuthorization: true,
	},
	{
		URI:                  "/post/{uuid}",
		Method:               http.MethodPatch,
		Function:             controllers.UpdatePost,
		RequireAuthorization: true,
	},
	{
		URI:                  "/post/{uuid}",
		Method:               http.MethodDelete,
		Function:             controllers.DeletePost,
		RequireAuthorization: true,
	},
	{
		URI:                  "/user/{userID}/posts",
		Method:               http.MethodGet,
		Function:             controllers.SearchPostsByUser,
		RequireAuthorization: true,
	},
	{
		URI:                  "/post/{uuid}/like",
		Method:               http.MethodPost,
		Function:             controllers.LikePost,
		RequireAuthorization: true,
	},
}
