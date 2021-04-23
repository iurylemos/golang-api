package controller_publications

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/middlewares/authentication"
	"api-nos-golang/src/models"
	"api-nos-golang/src/repositories"
	"api-nos-golang/src/utils"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func UpdatePublication(w http.ResponseWriter, r *http.Request) {
	userID, erro := authentication.ExtractUserID(r)

	if erro != nil {
		utils.ResponseError(w, http.StatusUnauthorized, erro)
		return
	}

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

	publicationSalvedInDB, erro := repository.FindForID(publicationID)

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	if publicationSalvedInDB.AutorID != userID {
		utils.ResponseError(w, http.StatusForbidden, errors.New("não é possivel atualizar uma publicação que não é sua"))
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	var publication models.Publicacao

	if erro = json.Unmarshal(bodyRequest, &publication); erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	if erro = publication.Prepare(); erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.Update(publication.ID, publication); erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusNoContent, nil)
}
