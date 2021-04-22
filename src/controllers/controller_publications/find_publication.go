package controller_publications

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/repositories"
	"api-nos-golang/src/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FindPublication(w http.ResponseWriter, r *http.Request) {
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

	publication, erro := repository.FindForID(publicationID)

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, publication)
}
