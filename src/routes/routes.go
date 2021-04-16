package routes

import (
	"api-nos-golang/src/routes/route"

	"github.com/gorilla/mux"
)

func CreateRoutes() *mux.Router {
	r := mux.NewRouter()
	return route.SettingsRoute(r)
}
