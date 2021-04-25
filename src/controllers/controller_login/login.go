package controller_login

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/middlewares/authentication"
	"api-nos-golang/src/middlewares/security"
	"api-nos-golang/src/models"
	"api-nos-golang/src/repositories"
	"api-nos-golang/src/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Login(w http.ResponseWriter, r *http.Request) {
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

	db, erro := db.Connect()

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	//obs:
	repository := repositories.NewRepositoryUsers(db)

	userSalvedInDB, erro := repository.FindForEmail(user.Email)

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(user.Senha, userSalvedInDB.Senha); erro != nil {
		utils.ResponseError(w, http.StatusUnauthorized, erro)
		return
	}

	token, erro := authentication.CreateToken(userSalvedInDB.ID)

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	userID := strconv.FormatUint(userSalvedInDB.ID, 10)

	utils.ResponseJSON(w, http.StatusOK, models.DataAuth{ID: userID, Token: token})
}
