package routing

import (
	"encoding/json"
	"net/http"

	"github.com/kanca-studio/palang/delivery/http/schema"
)

func respondWithError(w http.ResponseWriter, code int, message string, err error) {
	var response schema.ErrorResponse
	response.Message = message
	response.Error = err.Error()
	respondWithJSON(w, code, response)
}

func respondWithSuccess(w http.ResponseWriter, payload interface{}) {
	respondWithJSON(w, http.StatusOK, payload)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}
