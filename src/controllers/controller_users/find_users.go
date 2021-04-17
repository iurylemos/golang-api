package controller_users

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/repositories"
	"api-nos-golang/src/utils"
	"net/http"
	"strings"
)

func FindUsers(w http.ResponseWriter, r *http.Request) {
	nameOrNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := db.Connect()

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)

	users, erro := repository.Find(nameOrNick)

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, users)
}
