package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/internal/grpcclient"
	"pawpawchat/internal/server/response"
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
// @Router       /signin [get]
func (r *userRoutes) SignIn(w http.ResponseWriter, req *http.Request) {
	signInReq := new(auth.SignInRequest)
	if err := json.NewDecoder(req.Body).Decode(&signInReq); err != nil {
		response.BadReq(w, err)
	}

	signResp, err := r.client.Auth().SignIn(context.TODO(), signInReq)
	if err != nil {
		response.InternalErr(w, err)
		return
	}

	response.OK(w, signResp)
}
