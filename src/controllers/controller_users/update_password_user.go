package controller_users

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/middlewares/authentication"
	"api-nos-golang/src/middlewares/security"
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

func UpdatePasswordUser(w http.ResponseWriter, r *http.Request) {
	userIDToken, erro := authentication.ExtractUserID(r)

	if erro != nil {
		utils.ResponseError(w, http.StatusUnauthorized, erro)
		return
	}

	parameteres := mux.Vars(r)
	userID, erro := strconv.ParseUint(parameteres["id"], 10, 64)

	if erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	if userIDToken != userID {
		utils.ResponseError(w, http.StatusForbidden, errors.New("não é possível modificar senha de outro usuário"))
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	var password models.Senha

	/*
		{
			"nova": "123456"
			"atual": "789101"
		}

	*/

	if erro = json.Unmarshal(bodyRequest, &password); erro != nil {
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

	// check if current password that come to request is different of password save in database
	// find password current this user in database for id

	passwordInDB, erro := repository.FindPasswordForID(userID)

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	// this password that be save in variable passwordInDB be with hash
	// then i need compare with the current password
	// or be, i need compare with the password came of body request

	if erro = security.VerifyPassword(password.Atual, passwordInDB); erro != nil {
		utils.ResponseError(w, http.StatusUnauthorized, errors.New("a senha atual não condiz com a que está salva no database"))
		return
	}

	// i need put on a hash for this password
	// for me to put on it inside the database, saving password with hash alredy

	passwordHash, erro := security.Hash(password.Nova)

	if erro != nil {
		utils.ResponseError(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repository.UpdatePassword(userID, string(passwordHash)); erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusNoContent, nil)
}
