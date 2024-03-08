package api

import (
	"encoding/json"
	"net/http"
)

type BalanceParams struct {
	Username string
}

type BalanceResponse struct {
	Status  int
	Balance int64
}

type Error struct {
	Status  int
	Message string
}

func writeError(w http.ResponseWriter, message string, status int) {
	resp := Error{
		Status:  status,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An Unexpected Error Occurred.", http.StatusInternalServerError)
	}
)
