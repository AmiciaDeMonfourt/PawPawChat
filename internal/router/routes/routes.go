package routes

import "net/http"

type User interface {
	SignUp(http.ResponseWriter, *http.Request)

	SignIn(http.ResponseWriter, *http.Request)

	Page(http.ResponseWriter, *http.Request)

	User(http.ResponseWriter, *http.Request)
}
