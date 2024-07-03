package model

import (
	"pawpawchat/generated/proto/users"
	"pawpawchat/pkg/users/utils"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	HashPass string `json:"-"`
}

func NewUser(req *users.CreateRequest) (*User, error) {
	hashPass, err := utils.EncryptString(req.GetPassword())
	if err != nil {
		return nil, err
	}

	return &User{
		Username: req.GetUsername(),
		Email:    req.GetEmail(),
		HashPass: hashPass,
	}, nil
}

func (u *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.HashPass), []byte(password))
}
