package utils

import (
	"CleanArchitectureGo/internal/entities"
	"encoding/json"
	"net/http"
)

func DecodeErr(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(&entities.ErrorResponse{Message: err.Error()})
}
