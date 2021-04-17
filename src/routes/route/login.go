package route

import (
	"api-nos-golang/src/controllers/controller_login"
	"net/http"
)

var routeLogin = Route{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Function:               controller_login.Login,
	RequiredAuthentication: false,
}
