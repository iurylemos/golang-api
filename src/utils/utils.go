package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

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

	if data != nil {
		if erro := json.NewEncoder(w).Encode(data); erro != nil {
			log.Fatal(erro)
		}
	}
}
