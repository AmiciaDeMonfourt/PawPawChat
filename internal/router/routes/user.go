package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/generated/proto/users"
	"pawpawchat/internal/grpcclient"
	"pawpawchat/internal/server/response"

	"github.com/gorilla/mux"
)

type userRoutes struct {
	client *grpcclient.Client
}

func NewUserRoutes(client *grpcclient.Client) *userRoutes {
	return &userRoutes{
		client: client,
	}
}

// @Summary      Sign up
// @Description  Registration
// @Param        requestBody    body      auth.SignInRequest	true	"Credentials"
// @Success      200  			{object}   auth.SignInResponse
// @Failure      400  			{object}  response.HTTPError
// @Failure      500  			{object}  response.HTTPError
// @Router       /signup [post]
func (r *userRoutes) SignUp(w http.ResponseWriter, req *http.Request) {
	signUpReq := new(auth.SignUpRequest)
	if err := json.NewDecoder(req.Body).Decode(&signUpReq); err != nil {
		response.BadReq(w, err)
		return
	}

	singnUpResp, err := r.client.Auth().SignUp(context.TODO(), signUpReq)
	if err != nil {
		response.InternalErr(w, err)
		return
	}

	response.Created(w, singnUpResp)
}

// @Summary      Sign in
// @Description  Authorization
// @Param        Authorization  header    string              	true  "Token"
// @Param        requestBody    body      auth.SignInRequest	true  "Credentials"
// @Success      200  			{object}  auth.SignInResponse
// @Failure      400  			{object}  response.HTTPError
// @Failure      500  			{object}  response.HTTPError
// @Router       /signin [post]
func (r *userRoutes) SignIn(w http.ResponseWriter, req *http.Request) {
	signInReq := new(auth.SignInRequest)
	if err := json.NewDecoder(req.Body).Decode(&signInReq); err != nil {
		response.BadReq(w, err)
		return
	}

	signResp, err := r.client.Auth().SignIn(context.TODO(), signInReq)
	if err != nil {
		response.InternalErr(w, err)
		return
	}

	response.OK(w, signResp)
}

// @Summary      Page
// @Description  User's page
// @Param        Authorization  header    string              	true  "Token"
// @Success      200  			{object}  users.GetByUsernameResponse
// @Failure      401  			{object}  response.HTTPError
// @Failure      500  			{object}  response.HTTPError
// @Router       /{username} [get]
func (r *userRoutes) Page(w http.ResponseWriter, req *http.Request) {
	username := mux.Vars(req)["username"]

	user, err := r.client.Users().GetByUsername(context.TODO(), &users.GetByUsernameRequest{Username: username})
	if err != nil {
		response.BadReq(w, err)
		return
	}

	if user == nil {
		response.NotFound(w, err)
		return
	}

	response.OK(w, user)
}
