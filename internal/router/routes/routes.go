package routes

import "net/http"

type User interface {
	SignUp(http.ResponseWriter, *http.Request)

	SignIn(http.ResponseWriter, *http.Request)

	Profile(http.ResponseWriter, *http.Request)

	GetInfo(http.ResponseWriter, *http.Request)
}
