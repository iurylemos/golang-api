package controller_publications

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/repositories"
	"api-nos-golang/src/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func LikePublication(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	publicationID, erro := strconv.ParseUint(parameters["id"], 10, 64)

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

	repository := repositories.NewRepositoryPublications(db)

	if erro = repository.Like(publicationID); erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusNoContent, nil)
}
