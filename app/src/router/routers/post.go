package routers

import (
	"net/http"
	"social-api/src/controller"
)

var routersPost = []Route{
	{
		URI:                  "/post",
		Method:               http.MethodPost,
		Function:             controller.CreatePost,
		RequireAuthorization: true,
	},
	{
		URI:                  "/posts",
		Method:               http.MethodGet,
		Function:             controller.SearchPosts,
		RequireAuthorization: true,
	},
	{
		URI:                  "/post/{uuid}",
		Method:               http.MethodGet,
		Function:             controller.SearchPostByUuid,
		RequireAuthorization: true,
	},
	{
		URI:                  "/post/{uuid}",
		Method:               http.MethodPatch,
		Function:             controller.UpdatePost,
		RequireAuthorization: true,
	},
	{
		URI:                  "/post/{uuid}",
		Method:               http.MethodDelete,
		Function:             controller.DeletePost,
		RequireAuthorization: true,
	},
	{
		URI:                  "/user/{userID}/posts",
		Method:               http.MethodGet,
		Function:             controller.SearchPostsByUser,
		RequireAuthorization: true,
	},
	{
		URI:                  "/post/{uuid}/like",
		Method:               http.MethodPost,
		Function:             controller.LikePost,
		RequireAuthorization: true,
	},
	{
		URI:                  "/post/{uuid}/like",
		Method:               http.MethodDelete,
		Function:             controller.UnlikePost,
		RequireAuthorization: true,
	},
	{
		URI:                  "/post/{uuid}/dislike",
		Method:               http.MethodPost,
		Function:             controller.DislikePost,
		RequireAuthorization: true,
	},
}
