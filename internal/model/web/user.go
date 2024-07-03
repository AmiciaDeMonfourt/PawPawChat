package web

import "pawpawchat/internal/model/domain"

type UserRequest struct {
	TokenStr string `json:"token_string"`
}

type UserResponse struct {
	User domain.User `json:"user"`
}

///

type UserSignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignUpResponse struct {
	User     domain.User `json:"user"`
	TokenStr string      `json:"token_string"`
	Error    string      `json:"error,omitempty"`
}

type UserSignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignInResponse struct {
	User     domain.User `json:"user"`
	TokenStr string      `json:"token_string"`
	Error    string      `json:"error,omitempty"`
}

type PageRequest struct {
	TokenStr string `json:"token_string"`
}

type PageResponse struct {
	User  domain.User `json:"user"`
	Error string      `json:"error,omitempty"`
}
