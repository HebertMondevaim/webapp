package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Erro representa a resposta de erro da API
type ErroApi struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal()
	}
}

// TratarStatusCodeDeErro trata as requisições com status code 400 ou superior.
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) {
	var erro ErroApi
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
}
