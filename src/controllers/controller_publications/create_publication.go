package controller_publications

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/middlewares/authentication"
	"api-nos-golang/src/models"
	"api-nos-golang/src/repositories"
	"api-nos-golang/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CreatePublication(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtractUserID(r)

	if erro != nil {
		utils.ResponseError(w, http.StatusUnauthorized, erro)
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		utils.ResponseError(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var publication models.Publicacao

	if erro = json.Unmarshal(bodyRequest, &publication); erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	publication.AutorID = userID

	db, erro := db.Connect()

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryPublications(db)

	publication.ID, erro = repository.Create(publication)

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, publication)
}
