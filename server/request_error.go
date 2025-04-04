package server

import (
	"encoding/json"
	"net/http"
)

func HttpInternalServerError(w http.ResponseWriter, err error) {
	http.Error(w, "INTERNAL_SERVER_ERROR", http.StatusInternalServerError)
}

func HttpBadRequestError(w http.ResponseWriter, err error) {
	http.Error(w, "BAD_REQUEST", http.StatusBadRequest)
}

func HttpMethodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed)
}

func HttpCreated(w http.ResponseWriter, data any) {
	includeData(w, data)
	w.WriteHeader(http.StatusCreated)
}

func HttpOk(w http.ResponseWriter, data any) {
	includeData(w, data)
	w.WriteHeader(http.StatusOK)
}

func includeData(w http.ResponseWriter, data any) {
	if data != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(data)
	}
}
