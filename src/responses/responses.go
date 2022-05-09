package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if body != nil {
		if erro := json.NewEncoder(w).Encode(body); erro != nil {
			log.Fatal(erro)
		}
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
