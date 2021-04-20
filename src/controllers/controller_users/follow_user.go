package controller_users

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/middlewares/authentication"
	"api-nos-golang/src/repositories"
	"api-nos-golang/src/utils"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func FollowUser(w http.ResponseWriter, r *http.Request) {
	followID, erro := authentication.ExtractUserID(r)

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

	if followID == userID {
		utils.ResponseError(w, http.StatusForbidden, errors.New("não é possível seguir você mesmo"))
		return
	}

	db, erro := db.Connect()

	if erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryUsers(db)

	if erro = repository.Follow(userID, followID); erro != nil {
		utils.ResponseError(w, http.StatusInternalServerError, erro)
		return
	}

	utils.ResponseJSON(w, http.StatusNoContent, nil)
}
