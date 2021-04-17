package controller_users

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/repositories"
	"api-nos-golang/src/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)

	if erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Connect()

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)

	if erro = repository.Delete(userID); erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusNoContent, nil)
}
