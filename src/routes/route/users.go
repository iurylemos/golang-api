package route

import (
	"api-nos-golang/src/controllers/controller_users"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                    "/usuarios",
		Method:                 http.MethodPost,
		Function:               controller_users.CreateUser,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/usuarios",
		Method:                 http.MethodGet,
		Function:               controller_users.FindUsers,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/usuarios/{id}",
		Method:                 http.MethodGet,
		Function:               controller_users.FindUser,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/usuarios/{id}",
		Method:                 http.MethodPut,
		Function:               controller_users.UpdateUSer,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/usuarios/{id}",
		Method:                 http.MethodDelete,
		Function:               controller_users.DeleteUser,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/usuarios/{id}/seguir",
		Method:                 http.MethodPost,
		Function:               controller_users.FollowUser,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/usuarios/{id}/parar-de-seguir",
		Method:                 http.MethodPost,
		Function:               controller_users.UnfollowUser,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/usuarios/{id}/seguidores",
		Method:                 http.MethodGet,
		Function:               controller_users.FindFollowersUser,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/usuarios/{id}/seguindo",
		Method:                 http.MethodGet,
		Function:               controller_users.FindFollowingUser,
		RequiredAuthentication: true,
	},
}
