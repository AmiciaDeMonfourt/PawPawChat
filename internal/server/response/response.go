package response

import (
	"encoding/json"
	"net/http"
)

var (
	ResponseStatusSuccess = "sucess"
	ResponseStatusError   = "error"
)

type Response struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Code    int         `json:"code,omitempty"`
}

func JSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	response := &Response{
		Status: ResponseStatusSuccess,
		Data:   data,
	}

	json.NewEncoder(w).Encode(response)
}

func Error(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Application-Type", "application/json")
	w.WriteHeader(code)

	response := &Response{
		Status:  ResponseStatusError,
		Message: msg,
	}

	json.NewEncoder(w).Encode(response)
}
