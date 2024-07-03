package response

import (
	"encoding/json"
	"net/http"
)

func writeJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

// SUCCESS
func OK(w http.ResponseWriter, data interface{}) {
	writeJSONResponse(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, data interface{}) {
	writeJSONResponse(w, http.StatusCreated, data)
}

// FAILED
func BadReq(w http.ResponseWriter, err error) {
	writeJSONResponse(w, http.StatusBadRequest, &HTTPError{Error: err.Error()})
}

func InternalErr(w http.ResponseWriter, err error) {
	writeJSONResponse(w, http.StatusInternalServerError, &HTTPError{Error: err.Error()})
}

func NotFound(w http.ResponseWriter, err error) {
	writeJSONResponse(w, http.StatusNotFound, &HTTPError{Error: err.Error()})
}

func Unauthorized(w http.ResponseWriter, err error) {
	writeJSONResponse(w, http.StatusUnauthorized, &HTTPError{Error: err.Error()})
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

// MODELS FOR SWAGGER
type HTTPError struct {
	Error string `json:"error"`
}
