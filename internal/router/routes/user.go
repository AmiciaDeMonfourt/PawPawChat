package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"pawpawchat/generated/proto/auth"
	"pawpawchat/generated/proto/users"
	"pawpawchat/internal/grpcclient"
	"pawpawchat/internal/middleware"
	"pawpawchat/internal/model/domain"
	"pawpawchat/internal/model/web"
	"pawpawchat/internal/server/response"

	"github.com/gorilla/mux"
)

type userRoutes struct {
	gRPCClient *grpcclient.Client
}

func NewUserRoutes(gRPCClient *grpcclient.Client) *userRoutes {
	return &userRoutes{
		gRPCClient: gRPCClient,
	}
}

// @Summary      Sign up
// @Description  Registration
// @Param        requestBody    body       auth.SignInRequest	true	"Credentials"
// @Success      201  			{object}   auth.SignInResponse
// @Failure      409  			{object}  response.HTTPError
// @Failure      500  			{object}  response.HTTPError
// @Router       /signup [post]
func (r *userRoutes) SignUp(w http.ResponseWriter, req *http.Request) {
	signUpReq := new(auth.SignUpRequest)
	if err := json.NewDecoder(req.Body).Decode(&signUpReq); err != nil {
		response.BadReq(w, err.Error())
		return
	}

	signUpResp, err := r.gRPCClient.Auth().SignUp(context.TODO(), signUpReq)
	if err != nil {
		response.InternalErr(w, err.Error())
		return
	}

	if signUpResp.GetError() != "" {
		response.Conflict(w, signUpResp.GetError())
		return
	}

	response.Created(w, signUpResp)
}

// @Summary      Sign in
// @Description  Authorization
// @Param        Authorization  header    string              	true  "Token"
// @Param        requestBody    body      auth.SignInRequest	true  "Credentials"
// @Success      200  			{object}  auth.SignInResponse
// @Failure      404  			{object}  response.HTTPError
// @Failure      500  			{object}  response.HTTPError
// @Router       /signin [post]
func (r *userRoutes) SignIn(w http.ResponseWriter, req *http.Request) {
	signInReq := new(auth.SignInRequest)
	if err := json.NewDecoder(req.Body).Decode(&signInReq); err != nil {
		response.BadReq(w, err.Error())
		return
	}

	signInResp, err := r.gRPCClient.Auth().SignIn(context.TODO(), signInReq)
	if err != nil {
		response.InternalErr(w, err.Error())
		return
	}

	if signInResp.GetError() != "" {
		response.Forbidden(w, signInResp.GetError())
		return
	}

	response.OK(w, signInResp)
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

	user, err := r.gRPCClient.Users().GetByUsername(context.TODO(), &users.GetByUsernameRequest{Username: username})
	if err != nil {
		response.BadReq(w, err.Error())
		return
	}

	if user == nil {
		response.NotFound(w, err.Error())
		return
	}

	response.OK(w, user)
}

// @Summary      User
// @Description  User info
// @Param        Authorization  header    string            	true	"Token"
// @Success      200  			{object}  web.UserResponse
// @Failure      404  			{object}  response.HTTPError
// @Failure      500  			{object}  response.HTTPError
// @Router       /api/user [get]
func (r *userRoutes) User(w http.ResponseWriter, req *http.Request) {
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
		response.InternalErr(w, "could not create a user model")
		return
	}

	response.OK(w, web.UserResponse{User: *user})
}
