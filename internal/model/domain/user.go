package domain

import (
	"pawpawchat/generated/proto/users"
)

type User struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUser(in any) *User {
	switch input := in.(type) {
	case *users.GetByIdResponse:
		return &User{
			Id:       input.GetUser().Id,
			Username: input.GetUser().Username,
			Email:    input.GetUser().Email,
		}
	}

	return nil
}
