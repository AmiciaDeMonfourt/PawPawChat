package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/generated/proto/users"
	"pawpawchat/internal/middleware"
	"pawpawchat/internal/model/domain"
	"pawpawchat/internal/model/web"
	"pawpawchat/internal/server/response"
)

// @Summary      Sign up
// @Description  Registration
// @Param        requestBody    body      web.SignUpRequest		true	"Credentials"
// @Success      201  			{object}  web.SignUpResponse
// @Failure      409  			{object}  response.HTTPError
// @Failure      500  			{object}  response.HTTPError
// @Router       /signup [post]
func (r *userRoutes) SignUp(w http.ResponseWriter, req *http.Request) {
	signUpReq := &web.SignUpRequest{}

	if err := json.NewDecoder(req.Body).Decode(&signUpReq); err != nil {
		response.BadReq(w, err.Error())
		return
	}

	if signUpReq.FirstName == "" || signUpReq.SecondName == "" || signUpReq.Email == "" || signUpReq.Password == "" {
		response.BadReq(w, fmt.Sprintf("missing parameters in request: %v", signUpReq))
		return
	}

	authSignUpReq := &auth.SignUpRequest{
		FirstName:  signUpReq.FirstName,
		SecondName: signUpReq.SecondName,
		Email:      signUpReq.Email,
		Password:   signUpReq.Password,
	}

	authResp, err := r.gRPCClient.Auth().SignUp(context.TODO(), authSignUpReq)
	if err != nil {
		response.InternalErr(w, err.Error())
		return
	}

	if authResp.GetError() != "" {
		response.Conflict(w, authResp.GetError())
		return
	}

	user := domain.NewUser(authResp)
	if user == nil {
		response.InternalErr(w, "failed to create a user model")
		return
	}

	response.Created(w, web.SignUpResponse{User: *user})
}

// @Summary      Sign in
// @Description  Authorization
// @Param        Authorization  header    string              	true	"Token"
// @Param        requestBody    body      web.SignInRequest		true	"Credentials"
// @Success      200  			{object}  web.SignInResponse
// @Failure      404  			{object}  response.HTTPError
// @Failure      500  			{object}  response.HTTPError
// @Router       /signin [post]
func (r *userRoutes) SignIn(w http.ResponseWriter, req *http.Request) {
	signInReq := &web.SignInRequest{}

	if err := json.NewDecoder(req.Body).Decode(&signInReq); err != nil {
		response.BadReq(w, err.Error())
		return
	}

	if signInReq.Email == "" || signInReq.Password == "" {
		response.BadReq(w, fmt.Sprintf("missing email or email in request: %v", signInReq))
		return
	}

	authSignInReq := &auth.SignInRequest{
		Email:    signInReq.Email,
		Password: signInReq.Password,
	}

	authResp, err := r.gRPCClient.Auth().SignIn(context.TODO(), authSignInReq)
	if err != nil {
		response.InternalErr(w, err.Error())
		return
	}

	if authResp.GetError() != "" {
		response.Forbidden(w, authResp.GetError())
		return
	}

	user := domain.NewUser(authResp)
	if user == nil {
		response.InternalErr(w, "failed to create a user model")
		return
	}

	response.OK(w, web.SignInResponse{User: *user})
}

// @Summary      User
// @Description  User info
// @Param        Authorization  header    string            	true	"Token"
// @Success      200  			{object}  web.GetUserInfoResponse
// @Failure      404  			{object}  response.HTTPError
// @Failure      500  			{object}  response.HTTPError
// @Router       /api/user [get]
func (r *userRoutes) GetInfo(w http.ResponseWriter, req *http.Request) {
	id := req.Context().Value(middleware.CtxString("user_id"))

	getByIdParams := &users.GetByIdRequest{Id: id.(uint64)}
	usersResp, err := r.gRPCClient.Users().GetById(context.TODO(), getByIdParams)
	if err != nil {
		response.NotFound(w, err.Error())
		return
	}

	if usersResp.GetError() != "" {
		response.InternalErr(w, usersResp.GetError())
		return
	}

	user := domain.NewUser(usersResp)
	if user == nil {
		response.InternalErr(w, "failed to create a user model")
		return
	}

	response.OK(w, web.GetUserInfoResponse{User: *user})
}
