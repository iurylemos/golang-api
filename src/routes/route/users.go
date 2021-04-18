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
		RequiredAuthentication: false,
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
		RequiredAuthentication: false,
	},
	{
		URI:                    "/usuarios/{id}",
		Method:                 http.MethodPut,
		Function:               controller_users.UpdateUSer,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/usuarios/{id}",
		Method:                 http.MethodDelete,
		Function:               controller_users.DeleteUser,
		RequiredAuthentication: false,
	},
}
