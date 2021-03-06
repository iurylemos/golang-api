package route

import (
	"api-nos-golang/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}

func SettingsRoute(r *mux.Router) *mux.Router {
	// received router that not go have no one router inside and return router ready

	routes := routesUsers

	routes = append(routes, routeLogin)

	routes = append(routes, routePublications...)

	for _, route := range routes {

		if route.RequiredAuthentication {
			r.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authenticate(route.Function))).
				Methods(route.Method)
		} else {
			r.HandleFunc(route.URI,
				middlewares.Logger(route.Function)).
				Methods(route.Method)
		}

	}

	return r
}
