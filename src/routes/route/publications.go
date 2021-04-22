package route

import (
	"api-nos-golang/src/controllers/controller_publications"
	"net/http"
)

var routePublications = []Route{
	{
		URI:                    "/publicacoes",
		Method:                 http.MethodGet,
		Function:               controller_publications.FindPublications,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/publicacoes",
		Method:                 http.MethodPost,
		Function:               controller_publications.CreatePublication,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/publicacoes/{id}",
		Method:                 http.MethodPut,
		Function:               controller_publications.FindPublications,
		RequiredAuthentication: true,
	},
	{
		URI:                    "/publicacoes/{id}",
		Method:                 http.MethodDelete,
		Function:               controller_publications.DeletePublication,
		RequiredAuthentication: true,
	},
}
