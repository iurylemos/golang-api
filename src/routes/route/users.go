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
}
