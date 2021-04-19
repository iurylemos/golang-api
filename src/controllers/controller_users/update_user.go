package controller_users

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

func UpdateUSer(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, erro := strconv.ParseUint(parameters["id"], 10, 64)

	if erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	userIDInToken, erro := authentication.ExtractUserID(r)

	if erro != nil {
		utils.ResponseError(w, http.StatusUnauthorized, erro)
		return
	}

	if userID != userIDInToken {
		utils.ResponseError(w, http.StatusForbidden, errors.New("not is possible update the user that no is your"))
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		utils.ResponseError(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.Usuario

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	if erro := user.Prepare("edit"); erro != nil {
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

	if erro = repository.Update(userID, user); erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusNoContent, nil)
}
