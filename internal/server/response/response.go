package response

import (
	"encoding/json"
	"net/http"
)

type HTTPError struct {
	Error string `json:"error"`
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func OK(w http.ResponseWriter, data interface{}) {
	jsonResponse(w, http.StatusOK, data)
}

func Created(w http.ResponseWriter, data interface{}) {
	jsonResponse(w, http.StatusCreated, data)
}

func HTTPErrorResp(w http.ResponseWriter, statusCode int, err error) {
	jsonResponse(w, statusCode, &HTTPError{Error: err.Error()})
}

func BadReq(w http.ResponseWriter, err error) {
	HTTPErrorResp(w, http.StatusBadRequest, err)
}

func Conflict(w http.ResponseWriter, err error) {
	HTTPErrorResp(w, http.StatusConflict, err)
}

func Forbidden(w http.ResponseWriter, err error) {
	HTTPErrorResp(w, http.StatusForbidden, err)
}

func InternalErr(w http.ResponseWriter, err error) {
	HTTPErrorResp(w, http.StatusInternalServerError, err)
}

func NotFound(w http.ResponseWriter, err error) {
	HTTPErrorResp(w, http.StatusNotFound, err)
}

func Unauthorized(w http.ResponseWriter, err error) {
	HTTPErrorResp(w, http.StatusUnauthorized, err)
}
