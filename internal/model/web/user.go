package web

import "pawpawchat/internal/model/domain"

type UserSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpResponse struct {
	TokenStr string `json:"token_string"`
}

type UserSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignInResponse struct {
	TokenStr string      `json:"token_string"`
	Error    string      `json:"error,omitempty"`
	User     domain.User `json:"user"`
}
