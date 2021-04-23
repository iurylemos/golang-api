package controller_publications

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/middlewares/authentication"
	"api-nos-golang/src/repositories"
	"api-nos-golang/src/utils"
	"net/http"
)

func FindPublications(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtractUserID(r)

	if erro != nil {
		utils.ResponseError(w, http.StatusUnauthorized, erro)
		return
	}

	db, erro := db.Connect()

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryPublications(db)

	publications, erro := repository.FindPublications(userID)

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusOK, publications)
}
