package controller_users

import (
	"api-nos-golang/src/db"
	"api-nos-golang/src/models"
	"api-nos-golang/src/repositories"
	"api-nos-golang/src/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bodyRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var user models.Usuario

	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := db.Connect()

	if erro != nil {
		log.Fatal(erro)
	}

	//repositories have the responsibility of connecting with the database
	repository := repositories.NewRepositoryUsers(db)
	userID, erro := repository.Create(user)

	if erro != nil {
		log.Fatal(erro)
	}

	w.WriteHeader(http.StatusOK)

	responseJson, erro := utils.ResponseJSON(fmt.Sprintf("ID Inserido: %d", userID))

	if erro != nil {
		w.Write([]byte(fmt.Sprintf("Something wrong happened: %v", erro)))
		return
	}

	w.Write(responseJson)
}
