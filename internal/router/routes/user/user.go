package userroutes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/internal/grpcclient"
	"pawpawchat/internal/server/response"
)

type userRoutes struct {
	client *grpcclient.Client
}

func New(client *grpcclient.Client) *userRoutes {
	return &userRoutes{
		client: client,
	}
}

func (r *userRoutes) SignUp(w http.ResponseWriter, req *http.Request) {
	signUpReq := new(auth.SignUpRequest)
	if err := json.NewDecoder(req.Body).Decode(&signUpReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	log.Printf("userRoutes signUpReq: %v", signUpReq)

	singnUpResp, err := r.client.Auth().SignUp(context.TODO(), signUpReq)

	log.Printf("userRoutes signUpResp: %v", singnUpResp)

	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, singnUpResp)
}

func (r *userRoutes) SignIn(w http.ResponseWriter, req *http.Request) {
	signInReq := new(auth.SignInRequest)
	if err := json.NewDecoder(req.Body).Decode(&signInReq); err != nil {
		response.Error(w, http.StatusBadRequest, err.Error())
	}

	signResp, err := r.client.Auth().SignIn(context.TODO(), signInReq)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, signResp)
}
