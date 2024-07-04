package domain

import (
	"pawpawchat/generated/proto/auth"
	"pawpawchat/generated/proto/users"
)

type User struct {
	Id         uint64 `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	Email      string `json:"email"`
	Password   string `json:"-"`
}

func NewUser(in any) *User {
	switch input := in.(type) {
	case *users.GetByIdResponse:
		return &User{
			Id:    input.GetUser().Id,
			Email: input.GetUser().Email,
		}
	case *auth.SignInResponse:
		return &User{
			Id:         input.GetUser().Id,
			FirstName:  input.GetUser().FirstName,
			SecondName: input.GetUser().SecondName,
			Email:      input.GetUser().Email,
		}
	case *auth.SignUpResponse:
		return &User{
			Id:         input.GetUser().Id,
			FirstName:  input.GetUser().FirstName,
			SecondName: input.GetUser().SecondName,
			Email:      input.GetUser().Email,
		}
	}
	return nil
}
