package service

import (
	"pawpawchat/generated/proto/users"
	"pawpawchat/pkg/users/database"
	"pawpawchat/pkg/users/grpc"
)

type usersService struct {
	grpc *grpc.Client
	db   *database.DataBase
	users.UnimplementedUsersServiceServer
}

func newUsersService(db *database.DataBase) *usersService {
	return &usersService{
		grpc: grpc.NewClient(),
		db:   db,
	}
}
