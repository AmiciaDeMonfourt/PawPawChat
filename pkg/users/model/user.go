package model

import "pawpawchat/generated/proto/users"

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
	HashPass string `json:"-"`
}

func NewUser(req *users.CreateRequest) (*User, error) {
	return &User{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}, nil
}
