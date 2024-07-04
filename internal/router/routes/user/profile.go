package user

import (
	"context"
	"net/http"
	"pawpawchat/generated/proto/users"
	"pawpawchat/internal/model/domain"
	"pawpawchat/internal/model/web"
	"pawpawchat/internal/server/response"

	"github.com/gorilla/mux"
)

// @Summary      Profile
// @Description  The page received basic information about the user
// @Param        Authorization  header    string              	true  "Token"
// @Success      200  			{object}  users.GetByUsernameResponse
// @Failure      401  			{object}  response.HTTPError
// @Failure      500  			{object}  response.HTTPError
// @Router       /{username} [get]
func (r *userRoutes) Profile(w http.ResponseWriter, req *http.Request) {
	username := mux.Vars(req)["username"]
	getByUsernameParams := &users.GetByUsernameRequest{Username: username}

	usersResp, err := r.gRPCClient.Users().GetByUsername(context.TODO(), getByUsernameParams)
	if err != nil {
		response.BadReq(w, err.Error())
		return
	}

	if usersResp.GetUser() == nil {
		response.NotFound(w, "user not found")
		return
	}

	user := domain.NewUser(usersResp)
	if user == nil {
		response.InternalErr(w, "failed to create a user model")
		return
	}

	response.OK(w, web.ProfileResponse{User: *user})
}
