package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// type response struct {
// 	Message string `json:"message"`
// }

// func ResponseJSON(message string) ([]byte, error) {
// 	responseStruct := response{Message: message}
// 	responseJson, erro := json.Marshal(responseStruct)

// 	if erro != nil {
// 		return nil, erro
// 	}

// 	return responseJson, nil
// }

func ResponseError(w http.ResponseWriter, statusCode int, erro error) {
	// struct already filling out
	ResponseJSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}

func ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(data); erro != nil {
		log.Fatal(erro)
	}
}
