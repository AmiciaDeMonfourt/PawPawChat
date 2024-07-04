package web

import "pawpawchat/internal/model/domain"

// /api/user
type GetUserInfoRequest struct {
}

type GetUserInfoResponse struct {
	User domain.User `json:"user"`
}

// /profile
type ProfileRequest struct {
}

type ProfileResponse struct {
	User domain.User `json:"user"`
}

// /signup
type SignUpRequest struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type SignUpResponse struct {
	User     domain.User `json:"user"`
	TokenStr string      `json:"token_string"`
}

// /signin
type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignInResponse struct {
	User     domain.User `json:"user"`
	TokenStr string      `json:"token_string"`
}
